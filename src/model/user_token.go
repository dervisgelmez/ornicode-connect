package model

import (
	"gorm.io/gorm"
	"time"
)

type UserToken struct {
	gorm.Model
	User         User
	AccessToken  string
	RefreshToken string
	UserPlatform UserPlatform
	ExpiredAt    *time.Time
}
