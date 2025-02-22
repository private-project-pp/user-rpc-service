package authentication

import (
	"github.com/private-project-pp/user-rpc-service/domain"
	"github.com/private-project-pp/user-rpc-service/repository/redis"
)

type authService struct {
	authRepo  domain.UsersAuthInformation
	redisRepo redis.Redis
}

func SetupAuthService(authRepo domain.UsersAuthInformation, redisRepo redis.Redis) Authentication {
	return &authService{
		authRepo:  authRepo,
		redisRepo: redisRepo,
	}
}
