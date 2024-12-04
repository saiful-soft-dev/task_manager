# Use Go official image
FROM golang:1.23

# Use the official Golang image as the base image
# FROM golang:1.20

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod go.sum ./

# Download dependencies early for better caching
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the app will run on
EXPOSE 8080

# Command to run the application
CMD ["./main"]
