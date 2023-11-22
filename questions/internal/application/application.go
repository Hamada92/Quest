package application

import (
	"context"

	"github.com/Hamada92/Quest/questions/internal/domain"
)

type (
	CreateQuestion struct {
		ID   string
		Body string
	}

	App interface {
		CreateQuestion(ctx context.Context, create CreateQuestion) error
	}

	Application struct {
		questions domain.QuestionRepository
	}
)

var _ App = (*Application)(nil)

func New(questions domain.QuestionRepository) *Application {
	return &Application{questions: questions}
}

func (a *Application) CreateQuestion(ctx context.Context, create CreateQuestion) error {
	question, err := domain.CreateQuestion(create.ID, create.Body)
	if err != nil {
		return err
	}
	return a.questions.Save(ctx, question)
}
