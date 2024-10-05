SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

SRC := $(shell find . -type f -name "*.go" -exec echo {} \;)

test:
	go test -v=false ./...

coverage:
	go test -v=false ./... -coverprofile=cover.out
	go tool cover -func=cover.out

clean: 
	rm -f cover.out
