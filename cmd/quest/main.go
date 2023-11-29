package main

import (
	"database/sql"

	"github.com/Hamada92/Quest/internal/config"
	"github.com/Hamada92/Quest/internal/monolith"
	"github.com/Hamada92/Quest/internal/waiter"
	"github.com/Hamada92/Quest/questions"
	"github.com/go-chi/chi"
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
	mono.mux = initMux()
	mono.waiter = waiter.New()
	if err := mono.StartUpModules(); err != nil {
		return err
	}

	mono.waiter.Add(
		mono.waitForWeb,
		mono.waitForRPC,
	)

}

func initRpc() *grpc.Server {
	server := grpc.NewServer()
	reflection.Register(server)

	return server
}

func initMux() *chi.Mux {
	return chi.NewMux()
}
