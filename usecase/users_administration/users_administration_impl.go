package users_administration

import "github.com/private-project-pp/user-rpc-service/model"

type userAdminService struct {
	model.UnimplementedUsersAdministrationServiceServer
}
