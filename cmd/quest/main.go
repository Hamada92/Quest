package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/Hamada92/Quest/answers"
	"github.com/Hamada92/Quest/internal/config"
	"github.com/Hamada92/Quest/internal/monolith"
	"github.com/Hamada92/Quest/questions"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"

	_ "github.com/jackc/pgx/v4/stdlib"

	"github.com/Hamada92/Quest/internal/waiter"
	"github.com/go-chi/chi/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal(err)
	}
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

	mono.rpc = initRpc()
	mono.mux = initMux()
	mono.waiter = waiter.New(waiter.CatchSignals())

	mono.modules = []monolith.Module{
		&questions.Module{},
		&answers.Module{},
	}

	if err := mono.StartUpModules(); err != nil {
		return err
	}

	fmt.Println("started mallbots application")
	defer fmt.Println("stopped mallbots application")

	mono.waiter.Add(
		mono.waitForWeb,
		mono.waitForRPC,
	)
	return mono.waiter.Wait()
}

func initRpc() *grpc.Server {
	server := grpc.NewServer(
		grpc.UnaryInterceptor(auth.UnaryServerInterceptor(authFn)),
	)
	reflection.Register(server)

	return server
}

func initMux() *chi.Mux {
	return chi.NewMux()
}

func authFn(ctx context.Context) (context.Context, error) {
	token, err := auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	// TODO: This is example only, perform proper Oauth/OIDC verification!
	if token != "yol" {
		return nil, status.Error(codes.Unauthenticated, "invalid auth token")
	}
	// NOTE: You can also pass the token in the context for further interceptors or gRPC service code.
	return context.WithValue(ctx, "authenticated", true), nil
}
