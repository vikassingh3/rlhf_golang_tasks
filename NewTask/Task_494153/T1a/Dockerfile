# Use the official Golang base image
FROM golang:1.20 AS build

# Set the working directory inside the container
WORKDIR /go/src/app

# Copy the go.mod and go.sum files
COPY go.mod go.sum ./

# Run go mod tidy to ensure all dependencies are up-to-date
RUN go mod tidy

# Copy the rest of the application code
COPY . .

# Build the application
RUN go build -o /go/bin/app

# Create a new stage that doesn't contain the Golang build environment
FROM scratch

# Copy the compiled binary to the root directory
COPY --from=build /go/bin/app /app

# Set the entry point for the Docker container
ENTRYPOINT ["./app"]