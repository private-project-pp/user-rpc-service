package handler

import (
	"context"

	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
)

func (s userService) UserAdd(ctx context.Context, in *model.UserAddRequest) (out *model.UserAddResponse, err error) {

	resp, err := s.userAdm.UserAdd(in.GetFullname(), in.GetEmail(), in.GetPhoneNumber())
	if err != nil {
		return out, err
	}

	out = &model.UserAddResponse{
		UserId:  resp.SecureId,
		Message: resp.Message,
	}
	return out, nil
}
