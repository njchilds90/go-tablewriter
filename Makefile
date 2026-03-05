.PHONY: build test lint fmt vet bench coverage

build:
	go build -o bin/tablewriter ./...

test:
	go test -v ./...

lint:
	golangci-lint run

fmt:
	go fmt ./...

vet:
	go vet ./...

bench:
	go test -bench=. ./...

coverage:
	go test -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
