package postgre

import (
	"github.com/private-project-pp/pos-general-lib/stacktrace"
	"github.com/private-project-pp/user-rpc-service/domain"
	"github.com/private-project-pp/user-rpc-service/entity"
	"gorm.io/gorm"
)

type users struct {
	db *gorm.DB
}

func SetupUsersRepo(db *gorm.DB) domain.Users {
	return &users{
		db: db,
	}
}

func (r users) GetUserDataById(id string) (out entity.Users, err error) {
	err = r.db.Where("id = ?", id).Scan(&out).Error
	if err != gorm.ErrRecordNotFound {
		return out, stacktrace.Cascade(err, stacktrace.DATA_NOT_FOUND, err.Error())
	}

	if err != nil {
		return out, stacktrace.Cascade(err, stacktrace.INTERNAL_SERVER_ERROR, err.Error())
	}

	return out, nil
}

func (r users) GetExistingUsers(email, phoneNumber string) (out entity.Users, err error) {
	return out, nil
}

func (r users) CreateUser(in entity.Users) (err error) {
	err = r.db.Create(&in).Error
	if err != nil {
		return err
	}
	return nil
}

func (r users) GetAllUsers() (out []entity.Users, err error) {
	err = r.db.Model(&entity.Users{}).Scan(&out).Error
	if err == gorm.ErrRecordNotFound {
		return out, nil
	}
	if err != nil {
		return out, err
	}
	return out, nil
}
