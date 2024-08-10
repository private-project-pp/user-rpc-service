package user

import "github.com/private-project-pp/user-rpc-service/domain"

type userService struct {
	authRepo domain.AuthInformation
	UnimplementedUserServiceServer
}

func SetupUserService(authRepo domain.AuthInformation) UserServiceServer {
	return &userService{
		authRepo: authRepo,
	}
}
