package main

import (
	"database/sql"

	"github.com/Hamada92/Quest/config"
	"github.com/Hamada92/Quest/internal/monolith"
	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type app struct {
	cfg     config.AppConfig
	db      *sql.DB
	logger  zerolog.Logger
	modules []monolith.Module
	mux     *chi.Mux
}
