package domain

import "context"

type AnswersRepository interface {
	Save(ctx context.Context, a *Answer) error
	GetAnswers(ctx context.Context, questionId string) (Answers, error)
}
