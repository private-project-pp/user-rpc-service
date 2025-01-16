package users_administration

import (
	"github.com/private-project-pp/user-rpc-service/usecase/users_administration/request"
	"github.com/private-project-pp/user-rpc-service/usecase/users_administration/responses"
)

type UsersAdministration interface {
	UserAdd(fullname, email, phoneNumber string) (out responses.UserAddResponse, err error)
	UserAccountRegistrationReq(in request.UserAccountRegistrationReq) (out responses.UserAccountRegistration, err error)
}
