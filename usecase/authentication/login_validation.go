package authentication

import (
	"context"

	"github.com/private-project-pp/user-rpc-service/model"
	"github.com/private-project-pp/user-rpc-service/shared/constant"
)

func (s authService) ValidateLogin(ctx context.Context, in *model.LoginValidationRequest) (out *model.LoginValidationResponse, err error) {
	var (
		isValidated bool
		token       string
		message     string
	)
	userData, err := s.authRepo.GetUserByCredential(in.Email, in.Username, in.Password)
	if err != nil {
		return out, err
	}
	if userData.Status == constant.SUSPENDED {
		message = "Akun ditangguhkan"
	}

	out = &model.LoginValidationResponse{
		IsValidated: isValidated,
		Token:       token,
		LastLogin:   "",
		Message:     message,
	}
	return out, nil
}
