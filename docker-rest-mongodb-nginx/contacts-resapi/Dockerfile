FROM golang as builder
WORKDIR /go/src/github.com/user/app
COPY . .
RUN set -x && \
    go get -d -v . && \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM scratch
WORKDIR /root/
COPY --from=builder /go/src/github.com/user/app .
EXPOSE 3000
CMD ["./app"]