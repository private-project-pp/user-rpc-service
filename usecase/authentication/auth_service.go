package authentication

import (
	"context"

	"github.com/private-project-pp/user-rpc-service/model"
)

type AuthService interface {
	ValidateLogin(context.Context, *model.LoginValidationRequest) (out *model.LoginValidationResponse, err error)
}
