
FROM golang:1.13-alpine3.10

RUN mkdir /goapp

WORKDIR /goapp

EXPOSE 9090:9090
