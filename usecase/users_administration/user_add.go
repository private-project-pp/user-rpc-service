package users_administration

import (
	"errors"

	"github.com/google/uuid"
	"github.com/private-project-pp/user-rpc-service/entity"
	"github.com/private-project-pp/user-rpc-service/shared/constant"
	"github.com/private-project-pp/user-rpc-service/usecase/users_administration/responses"
)

func (s userAdmin) UserAdd(fullname, email, phoneNumber string) (out responses.UserAddResponse, err error) {
	out.Message = constant.FAILED
	if fullname == "" {
		err = errors.New("fullname wajib diisi")
		return out, err
	}

	if email == "" {
		err = errors.New("email wajib diisi")
		return out, err
	}

	if phoneNumber == "" {
		err = errors.New("phoneNumber wajib diisi")
		return out, err
	}
	existingUser, err := s.userRepo.GetExistingUsers(email, phoneNumber)
	if err != nil {

	}

	if existingUser.Id != "" && existingUser.Status != constant.DELETED {
		err = errors.New("user exists")
		return out, err
	}

	newUser := entity.Users{
		Id:          uuid.New().String(),
		Fullname:    fullname,
		Email:       email,
		PhoneNumber: phoneNumber,
		Status:      constant.ACTIVE,
	}

	err = s.userRepo.CreateUser(newUser)
	if err != nil {
		return out, err
	}
	out.SecureId = newUser.Id
	out.Message = constant.SUCCESS

	return out, nil
}
