package main

import (
	"database/sql"

	"github.com/Hamada92/Quest/config"
)

func main() {
	run()
}

func run() error {
	cfg, err := config.InitConfig()
	if err != nil {
		return err
	}
	monolith := app{cfg: cfg}

	monolith.db, err = sql.Open("pgx", cfg.PG.Conn)
	if err != nil {
		return err
	}

}
