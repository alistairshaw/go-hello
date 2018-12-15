FROM golang:1.9-alpine

WORKDIR /go/src/go-hello

RUN apk add --no-cache git
RUN go get github.com/codegangsta/gin

ENTRYPOINT gin -a 80 run main.go