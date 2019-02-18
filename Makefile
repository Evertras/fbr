BINARY_NAME=fbr

all: test build

clean:
	rm -f $(BINARY_NAME)
	rm -f lib/static/build.go
	rm -f front/lib.wasm

test: lib/static/build.go
	go test -v ./lib/...

build: lib/static/build.go
	CG_ENABLED=0 go build -o $(BINARY_NAME) -v ./cmd/$(BINARY_NAME)/main.go

build-wasm: 
	GOARCH=wasm GOOS=js go build -o front/lib.wasm cmd/wasm/main.go

bench:
	go test -v -benchmem -bench . ./lib/...

run-dev: lib/static/build.go
	go run -race ./cmd/$(BINARY_NAME)/main.go -d -t 3

docker: lib/static/build.go
	docker build --rm -t evertras/fbr .

# These are not files, so always run them when asked to
.PHONY: all clean test build build-wasm bench run-dev

# Actual files/directories that must be generated
lib/static/build.go: front/lib.wasm
	go generate ./lib/static/

front/lib.wasm:
	GOARCH=wasm GOOS=js go build -o front/lib.wasm cmd/wasm/main.go
