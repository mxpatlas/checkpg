all: build

.PHONY: build

build:
	build/make_build.sh

test:
	go test -mod vendor ./...

lint:

