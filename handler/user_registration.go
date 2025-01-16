package handler

import (
	"context"

	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
)

func (u userService) UserRegistration(ctx context.Context, in *model.UserRegistrationRequest) (out *model.UserRegistrationResponse, err error) {
	return out, nil
}
