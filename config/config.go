package config

import (
	"log"
	"os"

	"github.com/subosito/gotenv"
)

type Config struct {
	AppName     string
	AppPort     int
	LogLevel    string
	Environment string
	JWTSecret   string
}

func Init() *Config {
	defaultEnv := ".env"

	if err := gotenv.Load(defaultEnv); err != nil {
		log.Fatal("failed load .env")
	}

	log.SetOutput(os.Stdout)

	appConfig := &Config{
		AppName:     GetString("APP_NAME"),
		AppPort:     GetInt("APP_PORT"),
		LogLevel:    GetString("LOG_LEVEL"),
		Environment: GetString("ENVIRONMENT"),
		JWTSecret:   GetString("JWT_SECRET"),
	}

	return appConfig
}
