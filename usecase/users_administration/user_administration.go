package users_administration

import "github.com/private-project-pp/user-rpc-service/usecase/users_administration/responses"

type UsersAdministration interface {
	UserAdd(fullname, email, phoneNumber string) (out responses.UserAddResponse, err error)
}
