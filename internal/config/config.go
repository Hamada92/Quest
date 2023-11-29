package config

import (
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type PGConfig struct {
	Conn string
}

type AppConfig struct {
	Environment     string
	PG              PGConfig
	ShutdownTimeout time.Duration `default:"30s"`
}

func InitConfig() (a AppConfig, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = envconfig.Process("", &a)
	return
}
