package entity

import "time"

type User struct {
	SecureId    string    `gorm:"column:secure_id"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	CreatedBy   string    `gorm:"column:created_by"`
	Email       string    `gorm:"column:email"`
	Password    string    `gorm:"column:password"`
	Username    string    `gorm:"column:username"`
	PhoneNumber string    `gorm:"column:phone_number"`
	Status      string    `gorm:"column:status"`
}

func (User) TableName() string {
	return "user"
}
