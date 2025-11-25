FROM node:20-alpine AS css-builder

WORKDIR /app

COPY package*.json tailwind.config.js postcss.config.js ./
RUN npm ci --only=production=false --no-audit --no-fund

# Copy all files needed for CSS processing
COPY static ./static/
COPY layouts ./layouts/
COPY content ./content/

# Create symlink for node_modules to make CSS imports work
RUN ln -sf /app/node_modules /app/static/

RUN npm run build:css:prod

FROM klakegg/hugo:ext-alpine AS builder

WORKDIR /app

# Copy Hugo configuration and content files first
COPY hugo.toml ./
COPY content ./content/
COPY layouts ./layouts/

# Copy static files except CSS (will be copied from css-builder)
COPY static ./static/
# Copy the built CSS from the css-builder stage (overwrite the original)
COPY --from=css-builder /app/static/css/main.min.css ./static/css/main.min.css

# Build the site
RUN hugo --minify

# Clean up unnecessary files from the public directory
RUN find /app/public -name "*.map" -delete && \
    find /app/public -name "*.css" ! -name "*.min.css" -delete

FROM rtsp/lighttpd

RUN mkdir -p /var/log/lighttpd /var/cache/lighttpd/uploads /var/cache/lighttpd/compress && \
    chmod -R 755 /var/cache/lighttpd

COPY lighttpd.conf /etc/lighttpd/lighttpd.conf
COPY --from=builder /app/public /var/www/html/

EXPOSE 80/tcp
