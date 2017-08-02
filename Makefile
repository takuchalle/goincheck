TEST ?= $(shell go list ./... | grep -v sample)
VERSION = $(shell grep 'version =' version.go | sed -E 's/.*"(.+)"$$/\1/')

all:

deps:
	go get -d -v ./...

test: deps
	go test $(TEST) -timeout=3s -parallel=4
	go vet $(TEST)
	go test $(TEST) -race

version:
	@echo $(VERSION)

.PTHONY: all deps test version
