VERSION ?= $(shell git describe --tags --always)

IMAGE = turing-ml/turing-api
PKG = github.com/turing-ml/turing-api
PKGS = $(shell go list ./... | grep -v /vendor/)

LDFLAGS = "-s -w -X github.com/turing-ml/turing-api/pkg/version.Version=$(VERSION)"

OS ?= linux
ARCH ?= amd64

build:
	GOOS=$(OS) GOARCH=$(ARCH) go build -o bin/turing -a -tags netgo -ldflags $(LDFLAGS)
 
test:
	@go test $(PKGS)

lint:
	@for pkg in $(PKGS) ; do golint $$pkg ; done

documentation:
	@swagger -apiPackage="github.com/turing-ml/turing-api" -mainApiFile=main.go -output=./docs/API.md -format=markdown

vet:
	@go vet $(PKGS)

coverage:
	@go test ./server -coverprofile=./coverage/coverage.out -o ./coverage/coverage.html
	@go tool cover -html=./coverage/coverage.out
	@go test ./server -covermode=count -coverprofile=./coverage/count.out fmt
	@go tool cover -func=./coverage/count.out

docker-all: docker-build docker-image

docker-build:
	@docker run -i --rm -v "$(PWD):/go/src/$(PKG)" -w /go/src/$(PKG) golang:1.10 make build OS=linux ARCH=amd64

docker-test:
	@docker run -i --rm -v "$(PWD):/go/src/$(PKG)" -w /go/src/$(PKG) --network greeny_default golang:1.10 make test

docker-image:
	@docker build -t $(IMAGE):$(VERSION) .
	@docker tag $(IMAGE):$(VERSION) $(IMAGE):latest
	@echo " ---> $(IMAGE):$(VERSION)\n ---> $(IMAGE):latest"

docker-push:
	@docker push $(IMAGE):$(VERSION)
	@docker push $(IMAGE):latest