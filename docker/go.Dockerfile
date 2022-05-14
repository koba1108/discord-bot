FROM golang:1.17.5-alpine3.15

ENV CGO_ENABLED=0

RUN apk update && apk upgrade && apk add --update figlet && go get -u github.com/cespare/reflex
