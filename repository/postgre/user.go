package postgre

import (
	"github.com/private-project-pp/user-rpc-service/domain"
	"gorm.io/gorm"
)

type user struct {
	db *gorm.DB
}

func SetupUserRepo(db *gorm.DB) domain.User {
	return &user{
		db: db,
	}
}
