package domain

import "context"

type AnswersRepository interface {
	GetAnswers(ctx context.Context, questionID string) ([]*Answer, error)
}
