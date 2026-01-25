# -------- BUILD STAGE --------
FROM golang:1.24-alpine AS builder

WORKDIR /app

# Copy go files
COPY go.mod ./
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -o app ./cmd/api

# -------- RUN STAGE --------
FROM alpine:latest

WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/app .

EXPOSE 8080

# Run app
CMD ["./app"]