FROM golang:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download && go mod verify

RUN go install github.com/cosmtrek/air@latest

# Copy the source from the current directory to the workspace
COPY ./ /app/

# Build the Go app
RUN go build -o main cmd/server/main.go

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD "air"