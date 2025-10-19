# Builder stage: compile the Go binary with CGO enabled (required by mattn/go-sqlite3)
FROM golang:1.25.1 AS builder
WORKDIR /src

# Copy go.mod and go.sum first for caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source and build
COPY . .

# Build for Linux AMD64 with CGO enabled so the sqlite3 c library is linked dynamically
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

RUN mkdir -p /app/bin
RUN go build -v -o /app/bin/app ./cmd


# Final stage: slim runtime image with libsqlite3 available
FROM debian:bookworm-slim

# Install sqlite runtime dependency and ca-certificates for TLS
RUN apt-get update && apt-get install -y --no-install-recommends \
    libsqlite3-0 \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

RUN mkdir -p /app/bin /data

# Copy binary from builder
COPY --from=builder /app/bin/app /app/bin/app

# Create a non-root user and set ownership
RUN useradd -m appuser && chown -R appuser:appuser /app /data

USER appuser
WORKDIR /app

ENV PATH=/app/bin:$PATH

EXPOSE 8080

ENTRYPOINT ["/app/bin/app"]
