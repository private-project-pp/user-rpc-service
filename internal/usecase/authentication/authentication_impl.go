package authentication

import (
	"github.com/private-project-pp/user-rpc-service/internal/model"
	"github.com/private-project-pp/user-rpc-service/internal/repository/redis"
)

type authService struct {
	authRepo  model.UsersAuthInformation
	redisRepo redis.Redis
}

func SetupAuthService(authRepo model.UsersAuthInformation, redisRepo redis.Redis) Authentication {
	return &authService{
		authRepo:  authRepo,
		redisRepo: redisRepo,
	}
}
