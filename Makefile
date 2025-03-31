GOOS ?= $(shell go env GOOS)
GOARCH ?= $(shell go env GOARCH)

.PHONY: build
build:
	mkdir -p dist/$(GOOS)/$(GOARCH)
	go build -tags 'netgo,disable_crypt' -buildmode=pie -trimpath -ldflags '-s' -o dist/$(GOOS)/$(GOARCH)/operator-podman -v ./cmd/operator-podman

.PHONY: clean
clean:
	rm -rf dist

.PHONY: default
default: build