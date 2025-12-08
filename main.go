package main

import (
	"bytes"
	"embed"
	"errors"
	"io"
	"io/fs"
	"log"
	"net"
	"net/http"
	"os"
	"path"
	"strings"
	"time"
)

//go:embed public
var staticFiles embed.FS

func main() {
	publicFS, err := fs.Sub(staticFiles, "public")
	if err != nil {
		log.Fatalf("failed to load embedded files: %v", err)
	}

	fileServer := http.FileServer(http.FS(publicFS))
	handler := loggingMiddleware(notFoundMiddleware(publicFS, errorPageMiddleware(publicFS, fileServer)))

	host := os.Getenv("HOST")
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	addr := net.JoinHostPort(host, port)
	log.Printf("listening on %s", addr)

	if err := http.ListenAndServe(addr, handler); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Fatalf("server error: %v", err)
	}
}

// notFoundMiddleware renders the custom 404 page when the underlying handler returns 404.
func notFoundMiddleware(files fs.FS, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Buffer the downstream response to detect a 404 before sending anything.
		br := newBufferRecorder(w)
		next.ServeHTTP(br, r)

		if br.status == http.StatusNotFound && statusForPath(r.URL.Path) == 0 {
			serveErrorFilePath(files, w, r, http.StatusNotFound, "/404.html")
			return
		}

		br.flush()
	})
}

type bufferRecorder struct {
	sink        http.ResponseWriter
	status      int
	header      http.Header
	buf         bytes.Buffer
	wroteHeader bool
}

func newBufferRecorder(w http.ResponseWriter) *bufferRecorder {
	return &bufferRecorder{
		sink:   w,
		status: http.StatusOK,
		header: make(http.Header),
	}
}

func (r *bufferRecorder) Header() http.Header {
	return r.header
}

func (r *bufferRecorder) WriteHeader(status int) {
	if r.wroteHeader {
		return
	}
	r.wroteHeader = true
	r.status = status
}

func (r *bufferRecorder) Write(b []byte) (int, error) {
	if !r.wroteHeader {
		r.WriteHeader(http.StatusOK)
	}
	return r.buf.Write(b)
}

func (r *bufferRecorder) flush() {
	// copy headers set by downstream
	dst := r.sink.Header()
	for k, vv := range r.header {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}

	r.sink.WriteHeader(r.status)
	if r.buf.Len() > 0 {
		_, _ = r.sink.Write(r.buf.Bytes())
	}
}

// errorPageMiddleware intercepts direct requests to error pages and ensures matching HTTP status codes.
func errorPageMiddleware(files fs.FS, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		code := statusForPath(r.URL.Path)
		if code == 0 {
			next.ServeHTTP(w, r)
			return
		}

		serveErrorFilePath(files, w, r, code, r.URL.Path)
	})
}

func statusForPath(p string) int {
	switch path.Clean(p) {
	case "/400.html":
		return http.StatusBadRequest
	case "/401.html":
		return http.StatusUnauthorized
	case "/403.html":
		return http.StatusForbidden
	case "/404.html":
		return http.StatusNotFound
	case "/405.html":
		return http.StatusMethodNotAllowed
	case "/408.html":
		return http.StatusRequestTimeout
	case "/410.html":
		return http.StatusGone
	case "/500.html":
		return http.StatusInternalServerError
	case "/502.html":
		return http.StatusBadGateway
	case "/503.html":
		return http.StatusServiceUnavailable
	case "/429.html":
		return http.StatusTooManyRequests
	case "/451.html":
		return http.StatusUnavailableForLegalReasons
	case "/504.html":
		return http.StatusGatewayTimeout
	default:
		return 0
	}
}

func serveErrorFilePath(files fs.FS, w http.ResponseWriter, r *http.Request, code int, p string) {
	filePath := path.Clean(p)
	filePath = strings.TrimPrefix(filePath, "/")

	file, err := files.Open(filePath)
	if err != nil {
		w.WriteHeader(code)
		return
	}
	defer file.Close()

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.WriteHeader(code)

	if r.Method == http.MethodHead {
		return
	}

	_, _ = io.Copy(w, file)
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := &statusRecorder{ResponseWriter: w, status: 200}
		next.ServeHTTP(rec, r)
		duration := time.Since(start)

		log.Printf(
			`{ "time": "%s", "remote_ip": "%s", "method": "%s", "uri": "%s", "query_string": "%s", "status": %d, "bytes_sent": %d, "referer": "%s", "user_agent": "%s", "duration_ms": %.3f }`,
			time.Now().Format(time.RFC3339Nano),
			r.RemoteAddr,
			r.Method,
			r.URL.Path,
			r.URL.RawQuery,
			rec.status,
			rec.bytes,
			r.Referer(),
			r.UserAgent(),
			float64(duration.Microseconds())/1000.0,
		)
	})
}

type statusRecorder struct {
	http.ResponseWriter
	status int
	bytes  int64
}

func (r *statusRecorder) WriteHeader(status int) {
	r.status = status
	r.ResponseWriter.WriteHeader(status)
}

func (r *statusRecorder) Write(b []byte) (int, error) {
	n, err := r.ResponseWriter.Write(b)
	r.bytes += int64(n)
	return n, err
}
