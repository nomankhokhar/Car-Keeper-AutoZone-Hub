FROM golang:1.24.4-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum from the root of the project (relative to the Docker build context)
COPY ../go.mod ../go.sum ./
RUN go mod download

# Copy the rest of the project files into the container
COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main cmd/api/main.go

FROM alpine:latest

# Install certificates (needed for HTTPS requests)
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the built binary from the builder stage
COPY --from=builder /app/main .

EXPOSE 8080

CMD ["./main"]
