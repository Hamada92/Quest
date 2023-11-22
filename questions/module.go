package questions

import (
	"context"

	"github.com/Hamada92/Quest/internal/monolith"
	"github.com/Hamada92/Quest/questions/internal/application"
	"github.com/Hamada92/Quest/questions/internal/grpc"
	"github.com/Hamada92/Quest/questions/internal/postgres"
)

type Module struct{}

func (m Module) StartUp(ctx context.Context, mono monolith.Monolith) error {
	questions := postgres.NewQuestionRepository(mono.DB(), "questions.questions")

	var app application.App
	app = application.New(questions)

	grpc.RegisterServer(app, mono.RPC())
	return nil
}
