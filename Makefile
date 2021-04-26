LINT_ARGS=


all: build

.PHONY: build

build:
	build/make_build.sh

test:
	go test -mod vendor ./...

lint:
	golangci-lint -c .golangci.yml $(LINT_ARGS) run
