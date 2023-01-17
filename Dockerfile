FROM golang:latest

# Set the working directory
WORKDIR /app

# Copy the source code into the container
COPY . .

# Build the Go binary
RUN go build -o main .

# Expose the port that the API will run on
EXPOSE 8080

# Start the API
CMD ["./main"]