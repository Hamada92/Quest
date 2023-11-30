package main

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"net/http"
	"time"

	"github.com/Hamada92/Quest/internal/config"
	"github.com/Hamada92/Quest/internal/monolith"
	"github.com/Hamada92/Quest/internal/waiter"
	"github.com/go-chi/chi/v5"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

type app struct {
	cfg     config.AppConfig
	db      *sql.DB
	mux     *chi.Mux
	modules []monolith.Module
	rpc     *grpc.Server
	waiter  waiter.Waiter
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

func (a *app) waitForWeb(ctx context.Context) error {
	WebServer := &http.Server{Addr: a.cfg.Web.Address(),
		Handler: a.mux,
	}

	g, gCtx := errgroup.WithContext(ctx)

	g.Go(func() error {
		fmt.Println("web server starting")
		defer fmt.Println("web server shutting down")
		if err := WebServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			return err
		}
		return nil
	})

	g.Go(func() error {
		<-gCtx.Done()
		fmt.Println("web server shutting down")
		ctx, cancel := context.WithTimeout(ctx, a.cfg.ShutdownTimeout)
		defer cancel()
		if err := WebServer.Shutdown(ctx); err != nil {
			return err
		}
		return nil
	})

	return g.Wait()
}

func (a *app) waitForRPC(ctx context.Context) error {
	listener, err := net.Listen("tcp", a.cfg.Rpc.Address())
	if err != nil {
		return err
	}

	group, gCtx := errgroup.WithContext(ctx)
	group.Go(func() error {
		fmt.Println("rpc server started")
		defer fmt.Println("rpc server shutdown")
		if err := a.RPC().Serve(listener); err != nil && err != grpc.ErrServerStopped {
			return err
		}
		return nil
	})
	group.Go(func() error {
		<-gCtx.Done()
		fmt.Println("rpc server to be shutdown")
		stopped := make(chan struct{})
		go func() {
			a.RPC().GracefulStop()
			close(stopped)
		}()
		timeout := time.NewTimer(a.cfg.ShutdownTimeout)
		select {
		case <-timeout.C:
			// Force it to stop
			a.RPC().Stop()
			return fmt.Errorf("rpc server failed to stop gracefully")
		case <-stopped:
			return nil
		}
	})

	return group.Wait()
}
