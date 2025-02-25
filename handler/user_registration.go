package handler

import (
	"context"

	"github.com/private-project-pp/pos-general-lib/stacktrace"
	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
)

func (u userService) UserRegistration(ctx context.Context, in *model.UserRegistrationRequest) (out *model.UserRegistrationResponse, err error) {

	out, err = u.userAdm.UserAccountRegistration(in.GetPayload())
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	return out, nil
}
