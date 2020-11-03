package buildkit

import (
	_ "github.com/go-openapi/errors"
	_ "github.com/go-openapi/loads"
	_ "github.com/go-openapi/runtime"
	_ "github.com/go-openapi/spec"
	_ "github.com/go-openapi/strfmt"
	_ "github.com/go-openapi/swag"
	_ "github.com/go-openapi/validate"
	_ "github.com/jessevdk/go-flags"
)

//go:generate swagger generate server --target ./pkg/api/v1 --spec ./api/openapi-spec/v1.yml --exclude-main --regenerate-configureapi
