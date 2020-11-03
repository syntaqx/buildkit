package v1

import (
	"net/http"

	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"

	"github.com/syntaqx/buildkit/pkg/api/v1/handler"
	"github.com/syntaqx/buildkit/pkg/api/v1/restapi"
	"github.com/syntaqx/buildkit/pkg/api/v1/restapi/operations"
)

//go:generate swagger generate server --target . --spec ../../../api/openapi-spec/v1.yml --exclude-main --regenerate-configureapi

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

	usersvc := handler.UserService{}

	usersvc.ConfigureHandlers(api)

	return &API{
		Handler: api.Serve(nil),
	}, nil
}
