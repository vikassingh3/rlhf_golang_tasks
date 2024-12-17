# Start from the official Golang base image
FROM golang:1.20-alpine

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod ./
COPY go.sum ./

# Download all dependencies. Dependencies are cached if the go.mod and go.sum files are not changed
RUN go mod tidy

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 8080 to the outside world
EXPOSE 8080

# Run the Go app
CMD ["./main"]
