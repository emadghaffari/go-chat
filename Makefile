
GOPATH:=$(shell go env GOPATH)
.PHONY: build
build:
	go build -o micro *.go

.PHONY: test
test:
	go test -v ./... -cover -race

.PHONY: vendor
vendor:
	go get ./...
	go mod vendor
	go mod verify

.PHONY: config
config:
	cp -rf ./config.example.yaml ./config.yaml
	cp -rf ./config.example.yaml ./config.test.yaml
