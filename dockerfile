# Dockerfile
######################################################################
# Stage 1: Build the Go application
######################################################################
FROM golang:1.25-alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . .

RUN go build -o /app/hugo_comments src/main.go

######################################################################
# Stage 2: Create the minimal production image
######################################################################
FROM alpine:latest

RUN mkdir -p /opt/hugo_comments
WORKDIR /opt/hugo_comments
COPY --from=builder /app/hugo_comments .

EXPOSE 8080
CMD ["./hugo_comments"]