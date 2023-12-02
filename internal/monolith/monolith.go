package monolith

import (
	"context"
	"database/sql"

	"github.com/Hamada92/Quest/internal/config"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
)

type Monolith interface {
	Config() config.AppConfig
	DB() *sql.DB
	RPC() *grpc.Server
	Mux() *chi.Mux
}

type Module interface {
	StartUp(context.Context, Monolith) error
}
