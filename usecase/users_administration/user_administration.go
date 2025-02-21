package users_administration

import (
	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
	"github.com/private-project-pp/user-rpc-service/usecase/users_administration/request"
	"github.com/private-project-pp/user-rpc-service/usecase/users_administration/responses"
)

type UsersAdministration interface {
	UserAdd(in request.UserAddRequest) (out responses.UserAddResponse, err error)
	UserList() (out []*model.UserData, err error)
	UserAccountRegistration(in request.UserAccountRegistrationReq) (out *model.UserRegistrationResponse, err error)
}
