# Use an official Golang runtime as a parent image
FROM golang:1.20.3-alpine

# Set the working directory to /api
WORKDIR /api

# Copy go.mod and go.sum files to the container
COPY go.mod ./

# Download the dependencies
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go application
RUN go build -o bin/api ./cmd/sendit

# Expose port 8080
EXPOSE 8080

# Set the entry point of the container to the executable
CMD ["./bin/api"]

