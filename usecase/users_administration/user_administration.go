package users_administration

import (
	"context"

	"github.com/private-project-pp/user-rpc-service/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UsersAdministration interface {
	UsersList(context.Context, *emptypb.Empty) (*model.UsersListResponse, error)
	UserAdd(ctx context.Context, in *model.UserAddRequest) (out *model.UserAddResponse, err error)
}
