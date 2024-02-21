FROM golang:1.22-alpine3.19 as builder
RUN mkdir /app
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
ENTRYPOINT [ "air","dev" ]