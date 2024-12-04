# Use Go official image
FROM golang:1.23

# Set the working directory
WORKDIR /app

# Copy Go modules and source code
COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Build the Go application
RUN go build -o main .

# Expose port 8080
EXPOSE 8080

# Command to run the app
CMD ["./main"]
