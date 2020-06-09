SHELL   := /bin/bash
VERSION := v1.1.0
GOOS    := $(shell go env GOOS)
GOARCH  := $(shell go env GOARCH)

.PHONY: all
all: build

.PHONY: build
build:
	go build -ldflags "-X main.version=$(VERSION)" ./cmd/rds-download-log-file

.PHONY: package
package: clean build
	gzip rds-download-log-file -c > rds-download-log-file_$(VERSION)_$(GOOS)_$(GOARCH).gz
	sha1sum rds-download-log-file_$(VERSION)_$(GOOS)_$(GOARCH).gz > rds-download-log-file_$(VERSION)_$(GOOS)_$(GOARCH).gz.sha1sum

.PHONY: clean
clean:
	rm -f rds-download-log-file
