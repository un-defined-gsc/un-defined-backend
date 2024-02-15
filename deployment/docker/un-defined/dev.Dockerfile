FROM golang:1.21.5-alpine3.19 as builder
RUN mkdir /app
WORKDIR /app
RUN go install github.com/cosmtrek/air@latest
ENTRYPOINT [ "air","dev" ]