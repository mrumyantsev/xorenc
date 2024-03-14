.SILENT:
.DEFAULT_GOAL := build

.PHONY: build
build:
	go build -o ./build/xor ./cmd/xor

.PHONY: test
test:
	go test -v ./...
