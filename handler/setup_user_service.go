package handler

import (
	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
	"github.com/private-project-pp/user-rpc-service/usecase/authentication"
	"github.com/private-project-pp/user-rpc-service/usecase/users_administration"
)

type userService struct {
	auth    authentication.Authentication
	userAdm users_administration.UsersAdministration
	model.UnimplementedUserServiceServer
}

func SetupUserService(
	auth authentication.Authentication,
	userAdm users_administration.UsersAdministration,
) model.UserServiceServer {
	return &userService{
		auth:    auth,
		userAdm: userAdm,
	}
}
