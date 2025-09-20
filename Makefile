# Makefile for precise-calc

.PHONY: build test clean format vet

# Build the application
build:
	go build -o bin/precise-calc cmd/precise-calc/main.go

# Run all tests
test:
	go test ./...

# Run tests with coverage
test-coverage:
	go test -cover ./...

# Clean build artifacts
clean:
	rm -rf bin/

# Format code
format:
	go fmt ./...

# Run go vet
vet:
	go vet ./...

# Run all checks
check: format vet test

# Build and install
install: build
	cp bin/precise-calc /usr/local/bin/

.DEFAULT_GOAL := build