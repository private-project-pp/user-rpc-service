package users_administration

import (
	"github.com/private-project-pp/pos-general-lib/stacktrace"
	model "github.com/private-project-pp/pos-grpc-contract/model/user_service"
)

func (s userAdmin) UserList() (out []*model.UserData, err error) {
	userData, err := s.userRepo.GetAllUsers()
	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	for _, data := range userData {
		out = append(out, &model.UserData{
			SecureId: data.Id,
			FullName: data.Fullname,
		})
	}

	return out, nil
}
