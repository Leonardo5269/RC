# Step 1: Build Stage
FROM golang:1.23-alpine AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Copy the .env file
COPY .env .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -o server ./cmd/app

# Step 2: Run Stage
FROM alpine:latest

# Set the Current Working Directory inside the container
WORKDIR /root/

# Copy the Pre-built binary file from the build stage
COPY --from=builder /app/server .

# Copy templates and public files
COPY --from=builder /app/views ./views
COPY --from=builder /app/public ./public

# Copy the .env file
COPY --from=builder /app/.env .

# Expose port
EXPOSE 80

# Command to run the executable
CMD ["./server"]
