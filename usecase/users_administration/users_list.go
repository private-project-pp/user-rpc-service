package users_administration

import (
	"context"

	"github.com/private-project-pp/user-rpc-service/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s userAdminService) UsersList(context.Context, *emptypb.Empty) (out *model.UsersListResponse, err error) {
	return out, nil
}
