# Build stage
FROM golang:1.21-alpine3.19 AS builder
LABEL authors="Jake Willmsen"

WORKDIR /app

COPY . .

RUN go build -o bin .

# Run stage
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/bin .
COPY --from=builder /app/config.json .

ENTRYPOINT ["/app/bin"]