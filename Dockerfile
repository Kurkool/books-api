# --- Stage 1: Build ---
# Use an updated version of the Golang image that matches the go.mod requirement.
FROM golang:1.24-alpine AS builder

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to download dependencies
COPY go.mod go.sum ./
# Download dependencies. This is cached if go.mod and go.sum don't change.
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the Go application
# -o main: specifies the output file name
# -ldflags="-w -s": strips debugging information, reducing binary size
RUN CGO_ENABLED=0 GOOS=linux go build -o main -ldflags="-w -s" .

# --- Stage 2: Run ---
# Use a minimal base image for the final container.
# alpine is a good choice for its small size.
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the application
# The binary is the entry point.
CMD ["./main"]
