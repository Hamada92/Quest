package application

import (
	"context"

	"github.com/Hamada92/Quest/answers/internal/domain"
)

type CreateAnswer struct {
	Id         string
	Body       string
	QuestionId string
}

type GetAnswers struct {
	QuestionId string
}

type App interface {
	CreateAnswer(ctx context.Context, create CreateAnswer) error
	GetAnswers(ctx context.Context, get GetAnswers) (domain.Answers, error)
}

type Application struct {
	repository domain.AnswersRepository
}

var _ App = (*Application)(nil)

func New(answers domain.AnswersRepository) *Application {
	return &Application{repository: answers}
}

func (a *Application) CreateAnswer(ctx context.Context, create CreateAnswer) error {
	answer, err := domain.CreateAnswer(create.Id, create.Body, create.QuestionId)

	if err != nil {
		return err
	}

	return a.repository.Save(ctx, answer)
}

func (a *Application) GetAnswers(ctx context.Context, get GetAnswers) (domain.Answers, error) {
	return a.repository.GetAnswers(ctx, get.QuestionId)
}
