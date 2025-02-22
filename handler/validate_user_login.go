package handler

import (
	"context"

	"github.com/private-project-pp/pos-general-lib/stacktrace"
	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
)

func (rpc userService) ValidateUserLogin(ctx context.Context, in *model.LoginValidationRequest) (out *model.LoginValidationResponse, err error) {

	out, err = rpc.auth.ValidateLogin(in)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	return out, nil
}
