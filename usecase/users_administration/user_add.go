package users_administration

import (
	"errors"

	"github.com/google/uuid"
	"github.com/private-project-pp/pos-general-lib/stacktrace"
	"github.com/private-project-pp/user-rpc-service/entity"
	"github.com/private-project-pp/user-rpc-service/shared/constant"
	"github.com/private-project-pp/user-rpc-service/usecase/users_administration/request"
	"github.com/private-project-pp/user-rpc-service/usecase/users_administration/responses"
)

func (s userAdmin) UserAdd(in request.UserAddRequest) (out responses.UserAddResponse, err error) {
	out.Message = constant.FAILED
	if in.Fullname == "" {
		err = errors.New("Nama lengkap wajib diisi")
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	if in.PhoneNumber == "" {
		err = errors.New("Nomor telepon wajib diisi")
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
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
