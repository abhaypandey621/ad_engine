# Makefile for Targeting Engine

.PHONY: all build run test lint docker clean help

BINARY=bin/server
PKG=./...

all: build

build:
	go mod tidy
	go build -o $(BINARY) cmd/server/main.go

run: build
	./$(BINARY)

test:
	go test -v $(PKG)

lint:
	@golangci-lint run || echo 'Install golangci-lint for full linting.'

docker:
	docker build -t targeting-engine .

clean:
	rm -rf bin/

help:
	@echo "Targets:"
	@echo "  build   - Build the server binary"
	@echo "  run     - Build and run the server"
	@echo "  test    - Run all tests"
	@echo "  lint    - Run linter (golangci-lint)"
	@echo "  docker  - Build Docker image"
	@echo "  clean   - Remove build artifacts" 