# syntax=docker/dockerfile:1

FROM golang:1.21.0-alpine3.18

# Destination where all commands will run now.
WORKDIR /app

# Download go modules 
COPY go.mod go.sum ./
RUN go mod download
RUN go mod verify

# Copy source code
COPY . ./

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /go-service

EXPOSE 8085

# First command to run as container starts
CMD ["/go-service"]

