FROM node:20-alpine AS css-builder

WORKDIR /app

COPY package*.json tailwind.config.js postcss.config.js ./
RUN npm ci --only=production=false --no-audit --no-fund

COPY . .
RUN npm run build:css && \
    rm -rf node_modules package*.json tailwind.config.js postcss.config.js

FROM klakegg/hugo:ext-alpine AS builder

WORKDIR /app

COPY . .
COPY --from=css-builder /app/static/css/main.min.css ./static/css/main.min.css

RUN hugo --minify && \
    rm -rf content layouts static hugo.toml && \
    find /app/public -name "*.css" ! -name "*.min.css" -delete && \
    find /app/public -name "*.map" -delete

FROM rtsp/lighttpd

RUN mkdir -p /var/log/lighttpd /var/cache/lighttpd/uploads /var/cache/lighttpd/compress && \
    chmod -R 755 /var/cache/lighttpd

COPY lighttpd.conf /etc/lighttpd/lighttpd.conf
COPY --from=builder /app/public /var/www/html/

EXPOSE 80/tcp
