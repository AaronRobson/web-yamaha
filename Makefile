.DEFAULT_GOAL := all

.PHONY: all
all: format test

.PHONY: format
format:
	go fmt

.PHONY: test
test:
	go test ./... -cover

.PHONY: run
run:
	go run main.go funcs.go url.go
