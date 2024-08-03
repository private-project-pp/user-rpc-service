package user

import "github.com/private-project-pp/user-rpc-service/domain"

type user struct {
	userRepo domain.User
}

func SetupUserService(
	userRepo domain.User,
) User {
	return &user{
		userRepo: userRepo,
	}
}
