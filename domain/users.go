package domain

import "github.com/private-project-pp/user-rpc-service/entity"

type Users interface {
	GetExistingUsers(email, phoneNumber string) (out entity.Users, err error)
	CreateUser(in entity.Users) (err error)
	GetAllUsers() (out []entity.Users, err error)
	// GetUserAuthList() (out []entity.AuthInformation, err error)
}
