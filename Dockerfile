# syntax=docker/dockerfile:1

FROM golang:1.16-alpine as builder

RUN apk update \
  && apk add --no-cache git \
  && go get -u github.com/cosmtrek/air \
  && chmod +x ${GOPATH}/bin/air

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . /app

RUN GOOS=linux GOARCH=amd64 go build -o /main

FROM alpine:3.9

COPY --from=builder /main .

ENV PORT=${PORT}
ENTRYPOINT ["/main"]