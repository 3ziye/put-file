# Use official Go image as build environment
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download dependencies (if any)
RUN go mod download || echo "No dependencies to download"

# Copy source code
COPY . .

# Compile Go program
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o put-file cmd/server/main.go

# Use Alpine as runtime environment
FROM alpine:3.18

# Add ca-certificates to support HTTPS (if needed)
RUN apk --no-cache add ca-certificates

# Create user to run the app as non-root
RUN adduser -D -g 'app' appuser

# Set working directory
WORKDIR /app

# Copy compiled binary from build stage
COPY --from=builder /app/put-file .

# Create upload directory and set permissions
RUN mkdir -p /app/uploads && chown -R appuser:appuser /app/uploads

# Switch to non-root user
USER appuser

# Expose port (default 8080)
EXPOSE 8080

# Set environment variables, can be overridden at runtime
ENV SERVER_PORT=8080
ENV SERVER_ROOT=/app/uploads
ENV SERVER_USERNAME=admin
ENV SERVER_PASSWORD=admin123
ENV SERVER_LOG_LEVEL=INFO
ENV SERVER_LOG_FILE=

# Run command, using environment variables as parameters
CMD ["sh", "-c", "./put-file --port=$SERVER_PORT --root=$SERVER_ROOT --username=$SERVER_USERNAME --password=$SERVER_PASSWORD --log-level=$SERVER_LOG_LEVEL --log-file=$SERVER_LOG_FILE"]