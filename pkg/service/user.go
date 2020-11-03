package service

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

func (s *UserService) ConfigureAPI(api *operations.BuildKitAPI) {
	api.UserGetUserHandler = user.GetUserHandlerFunc(s.GetUser)
}

func (s *UserService) GetUser(params user.GetUserParams) middleware.Responder {
	var username = "fixture"

	payload := &models.User{
		ID:        strfmt.UUID(uuid.Must(uuid.NewV4()).String()),
		Username:  username,
		Email:     strfmt.Email(fmt.Sprintf("%s@domain.com", username)),
		CreatedAt: strfmt.DateTime(time.Now()),
		UpdatedAt: strfmt.DateTime(time.Now()),
	}

	return user.NewGetUserOK().WithPayload(payload)
}
