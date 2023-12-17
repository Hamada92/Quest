package questions

import (
	"context"

	"github.com/Hamada92/Quest/internal/monolith"
	"github.com/Hamada92/Quest/questions/internal/application"
	"github.com/Hamada92/Quest/questions/internal/grpc"
	"github.com/Hamada92/Quest/questions/internal/postgres"
	"github.com/Hamada92/Quest/questions/internal/rest"
)

type Module struct{}

func (m Module) StartUp(ctx context.Context, mono monolith.Monolith) error {
	conn, err := grpc.Dial(ctx, mono.Config().Rpc.Address())
	if err != nil {
		return err
	}

	questions := postgres.NewQuestionRepository(mono.DB(), `questions.questions`)
	answers := grpc.NewAnswersRepository(conn)
	var app application.App
	app = application.New(questions, answers)

	grpc.RegisterServer(app, mono.RPC())

	if err := rest.RegisterGateway(ctx, mono.Mux(), mono.Config().Rpc.Address()); err != nil {
		return err
	}
	return nil
}
