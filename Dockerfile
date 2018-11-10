FROM golang:alpine
MAINTAINER Davide Berdin "davideberdin@gmail.com"

WORKDIR /go/src/github.com/turing-ml/turing-api

COPY . .

RUN mkdir -p ./bin && make build

ENV GIN_MODE release

EXPOSE 8080

CMD ["./bin/turing-api"]