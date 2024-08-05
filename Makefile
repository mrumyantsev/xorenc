.DEFAULT_GOAL := build
.SILENT:

.PHONY: build
build:
	go build -o ./build/xor ./cmd/xor

.PHONY: test
test:
	go test -v ./...

.PHONY: install
install:
	cp ./build/xor /usr/local/bin
