.DEFAULT_GOAL := goapp

.PHONY: all
all: clean goapp

.PHONY: goapp
goapp:
	mkdir -p bin
	go build -o bin ./...

.PHONY: client
client:
	mkdir -p bin
	go build -o bin/client ./cmd/client

.PHONY: clean
clean:
	go clean
	rm -f bin/*
