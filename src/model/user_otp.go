package model

import (
	"gorm.io/gorm"
	"time"
)

type UserOtp struct {
	gorm.Model
	Otp          string
	Key          string
	Type         string
	UserPlatform UserPlatform
	ExpiredAt    *time.Time
}
