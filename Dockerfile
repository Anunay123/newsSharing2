# Start from a minimal Go image
FROM public.ecr.aws/docker/library/golang:alpine3.19 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

#Expose port 8080
EXPOSE 8080

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Command to run the executable
CMD ["./main"]