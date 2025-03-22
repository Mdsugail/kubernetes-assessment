# Use the official Golang image as a base
FROM golang:1.19

# Set the working directory
WORKDIR /app

# Copy the Go files into the container
COPY main.go .

# Build the Go application
RUN go build -o server main.go

# Expose the port
EXPOSE 8080

# Run the Go application
CMD ["/app/server"]
