# syntax=docker/dockerfile:1

FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . /app

RUN go build -o /reservation_sample_api

EXPOSE 8080

CMD [ "/reservation_sample_api" ]