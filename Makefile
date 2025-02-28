GOROOT := $(shell go env GOROOT)

.PHONY: build-web
build-web:
	@rm -rf ./dist
	@env GOOS=js GOARCH=wasm go build -o ./dist/pongo.wasm github.com/cauesmelo/pongo
	@cp $(GOROOT)/lib/wasm/wasm_exec.js ./dist
	@cp ./html/* ./dist


.PHONY: serve
serve:
	cd dist && python3 -m http.server 8080

.PHONY: watch
watch:
	@echo "Starting http server on port 8080..."
	@python3 -m http.server 8080 -d dist & \
	PID=$$!; \
	echo "HTTP server started with PID $$PID"; \
	trap "echo 'Killing HTTP server...'; kill $$PID" EXIT; \
	echo "Watching for changes in main.go, entity, and html..."; \
	fswatch -o main.go entity html | while read num; do \
		echo "Changes detected in code. Rebuilding..."; \
		make build-web; \
		echo "Dist files updated."; \
	done