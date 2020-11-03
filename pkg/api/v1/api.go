package v1

import (
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"

	"github.com/syntaqx/buildkit/pkg/api/v1/restapi"
	"github.com/syntaqx/buildkit/pkg/api/v1/restapi/operations"
)

type API struct {
	Handler http.Handler
}

func New() (*API, error) {
	spec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		return nil, err
	}

	api := operations.NewBuildKitAPI(spec)

	api.Middleware = func(b middleware.Builder) http.Handler {
		return middleware.Spec("", nil, api.Context().RoutesHandler(b))
	}

	return &API{
		Handler: api.Serve(nil),
	}, nil
}
