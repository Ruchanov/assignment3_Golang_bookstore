## Build stage
#FROM golang:1.19.4-alpine3.16 AS builder
#WORKDIR /app
#COPY . .
#RUN go build -o main main.go
#
## Run stage
#FROM alpine:3.16
#WORKDIR /app
#COPY --from=builder /app/main .
#COPY wait-for.sh .
#
#EXPOSE 8080
#CMD ["/app/main"]
#
FROM golang:latest

WORKDIR /app

COPY ./ ./
# Copy the Go module files
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code to the container
COPY . .

# Build the application binary
RUN go build -o app .

# Expose the port that the application listens on
EXPOSE 8080

# Set the command to run when the container starts
CMD ["./app"]