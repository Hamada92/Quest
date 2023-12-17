package answers

import (
	"context"

	"github.com/Hamada92/Quest/answers/internal/application"
	"github.com/Hamada92/Quest/answers/internal/grpc"
	"github.com/Hamada92/Quest/answers/internal/postgres"
	"github.com/Hamada92/Quest/answers/internal/rest"
	"github.com/Hamada92/Quest/internal/monolith"
)

type Module struct{}

func (m Module) StartUp(ctx context.Context, mono monolith.Monolith) error {
	answers := postgres.NewAnswersRepository(mono.DB(), "answers.answers")

	app := application.New(answers)

	grpc.RegisterServer(app, mono.RPC())

	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	return nil

}
