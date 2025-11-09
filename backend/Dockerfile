# syntax=docker/dockerfile:1
FROM golang:1.25-alpine AS builder

# Install build dependencies (sqlite headers)
RUN apk add --no-cache gcc musl-dev sqlite-dev

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build for Linux AMD64 with CGO enabled so the sqlite3 c library is linked dynamically
ENV CGO_ENABLED=1
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o server ./cmd

# Final minimal image
FROM alpine:3.18

# Install SQLite runtime lib
RUN apk add --no-cache sqlite-libs

WORKDIR /app
COPY --from=builder /app/server .

# Create and mount database directory
VOLUME ["/app/database"]

EXPOSE 8080
CMD ["./server"]
