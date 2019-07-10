.DEFAULT_GOAL := all

.PHONY: all
all: test run

.PHONY: test
test:
	go test ./... -cover

.PHONY: run
run:
	go run main.go
