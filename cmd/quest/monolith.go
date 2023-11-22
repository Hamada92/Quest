package main

import (
	"context"
	"database/sql"

	"github.com/Hamada92/Quest/internal/config"
	"github.com/Hamada92/Quest/internal/monolith"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
)

type app struct {
	cfg     config.AppConfig
	db      *sql.DB
	mux     *chi.Mux
	modules []monolith.Module
	rpc     *grpc.Server
}

func (a *app) StartUpModules() error {
	for _, m := range a.modules {
		if err := m.StartUp(context.TODO(), a); err != nil {
			return err
		}
	}
	return nil
}

func (a *app) DB() *sql.DB {
	return a.db
}

func (a *app) Config() config.AppConfig {
	return a.cfg
}

func (a *app) RPC() *grpc.Server {
	return a.rpc
}
