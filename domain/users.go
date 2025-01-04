package domain

import "github.com/private-project-pp/user-rpc-service/entity"

type Users interface {
	GetAllUsers() (out []entity.Users, err error)
	// GetUserAuthList() (out []entity.AuthInformation, err error)
}
