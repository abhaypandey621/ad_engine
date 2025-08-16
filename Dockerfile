# syntax=docker/dockerfile:1

# --- Build stage ---
FROM golang:1.24.3-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o server cmd/server/main.go

# --- Run stage ---
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app/server ./server
COPY configs ./configs
COPY res ./res
COPY docs ./docs
EXPOSE 8080
ENTRYPOINT ["./server"] 
