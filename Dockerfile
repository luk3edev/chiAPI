# Dockerfile
FROM golang:1.23-alpine

# Set working directory
WORKDIR /app

# Install dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy app source code
COPY . .

# Build the Go app
RUN go build -o app .

# Expose API port
EXPOSE 8080

# Start the app
CMD ["./app"]
