package main

import (
	"database/sql"

	"github.com/Hamada92/Quest/internal/config"
	"github.com/Hamada92/Quest/internal/monolith"
	"github.com/Hamada92/Quest/questions"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	run()
}

func run() error {
	cfg, err := config.InitConfig()
	if err != nil {
		return err
	}
	mono := app{cfg: cfg}

	mono.db, err = sql.Open("pgx", cfg.PG.Conn)
	if err != nil {
		return err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			return
		}
	}(mono.db)

	mono.modules = []monolith.Module{
		&questions.Module{},
	}

	mono.rpc = initRpc()

	if err := mono.StartUpModules(); err != nil {
		return err
	}

	return nil

}

func initRpc() *grpc.Server {
	server := grpc.NewServer()
	reflection.Register(server)

	return server
}
