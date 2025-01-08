# Use official Golang image as base
FROM golang:1.21-alpine

# Set working directory
WORKDIR /app

# Install required system dependencies
RUN apk update && apk add --no-cache \
    git \
    gcc \
    musl-dev \
    bash

# Install Beego and Bee tool
RUN go install github.com/beego/bee/v2@latest && \
    go install github.com/beego/beego/v2@latest

# Copy go mod files first for better caching
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application
COPY . .

# Build the application
RUN go build -o main .

# Expose the port the app runs on
EXPOSE 8080

# Command to run the application using bee (for hot reload during development)
CMD ["bee", "run"]