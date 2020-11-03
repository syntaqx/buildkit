SHELL := bash
PATH  := $(shell go env GOPATH)/bin:$(PATH)

ifndef VERBOSE
.SILENT:
endif

.PHONY: sync
sync:
	go mod download

.PHONY: generate
generate: tools
	go generate ./pkg/api/v1/...

.PHONY: tools
tools:
	go list -f '{{range .Imports}}{{.}} {{end}}' ./tools/tools.go | xargs go install

.PHONY: clean
clean:
	go clean -i ./...
	rm -rf pkg/api/v1/models pkg/api/v1/restapi
