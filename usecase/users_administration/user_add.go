package users_administration

import (
	"errors"

	"github.com/google/uuid"
	"github.com/private-project-pp/user-rpc-service/entity"
	"github.com/private-project-pp/user-rpc-service/shared/constant"
	"github.com/private-project-pp/user-rpc-service/usecase/users_administration/responses"
)

func (s userAdminService) UserAdd(fullname, email, phoneNumber string) (out responses.UserAddResponse, err error) {
	out.Message = constant.FAILED
	existingUser, err := s.userRepo.GetExistingUsers(email, phoneNumber)
	if err != nil {

	}

	if existingUser.SecureId != "" && existingUser.Status != constant.DELETED {
		err = errors.New("user exists")
		return out, err
	}

	newUser := entity.Users{
		SecureId: uuid.NewString(),
	}

	err = s.userRepo.SaveOrUpdateUser(newUser)
	if err != nil {
		return out, err
	}
	out.SecureId = newUser.SecureId
	out.Message = constant.SUCCESS

	return out, nil
}
