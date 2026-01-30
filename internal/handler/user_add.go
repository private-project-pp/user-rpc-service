package handler

import (
	"context"

	"github.com/private-project-pp/pos-general-lib/stacktrace"
	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
)

func (s userService) UserAdd(ctx context.Context, in *model.UserAddRequest) (out *model.UserAddResponse, err error) {

	resp, err := s.userAdm.UserAdd(in.GetPayload())
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	out = &model.UserAddResponse{
		UserId:  resp.SecureId,
		Message: resp.Message,
	}
	return out, nil
}
