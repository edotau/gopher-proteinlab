build:
	docker build . -t gopher-proteinlab

clean:
	go fmt ./...
	golangci-lint run ./... --fix
	yamlfmt .

install:
	go mod download && go mod verify
	go install ./...
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.61.0
	go install github.com/google/yamlfmt/cmd/yamlfmt@v0.13.0

lint:
	golangci-lint run ./...
	yamlfmt -dry .

list:
	@grep '^[^#[:space:]].*:' Makefile

test:
	go test -covermode=atomic ./...
