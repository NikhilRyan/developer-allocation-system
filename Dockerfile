# Build Stage
FROM golang:1.23.2 AS builder

WORKDIR /app

# Copy go.mod and go.sum for dependency resolution
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the source code
COPY . .

# Build the application binary
RUN go build -o main ./cmd/app/main.go
RUN go build -o migrate ./cmd/migrate/main.go

# Run Stage
FROM alpine:latest

WORKDIR /root/

# Install necessary libraries
RUN apk add --no-cache libc6-compat

# Copy the compiled binaries from the builder stage
COPY --from=builder /app/main .
COPY --from=builder /app/migrate .
COPY ./entrypoint.sh .

# Make the entrypoint script executable
RUN chmod +x ./entrypoint.sh

# Expose the application port
EXPOSE 8080

# Use the entrypoint script
ENTRYPOINT ["./entrypoint.sh"]

# Default command to start the application
CMD ["./main"]
