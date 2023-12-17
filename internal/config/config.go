package config

import (
	"fmt"
	"log"
	"time"

	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
)

type PGConfig struct {
	Conn string
}

type WebConfig struct {
	Host string `default:"0.0.0.0"`
	Port string `default:":8080"`
}

type RpcConfig struct {
	Host string `default:"0.0.0.0"`
	Port string `default:":8085"`
}

func (c RpcConfig) Address() string {
	return fmt.Sprintf("%s%s", c.Host, c.Port)
}

func (c WebConfig) Address() string {
	return fmt.Sprintf("%s%s", c.Host, c.Port)
}

type AppConfig struct {
	Environment     string
	PG              PGConfig
	Web             WebConfig
	Rpc             RpcConfig
	ShutdownTimeout time.Duration `default:"30s"`
	Secret          string
}

func InitConfig() (a AppConfig, err error) {
	err = godotenv.Load("./.env")
	if err != nil {
		log.Fatal("Error loading .env file", err)
	}

	err = envconfig.Process("", &a)
	return
}
