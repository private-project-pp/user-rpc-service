package handler

import (
	"context"

	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s userService) UsersList(context.Context, *emptypb.Empty) (out *model.UsersListResponse, err error) {
	return out, nil
}
