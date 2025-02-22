package users_administration

import (
	"errors"

	"github.com/google/uuid"
	"github.com/private-project-pp/pos-general-lib/stacktrace"
	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
	"github.com/private-project-pp/user-rpc-service/entity"
	"github.com/private-project-pp/user-rpc-service/shared/constant"
	"github.com/private-project-pp/user-rpc-service/usecase/users_administration/responses"
)

func (s userAdmin) UserAdd(in *model.UserAddRequestPayload) (out responses.UserAddResponse, err error) {
	out.Message = constant.FAILED
	if in.Fullname == "" {
		err = errors.New("nama lengkap wajib diisi")
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error())
	}

	if in.PhoneNumber == "" {
		err = errors.New("nomor telepon wajib diisi")
		return out, stacktrace.Cascade(err, stacktrace.INVALID_INPUT, err.Error())
	}

	newUser := entity.Users{
		Id:          uuid.New().String(),
		Fullname:    in.Fullname,
		PhoneNumber: in.PhoneNumber,
		Status:      constant.ACTIVE,
	}

	err = s.userRepo.CreateUser(newUser)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	out.SecureId = newUser.Id
	out.Message = constant.SUCCESS

	return out, nil
}
