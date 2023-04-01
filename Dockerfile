# syntax=docker/dockerfile:1


FROM golang:1.17-alpine
RUN apk add build-base
RUN mkdir /app
WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY . ./
VOLUME /database

RUN GO111MODULE=on CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o /gorala-link-shortener

EXPOSE 1333

CMD [ "/gorala-link-shortener" ]