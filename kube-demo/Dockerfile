FROM golang:alpine as builder
LABEL authors="Jake Willmsen"

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o hellogopher -ldflags="-w -s"

ENTRYPOINT ["./hellogopher"]