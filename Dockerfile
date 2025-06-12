# FROM golang:1.24 AS builder
FROM golang:1.23-bookworm AS builder
WORKDIR /app

COPY go.* ./

# Download the Go module dependencies
RUN go mod download

COPY . ./

RUN go build -v -o /server ./cmd/server
 
# FROM alpine:latest AS run
FROM debian:bookworm-slim
RUN set -x && apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    rm -rf /var/lib/apt/lists/*
    
# Copy the application executable from the build image
COPY --from=builder /server /server

# WORKDIR /app
EXPOSE 8080
CMD ["/server"]

