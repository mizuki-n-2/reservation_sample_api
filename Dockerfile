# syntax=docker/dockerfile:1

FROM golang:1.16-alpine as builder

RUN apk update \
  && apk add --no-cache git curl \
  && go get -u github.com/cosmtrek/air \
  && go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /app

RUN go get github.com/google/wire/cmd/wire \
  && wire /app/di/wire.go

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . /app

RUN GOOS=linux GOARCH=amd64 go build -o /main

FROM alpine:3.9

COPY --from=builder /main .

ENV PORT=${PORT}
ENTRYPOINT ["/main"]