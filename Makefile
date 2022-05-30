.PHONY: all
all: vet build

.PHONY: build
build:
	go build ./cmd/rds-download-log-file

.PHONY: vet
vet:
	go vet ./...

.PHONY: clean
clean:
	rm -f rds-download-log-file rds-download-log-file.exe
