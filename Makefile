.PHONY: build
build:
	@go build -o ./build/xore ./cmd/xore/main.go

.PHONY: run/demo
run/demo:
	@go run ./cmd/xore-demo/main.go
