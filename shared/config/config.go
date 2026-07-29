package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadConfig() {
	appEnv := os.Getenv("APP_ENV")

	if appEnv == "" {
		appEnv = "development"
	}

	var err error

	switch appEnv {
	case "development":
		err = godotenv.Load(".env.development")
	case "production":
		err = godotenv.Load(".env")
	default:
		err = godotenv.Load()
	}

	if err != nil {
		log.Printf("No env file loaded (%s). Using system environment variables.\n", appEnv)
	}
}
