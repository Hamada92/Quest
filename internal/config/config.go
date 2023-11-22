package config

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type PGConfig struct {
	Conn string
}

type AppConfig struct {
	Environment string
	PG          PGConfig
}

func InitConfig() (a AppConfig, err error) {
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	err = envconfig.Process("", &a)
	return
}
