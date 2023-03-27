# Makefile for building binaries for multiple platforms

BINARY_NAME=seligman
VERSION=1.0.0

build:
	GOOS=linux GOARCH=amd64 go build -o build/$(VERSION)/linux/amd64/$(BINARY_NAME)
	GOOS=linux GOARCH=arm64 go build -o build/$(VERSION)/linux/arm64/$(BINARY_NAME)
	GOOS=darwin GOARCH=amd64 go build -o build/$(VERSION)/darwin/amd64/$(BINARY_NAME)
	GOOS=darwin GOARCH=arm64 go build -o build/$(VERSION)/darwin/arm64/$(BINARY_NAME)
	GOOS=windows GOARCH=amd64 go build -o build/$(VERSION)/windows/amd64/$(BINARY_NAME)

clean:
	rm -rf build

.PHONY: build clean


