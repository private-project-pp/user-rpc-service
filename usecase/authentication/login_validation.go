package authentication

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/private-project-pp/pos-general-lib/metadata"
	constant_lib "github.com/private-project-pp/pos-general-lib/shared/constant"
	"github.com/private-project-pp/pos-general-lib/stacktrace"
	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
	"github.com/private-project-pp/user-rpc-service/shared/constant"
	"github.com/private-project-pp/user-rpc-service/shared/utils"
)

func (s authService) ValidateUser(in *model.LoginValidationRequest) (out *model.LoginValidationResponse, err error) {
	var (
		isValidated bool
		message     string
	)
	userData, err := s.authRepo.GetUserByCredential(in.GetEmail(), in.GetUsername(), in.GetPassword())
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	if userData.Status == constant.SUSPENDED {
		message = "akun ditangguhkan"
		err = errors.New(message)
		return out, stacktrace.CascadeWithClientMessage(err, stacktrace.BAD_PROCESSING, err.Error())
	}

	lastLogin, err := s.authRepo.GetLastLoginByUserId(userData.UserId)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	sessionInfo := map[string]string{
		metadata.XAuthUserId:   userData.UserId,
		metadata.XAuthUserRole: userData.Role,
		"X-CreatedTime":        time.Now().UTC().Format(constant_lib.YYMMDDHiS),
	}

	generatedKey, err := utils.GenerateUserToken(constant_lib.SHA256, sessionInfo)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	splitSession := strings.Split(generatedKey, ".")
	session := fmt.Sprintf("%s.%s", splitSession[0], splitSession[1])

	err = s.redisRepo.SaveUserLoginCache(session, splitSession[2])
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	out = &model.LoginValidationResponse{
		IsValidated: isValidated,
		Token:       splitSession[2],
		LastLogin:   lastLogin.CreatedAt.Format(constant_lib.YYMMDDHiS),
		Message:     message,
	}
	return out, nil
}
