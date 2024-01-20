.SILENT:

.PHONY: build
build:
	go build -o ./build/xor ./cmd/xor/main.go

.PHONY: test
test:
	go test -v ./...
