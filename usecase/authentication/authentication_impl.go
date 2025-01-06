package authentication

import (
	"github.com/private-project-pp/user-rpc-service/domain"
)

type authService struct {
	authRepo domain.UsersAuthInformation
}

func SetupAuthService(authRepo domain.UsersAuthInformation) Authentication {
	return &authService{
		authRepo: authRepo,
	}
}
