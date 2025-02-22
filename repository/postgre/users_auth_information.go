package postgre

import (
	"github.com/private-project-pp/pos-general-lib/stacktrace"
	"github.com/private-project-pp/user-rpc-service/domain"
	"github.com/private-project-pp/user-rpc-service/entity"
	"gorm.io/gorm"
)

type authInfoRepo struct {
	db *gorm.DB
}

func SetupAuthInformationRepo(db *gorm.DB) domain.UsersAuthInformation {
	return &authInfoRepo{
		db: db,
	}
}

func (r authInfoRepo) GetUserByCredential(email, username, password string) (out entity.UsersAuthInformation, err error) {
	err = r.db.Model(&entity.UsersAuthInformation{}).
		Joins("inner join users u on u.id = users_auth_information.user_id").
		Where("(u.email = ? OR username = ?) AND password = ?", email, username, password).
		Scan(&out).Error
	if err == gorm.ErrRecordNotFound {
		return out, nil
	}

	if err != nil {
		return out, err
	}
	return out, nil
}

func (r authInfoRepo) GetUserAuthList() (out []entity.UsersAuthInformation, err error) {
	err = r.db.Model(&entity.UsersAuthInformation{}).Scan(&out).Error
	if err == gorm.ErrRecordNotFound {
		return out, nil
	}

	if err != nil {
		return out, err
	}
	return out, nil
}

func (r authInfoRepo) SaveUserAuthInformation(in entity.UsersAuthInformation) (err error) {
	err = r.db.Save(&in).Error
	if err != nil {
		return stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}
	return nil
}
