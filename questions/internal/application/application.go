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

	GetQuestion struct {
		ID string
	}

	App interface {
		CreateQuestion(ctx context.Context, create CreateQuestion) error
		GetQuestion(ctx context.Context, get GetQuestion) (*domain.Question, error)
		GetAnswers(ctx context.Context, get GetQuestion) ([]*domain.Answer, error)
	}

	Application struct {
		questions domain.QuestionRepository
		answers   domain.AnswersRepository
	}
)

var _ App = (*Application)(nil)

func New(questions domain.QuestionRepository, answers domain.AnswersRepository) *Application {
	return &Application{questions: questions, answers: answers}
}

func (a *Application) CreateQuestion(ctx context.Context, create CreateQuestion) error {
	question, err := domain.CreateQuestion(create.ID, create.Body)
	if err != nil {
		return err
	}
	return a.questions.Save(ctx, question)
}

func (a *Application) GetQuestion(ctx context.Context, get GetQuestion) (*domain.Question, error) {
	return a.questions.Find(ctx, get.ID)
}

func (a *Application) GetAnswers(ctx context.Context, get GetQuestion) ([]*domain.Answer, error) {
	return a.answers.GetAnswers(ctx, get.ID)
}
