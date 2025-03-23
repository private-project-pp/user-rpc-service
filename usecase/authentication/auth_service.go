package authentication

import (
	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
)

type Authentication interface {
	ValidateUser(in *model.LoginValidationRequest) (out *model.LoginValidationResponse, err error)
}
