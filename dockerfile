FROM golang:1.23-alpine

# Set working directory
WORKDIR /app

# Install required packages
RUN apk add --no-cache git tzdata

# Copy go.mod and go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the source code
COPY . .

# Build the Go app
RUN go build -o weatherapp

# Set environment variables
ENV GIN_MODE=release

# Run the app
CMD ["./weatherapp"]
