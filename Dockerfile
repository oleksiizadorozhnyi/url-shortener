# Use the official Golang image as the base image
FROM golang:1.18-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy the Go module files to the working directory
COPY go.mod go.sum ./

# Download and install the application dependencies
RUN go mod download

# Copy the rest of the application source code to the working directory
COPY . .

# Build the Go application
RUN go build cmd/server/main.go

# Expose the port that the application listens on
EXPOSE 8080

# Set the entrypoint command to run the application
CMD ["./main"]