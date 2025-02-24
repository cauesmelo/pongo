GOROOT := $(shell go env GOROOT)

.PHONY: build-web
build-web:
	@rm -rf ./dist
	@env GOOS=js GOARCH=wasm go build -o ./dist/pongo.wasm github.com/cauesmelo/pongo
	@cp $(GOROOT)/misc/wasm/wasm_exec.js ./dist
	@cp ./html/* ./dist


.PHONY: serve
serve:
	cd dist && python3 -m http.server 8080
