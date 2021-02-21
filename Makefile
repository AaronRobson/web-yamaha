.DEFAULT_GOAL := all

.PHONY: all
all: format test

.PHONY: install
install:
	go get ${gobuild_args} ./...

.PHONY: format
format:
	go fmt

.PHONY: check
check:
	golint

.PHONY: test
test:
	go test ./... -cover

.PHONY: run
run:
	go run main.go funcs.go url.go messages.go
