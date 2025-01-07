package users_administration

import (
	"github.com/private-project-pp/user-rpc-service/domain"
)

type userAdmin struct {
	authRepo domain.UsersAuthInformation
	userRepo domain.Users
}

func SetupUserAdministration(
	authRepo domain.UsersAuthInformation,
	userRepo domain.Users,
) UsersAdministration {
	return &userAdmin{
		authRepo: authRepo,
		userRepo: userRepo,
	}
}
