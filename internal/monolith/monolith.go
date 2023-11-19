package monolith

import (
	"context"
	"database/sql"

	"github.com/Hamada92/Quest/config"
)

type Monolith interface {
	Config() config.AppConfig
	DB() *sql.DB
}

type Module interface {
	StartUp(context.Context, Monolith) error
}
