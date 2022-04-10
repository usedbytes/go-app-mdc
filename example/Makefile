all: server app

server:
	go build

app:
	GOARCH=wasm GOOS=js go build -o web/app.wasm

.PHONY: all server app
