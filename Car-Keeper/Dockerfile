FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the Go modules files and download dependencies
COPY go.mod go.sum ./

# Download Go module dependencies
RUN go mod download

# Copy the entire project into the container
COPY . .

# Build the Go application
RUN go build -o main .

# Expose the port the application will run on
EXPOSE 8080

# Command to run the application
CMD ["./main"]

