package model

import "github.com/private-project-pp/user-rpc-service/internal/model/entity"

type UsersAuthInformation interface {
	GetUserByCredential(email, username, password string) (out entity.UsersAuthInformation, err error)
	GetUserAuthList() (out []entity.UsersAuthInformation, err error)
	SaveUserAuthInformation(in entity.UsersAuthInformation) (err error)
	GetLastLoginByUserId(userId string) (out entity.UserAuditTrail, err error)
}
