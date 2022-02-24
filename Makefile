.DEFAULT_GOAL := all

.PHONY: all
all: format test

.PHONY: install
install:
	go get ${gobuild_args} ./...

.PHONY: docker-build
docker-build:
	docker build .

.PHONY: docker-run
docker-run: docker-build
	docker run \
		-p 8080:8080/tcp
		--read-only
		-t web-yamaha
		something

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
