package core

import (
	"github.com/joho/godotenv"
	"log"
	parameters "ornicode-connect/config"
)

func LoadEnv() {
	var envFileName = ".env"
	if parameters.AppPlatform == "debug" {
		envFileName = ".env.local"
	}
	err := godotenv.Load(envFileName)
	if err != nil {
		log.Fatal(err)
	}
}
