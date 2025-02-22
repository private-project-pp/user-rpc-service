package authentication

import (
	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
)

type Authentication interface {
	ValidateLogin(in *model.LoginValidationRequest) (out *model.LoginValidationResponse, err error)
}
