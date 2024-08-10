package domain

import "github.com/private-project-pp/user-rpc-service/entity"

type AuthInformation interface {
	GetUserAuthList() (out []entity.AuthInformation, err error)
}
