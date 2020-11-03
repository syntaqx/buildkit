package handler

import (
	"fmt"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/gofrs/uuid"

	"github.com/syntaqx/buildkit/pkg/api/v1/models"
	"github.com/syntaqx/buildkit/pkg/api/v1/restapi/operations"
	"github.com/syntaqx/buildkit/pkg/api/v1/restapi/operations/user"
)

type UserService struct {
}

func (s *UserService) ConfigureHandlers(api *operations.BuildKitAPI) {
	api.UserGetUserHandler = user.GetUserHandlerFunc(s.GetUser)
}

func (s *UserService) GetUser(params user.GetUserParams) middleware.Responder {
	var login = "fixture"

	payload := &models.User{
		ID:        strfmt.UUID(uuid.Must(uuid.NewV4()).String()),
		Login:     login,
		Email:     strfmt.Email(fmt.Sprintf("%s@domain.com", login)),
		CreatedAt: strfmt.DateTime(time.Now()),
		UpdatedAt: strfmt.DateTime(time.Now()),
	}

	return user.NewGetUserOK().WithPayload(payload)
}
