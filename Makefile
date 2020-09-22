.PHONY: build
build:
	go build -v ./httpd/main.go
	./main

.DEFAULT_GOAL := build