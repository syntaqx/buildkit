install-tools:
  go list -f '{{range .Imports}}{{.}} {{end}}' ./tools.go | xargs go install
