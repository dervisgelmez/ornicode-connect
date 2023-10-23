package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Username    string `gorm:"unique"`
	FirstName   string
	LastName    string
	Password    string
	Email       *string
	PhoneNumber *string
	Enabled     bool
	LockedAt    *time.Time
}

func (User) TableName() string {
	return "t_user"
}
