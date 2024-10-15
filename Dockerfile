# Stage 1: Build the Go binary
FROM golang:1.20-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files haven't changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o auth_microservice ./cmd/web

# Stage 2: Create a lightweight image for the application
FROM alpine:latest

# Set working directory in the final image
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/auth_microservice .

COPY .env .env


# Expose the port the app runs on
EXPOSE 3000

ENV APP_ENV=production

# Run the binary program as the entrypoint
CMD ["./auth_microservice"]
