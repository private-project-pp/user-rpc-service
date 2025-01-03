package entity

import (
	"time"

	"github.com/private-project-pp/user-rpc-service/shared/constant"
)

type Users struct {
	SecureId  string     `gorm:"column:id"`
	CreatedAt time.Time  `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	Fullname  string     `gorm:"column:fullname"`
	Email     string     `gorm:"column:email"`
	Status    string     `gorm:"column:status"`
}

func (Users) TableName() string {
	return constant.UsersEntity
}
