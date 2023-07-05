
SHELL := /bin/bash

.PHONY: all build test deps deps-cleancache

GOCMD=go
BUILD_DIR=build
BINARY_DIR=$(BUILD_DIR)/bin
CODE_COVERAGE=code-coverage

all: test build

${BINARY_DIR}:
	mkdir -p $(BINARY_DIR)

build-linux: ${BINARY_DIR} ## Build executable for Linux ( arm64)
	env GOOS=linux GOARCH=arm64 $(GOCMD) build 
# swag: ## Generate swagger docs
# 	swag init --parseDependency --parseInternal --parseDepth 1 -md ./documentation -o ./docs

swag: ## Generate swagger docs
	swag init -g ./main.go -o ./docs