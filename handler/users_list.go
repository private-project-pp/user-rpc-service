package handler

import (
	"context"

	"github.com/private-project-pp/pos-general-lib/stacktrace"
	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s userService) UsersList(context.Context, *emptypb.Empty) (out *model.UsersListResponse, err error) {
	result, err := s.userAdm.UserList()
	if err != nil {
		return nil, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	out = &model.UsersListResponse{
		Users: result,
	}
	return out, nil
}
