FROM golang
MAINTAINER Davide Berdin "davideberdin@gmail.com"

WORKDIR /go/src/github.com/turing-ml/turing-api

COPY . .

RUN mkdir -p ./bin && make build

EXPOSE 8000

CMD ["./bin/turing", "api"]