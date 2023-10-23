package model

import (
	"gorm.io/gorm"
)

type UserPlatform struct {
	gorm.Model
	Name string
}
