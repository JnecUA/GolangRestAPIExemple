.PHONY: build
build:
	go build ./cmd/apiserver
	./apiserver

.DEFAULT_GOAL := build