package postgre

import (
	"github.com/private-project-pp/user-rpc-service/domain"
	"github.com/private-project-pp/user-rpc-service/entity"
	"gorm.io/gorm"
)

type authInfo struct {
	db *gorm.DB
}

func SetupAuthInformationRepo(db *gorm.DB) domain.AuthInformation {
	return &authInfo{
		db: db,
	}
}

func (r authInfo) GetUserAuthList() (out []entity.AuthInformation, err error) {
	err = r.db.Model(&entity.AuthInformation{}).Scan(&out).Error
	if err == gorm.ErrRecordNotFound {
		return out, nil
	}

	if err != nil {
		return out, err
	}
	return out, nil
}
