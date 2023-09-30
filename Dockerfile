FROM golang:1.21.1-alpine3.18

WORKDIR /usr/src/app

COPY . .
RUN go mod tidy