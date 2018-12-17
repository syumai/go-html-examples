.PHONY: build
build:
	GOOS=js GOARCH=wasm go build -o minimum/index.wasm ./minimum
