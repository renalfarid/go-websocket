# Use the official Golang image as a base image
FROM golang:latest

# Set the working directory inside the container
WORKDIR /app

# Copy the go.mod and go.sum files to the working directory
COPY go.mod .
COPY go.sum .

# Download and install any dependencies
RUN go mod download

# Copy the rest of the application code to the working directory
COPY . .

# Build the application
RUN go build -o schoolapi .

# Expose the port that the application will run on
EXPOSE 3000
EXPOSE 3030

# Command to run the application
CMD ["./schoolapi"]
