package authentication

import (
	"github.com/private-project-pp/user-rpc-service/domain"
	"github.com/private-project-pp/user-rpc-service/model"
)

type authService struct {
	authRepo domain.UsersAuthInformation
	model.UnimplementedAuthenticationServiceServer
}

func SetupAuthService(authRepo domain.UsersAuthInformation) model.AuthenticationServiceServer {
	return &authService{
		authRepo: authRepo,
	}
}
