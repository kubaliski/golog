VERSION ?= v0.1.1
EXAMPLE_PKG := ./examples/simple

# detect os
ifeq ($(OS),Windows_NT)
    RM = del /q
    RMDIR = rmdir /s /q
    MKDIR = if not exist bin mkdir bin
    GOLINT = $(shell if exist tools\golangci-lint.exe (echo tools\golangci-lint.exe) else (echo golangci-lint))
else
    RM = rm -f
    RMDIR = rm -rf
    MKDIR = mkdir -p bin
    GOLINT = $(shell if [ -f ./tools/golangci-lint ]; then echo ./tools/golangci-lint; else echo golangci-lint; fi)
endif

.PHONY: all example test cover tag build clean lint

all: test

example:
	go run $(EXAMPLE_PKG)

test:
	go test ./... -coverprofile=coverage.out

cover: test
	go tool cover -func=coverage.out

tag:
	git tag -a $(VERSION) -m "release $(VERSION)"
	git push origin $(VERSION)

build:
	$(MKDIR)
	go build -o bin/golog ./...

lint:
	$(GOLINT) run --timeout 5m ./...

clean:
	$(RM) coverage.out
	$(RMDIR) bin