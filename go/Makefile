

APP_NAME=sample
BINARY_NAME=../wasmtest/src/wasmdemo/$(APP_NAME).wasm

build:
	@echo "Building..."
	GOOS=js GOARCH=wasm go build  -o $(BINARY_NAME) main.go
	@echo "Copying wasm_exec.js"
	cp $(shell go env GOROOT)/misc/wasm/wasm_exec.js ../wasmtest/public/

tinygo:
	@echo "Building with tinygo ..."
	tinygo build  -o $(BINARY_NAME) -target=wasm main.go
	@echo "Copying wasm_exec.js from tinygo"
	cp $(shell tinygo env TINYGOROOT)/targets/wasm_exec.js ../wasmtest/public/
