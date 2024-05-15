# Start with the official Golang base image
FROM golang:1.22-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go application
RUN go build -o ulehub ./cmd/ulehub

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["./ulehub"]