package core

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"ornicode-connect/src/model"
	"os"
)

var DB *gorm.DB

func ConnectDatabase() {
	var err error
	dsn := os.Getenv("DATABASE_DSN")
	DB, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

func SyncDatabase() {
	DB.AutoMigrate(&model.User{})
}
