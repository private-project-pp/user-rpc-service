package users_administration

import (
	"github.com/google/uuid"
	"github.com/private-project-pp/pos-general-lib/stacktrace"
	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
	"github.com/private-project-pp/user-rpc-service/entity"
	"github.com/private-project-pp/user-rpc-service/shared/constant"
	"github.com/private-project-pp/user-rpc-service/shared/utils"
)

func (u userAdmin) UserAccountRegistration(in *model.UserRegistrationRequestPayload) (out *model.UserRegistrationResponse, err error) {

	userData, err := u.userRepo.GetUserDataById(in.GetUserId())
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	hashedPassword, err := utils.GenerateHashedPassword(in.GetPassword())
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.BAD_PROCESSING, err.Error())
	}

	regUser := entity.UsersAuthInformation{
		SecureId:  uuid.New().String(),
		CreatedAt: utils.GetUtcTime(),
		CreatedBy: in.GetUserId(),
		UserId:    userData.Id,
		Password:  hashedPassword,
		Username:  in.GetUsername(),
		Status:    constant.ACTIVE,
	}
	err = u.authRepo.SaveUserAuthInformation(regUser)
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return out, nil
}
