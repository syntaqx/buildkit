PATH := $(shell go env GOPATH)/bin:$(PATH)

mod-download:
	echo $(go env GOPATH)
	go mod download

install-tools:
	go list -f '{{range .Imports}}{{.}} {{end}}' ./tools/tools.go | xargs go install

generate: install-tools
	go generate ./...
