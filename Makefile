TEST ?= $(shell go list ./... | grep -v sample)
VERSION = $(shell grep 'version =' version.go | sed -E 's/.*"(.+)"$$/\1/')

all: build

deps:
	go get -d -v ./...

build: deps
	go build

test: deps
	go test $(TEST) -timeout=3s -parallel=4
	go vet $(TEST)
	go test $(TEST) -race

version:
	@echo $(VERSION)

.PTHONY: all deps build test version
