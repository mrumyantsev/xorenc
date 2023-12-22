.PHONY: build
build:
	@go build -o ./build/xorenc ./cmd/xorenc/main.go

.PHONY: run/demo
run/demo:
	@go run ./cmd/xorenc-demo/main.go
