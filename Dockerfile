# Stage 1: Build the Go binary
FROM golang:1.23 AS builder

WORKDIR /app

# Copy go mod and sum first, then download deps
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source
COPY . .

# Build the Go app (static binary)
RUN CGO_ENABLED=0 GOOS=linux go build -o app main.go

# Stage 2: Minimal runtime image
FROM alpine:latest

WORKDIR /root/

# Copy static files
COPY --from=builder /app/static ./static

# Copy binary from builder
COPY --from=builder /app/app .

# Expose port
EXPOSE 3000

# Run app
CMD ["./app"]
