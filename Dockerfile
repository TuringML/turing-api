FROM golang
MAINTAINER Davide Berdin "davideberdin@gmail.com"

WORKDIR /go/src/github.com/turing-ml/turing-api

# for generating the documentation
RUN go get github.com/yvasiyarov/swagger

COPY . .

RUN make documentation

RUN mkdir -p ./bin && make build

EXPOSE 8000

CMD ["./bin/turing", "api"]