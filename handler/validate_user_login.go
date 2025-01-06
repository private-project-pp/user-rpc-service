package handler

import (
	"context"

	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
)

func (rpc userService) ValidateUserLogin(ctx context.Context, in *model.LoginValidationRequest) (out *model.LoginValidationResponse, err error) {

	return out, nil
}
