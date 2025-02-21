package handler

import (
	"context"

	"github.com/private-project-pp/pos-general-lib/stacktrace"
	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
	"github.com/private-project-pp/user-rpc-service/usecase/users_administration/request"
)

func (u userService) UserRegistration(ctx context.Context, in *model.UserRegistrationRequest) (out *model.UserRegistrationResponse, err error) {
	req, err := request.NewUserAccountRegistrationReq(in.GetPayload())
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error())
	}

	out, err = u.userAdm.UserAccountRegistration(req)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	return out, nil
}
