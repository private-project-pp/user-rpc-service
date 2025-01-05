package users_administration

import (
	"github.com/private-project-pp/user-rpc-service/domain"
	"github.com/private-project-pp/user-rpc-service/model"
)

type userAdminService struct {
	authRepo domain.UsersAuthInformation
	userRepo domain.Users
	model.UnimplementedUsersAdministrationServiceServer
}

func SetupUserAdministration(
	authRepo domain.UsersAuthInformation,
	userRepo domain.Users,
) model.UsersAdministrationServiceServer {
	return &userAdminService{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}
