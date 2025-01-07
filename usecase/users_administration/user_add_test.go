package users_administration

import (
	"testing"

	"github.com/private-project-pp/user-rpc-service/entity"
	mocks_repository "github.com/private-project-pp/user-rpc-service/mocks/repositories"
	"github.com/private-project-pp/user-rpc-service/shared/constant"
	"github.com/private-project-pp/user-rpc-service/usecase/users_administration/responses"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestUserAdd(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	type args struct {
		userRepo *mocks_repository.MockUsers
	}

	tests := []struct {
		name    string
		args    args
		usecase func(args)
		wants   func(*testing.T, responses.UserAddResponse, error)
	}{
		{
			name: "Success condition",
			args: args{
				userRepo: mocks_repository.NewMockUsers(mockCtrl),
			},
			usecase: func(a args) {
				a.userRepo.EXPECT().GetExistingUsers("", "").Return(entity.Users{}, nil).Times(1)
				a.userRepo.EXPECT().CreateUser(gomock.AssignableToTypeOf(entity.Users{})).Return(nil).Times(1)
			},
			wants: func(t *testing.T, uar responses.UserAddResponse, err error) {
				assert.NotEqual(t, uar.SecureId, "", "Hasil harusnya tidak kosong")
				assert.Equal(t, uar.Message, constant.SUCCESS, "Response message harusnya SUCCESS")
			},
		},
		{
			name: "Failed condition [USER EXISTS]",
			args: args{
				userRepo: mocks_repository.NewMockUsers(mockCtrl),
			},
			usecase: func(a args) {
				a.userRepo.EXPECT().GetExistingUsers("", "").Return(entity.Users{Id: "1", Status: constant.ACTIVE}, nil).Times(1)
				a.userRepo.EXPECT().CreateUser(gomock.AssignableToTypeOf(entity.Users{})).Return(nil).Times(0)
			},
			wants: func(t *testing.T, uar responses.UserAddResponse, err error) {
				assert.Equal(t, uar.SecureId, "", "Hasil harusnya kosong")
				assert.Equal(t, uar.Message, constant.FAILED, "Response message harusnya FAILED")
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			testedStruct := userAdmin{userRepo: test.args.userRepo}
			test.usecase(test.args)
			response, err := testedStruct.UserAdd("", "", "")
			test.wants(t, response, err)
		})
	}
}
