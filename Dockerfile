# Use the official Golang image as the base image
FROM golang:1.20-alpine

# Set the working directory inside the container
WORKDIR /app

# Copy go.mod file to the workspace
COPY go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod file is not changed
RUN go mod download

# Copy the source code to the workspace
COPY . .

# Build the Go app
RUN go build -o server .

# Expose port 8000 to the outside world
EXPOSE 8000

# Command to run the executable
CMD ["./server"]
