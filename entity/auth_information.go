package entity

import (
	"time"

	"github.com/private-project-pp/user-rpc-service/shared/constant"
)

type AuthInformation struct {
	SecureId    string    `gorm:"column:id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	CreatedBy   string    `gorm:"column:created_by"`
	UserId      string    `gorm:"column:user_id"`
	Email       string    `gorm:"column:email"`
	Password    string    `gorm:"column:password"`
	Username    string    `gorm:"column:username"`
	PhoneNumber string    `gorm:"column:phone_number"`
	Status      string    `gorm:"column:status"`
}

func (AuthInformation) TableName() string {
	return constant.AuthInformation
}
