package authentication

import (
	"context"

	"github.com/private-project-pp/user-rpc-service/model"
)

func (s authService) ValidateLogin(context.Context, *model.LoginValidationRequest) (out *model.LoginValidationResponse, err error) {
	out = &model.LoginValidationResponse{
		IsValidated: true,
		Token:       "Ini tokennya",
	}
	return out, nil
}
