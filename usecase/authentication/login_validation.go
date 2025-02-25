package authentication

import (
	"github.com/private-project-pp/pos-general-lib/stacktrace"
	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
	"github.com/private-project-pp/user-rpc-service/shared/constant"
)

func (s authService) ValidateLogin(in *model.LoginValidationRequest) (out *model.LoginValidationResponse, err error) {
	var (
		isValidated bool
		token       string
		message     string
	)
	userData, err := s.authRepo.GetUserByCredential(in.Email, in.Username, in.Password)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	if userData.Status == constant.SUSPENDED {
		message = "Akun ditangguhkan"
	}

	lastLogin, err := s.authRepo.GetLastLoginByUserId(userData.UserId)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	token, err = "", nil
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.BAD_PROCESSING, err.Error())
	}

	key, err := s.redisRepo.SaveUserLoginCache(token)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	out = &model.LoginValidationResponse{
		IsValidated: isValidated,
		Token:       key,
		LastLogin:   lastLogin.CreatedAt.Format(""),
		Message:     message,
	}
	return out, nil
}
