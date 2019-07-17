.DEFAULT_GOAL := all

.PHONY: all
all: format test run

.PHONY: format
format:
	go fmt

.PHONY: test
test:
	go test ./... -cover

.PHONY: run
run:
	go run main.go url.go
