BINARY_NAME=server

all: test build

clean:
	rm -f $(BINARY_NAME)
	rm -f lib/static/build.go
	rm -f front/lib.wasm

test: lib/static/build.go
	@echo 'NOTE: This may take a while on the first run due to OpenGL compilation...'
	go test -v ./lib/...

build: lib/static/build.go
	CGO_ENABLED=0 go build -o $(BINARY_NAME) -v ./cmd/$(BINARY_NAME)/main.go

build-wasm: 
	CGO_ENABLED=0 GOARCH=wasm GOOS=js go build -o front/lib.wasm cmd/wasm/main.go

bench:
	go test -v -benchmem -run 'xxx' -bench . ./lib/...

run-dev:
	go run -race ./cmd/$(BINARY_NAME)/main.go -d

docker: lib/static/build.go
	docker build --rm -t evertras/fbr .

# These are not files, so always run them when asked to
.PHONY: all clean test build build-wasm bench run-dev docker

# Actual files/directories that must be generated
lib/static/build.go: front/lib.wasm front/index.html front/style.css front/wasm_exec.js
	go generate ./lib/static/

front/lib.wasm:
	GOARCH=wasm GOOS=js go build -o front/lib.wasm cmd/wasm/main.go
