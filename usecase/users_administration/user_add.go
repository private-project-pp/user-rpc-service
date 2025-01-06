package users_administration

import (
	"context"

	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
)

func (s userAdminService) UserAdd(ctx context.Context, in *model.UserAddRequest) (out *model.UserAddResponse, err error) {
	return out, nil
}
