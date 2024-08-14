package authentication

import (
	"github.com/private-project-pp/user-rpc-service/domain"
	"github.com/private-project-pp/user-rpc-service/model"
)

type authService struct {
	authRepo domain.AuthInformation
	model.UnimplementedAuthenticationServiceServer
}

func SetupAuthService(authRepo domain.AuthInformation) model.AuthenticationServiceServer {
	return &authService{
		authRepo: authRepo,
	}
}
