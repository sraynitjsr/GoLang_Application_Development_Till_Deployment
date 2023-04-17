# Use an official Golang runtime as a parent image
FROM golang:1.16-alpine3.13

# Set the working directory to /app
WORKDIR /app

# Copy the current directory contents into the container at /app
COPY . /app

# Build the Go app
RUN go build -o hello

# Set the command to run the executable
CMD ["./hello"]
