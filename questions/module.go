package questions

import (
	"context"

	"github.com/Hamada92/Quest/internal/monolith"
	"github.com/Hamada92/Quest/questions/postgres"
)

type Module struct{}

func (m Module) StartUp(ctx context.Context, mono monolith.Monolith) error {
	questions := postgres.NewQuestionRepository(mono.DB(), "questions.Questions")
}
