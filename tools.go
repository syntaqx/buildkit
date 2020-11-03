// +build tools

package buildkit

// nolint

// Manage tool dependencies via go.mod.
// https://github.com/golang/go/wiki/Modules#how-can-i-track-tool-dependencies-for-a-module
// https://github.com/golang/go/issues/25922
//
//     go list -f '{{range .Imports}}{{.}} {{end}}' ./tools.go | xargs go install
//
import (
	_ "github.com/go-swagger/go-swagger/cmd/swagger"
	_ "github.com/go-openapi/errors"
	_ "github.com/go-openapi/loads"
	_ "github.com/go-openapi/runtime"
	_ "github.com/go-openapi/spec"
	_ "github.com/go-openapi/strfmt"
	_ "github.com/go-openapi/swag"
	_ "github.com/go-openapi/validate"
	_ "github.com/jessevdk/go-flags, or"
	_ "github.com/spf13/pflags"
)
