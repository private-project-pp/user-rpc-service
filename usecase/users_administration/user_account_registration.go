package users_administration

import (
	"github.com/google/uuid"
	"github.com/private-project-pp/pos-general-lib/stacktrace"
	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
	"github.com/private-project-pp/user-rpc-service/entity"
	"github.com/private-project-pp/user-rpc-service/shared/utils"
)

func (u userAdmin) UserAccountRegistration(in *model.UserRegistrationRequestPayload) (out *model.UserRegistrationResponse, err error) {

	userData, err := u.userRepo.GetUserDataById(in.GetUserId())
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	hashedPassword, err := utils.GeneratePasswordHash(in.Password, "")
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.BAD_PROCESSING, err.Error())
	}

	regUser := entity.UsersAuthInformation{
		SecureId: uuid.New().String(),
		// Username: in.Username,
		UserId:   userData.Id,
		Password: hashedPassword,
	}

	_ = regUser
	return out, nil
}
