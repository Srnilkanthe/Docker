# Stage 1: Builder
FROM golang:1.16 AS builder

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Download dependencies
RUN go mod download

# Build the application
RUN go build -o myapp

# Stage 2: Runner
FROM alpine:latest

# Set the working directory
WORKDIR /root/

# Copy the compiled binary from the builder stage
COPY --from=builder /app/myapp .

# Command to run the application
CMD ["./myapp"]