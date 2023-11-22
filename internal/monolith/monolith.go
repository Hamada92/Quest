package monolith

import (
	"context"
	"database/sql"

	"github.com/Hamada92/Quest/internal/config"
	"google.golang.org/grpc"
)

type Monolith interface {
	Config() config.AppConfig
	DB() *sql.DB
	RPC() *grpc.Server
}

type Module interface {
	StartUp(context.Context, Monolith) error
}
