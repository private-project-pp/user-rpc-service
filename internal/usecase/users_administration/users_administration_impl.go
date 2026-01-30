package users_administration

import (
	"github.com/private-project-pp/user-rpc-service/internal/model"
)

type userAdmin struct {
	authRepo model.UsersAuthInformation
	userRepo model.Users
}

func SetupUserAdministration(
	authRepo model.UsersAuthInformation,
	userRepo model.Users,
) UsersAdministration {
	return &userAdmin{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}
