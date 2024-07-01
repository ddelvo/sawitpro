.PHONY: clean all init

all: build/main

build/main: main.go
	@echo "Building..."
	go build -o $@ $<

init: clean
	go mod tidy
	go mod vendor

test:
	go clean -testcache
	go test -short -coverprofile coverage.out -short -v ./...
