all: install test lint

build:
	docker build . -t gonomics

clean:
	go fmt ./...
	golangci-lint run ./... --fix

install:
	go mod download && go mod verify
	go install ./...
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0

lint:
	golangci-lint run ./...

list:
	@grep '^[^#[:space:]].*:' Makefile

test:
	go test -covermode=atomic ./...
