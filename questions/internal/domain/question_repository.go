package domain

import "context"

type QuestionRepository interface {
	Find(ctx context.Context, questionID string) (*Question, error)
	Save(ctx context.Context, q *Question) error
	// Update(ctx context.Context, q *Question) error
	List(ctx context.Context, n int, limit int) (*QList, error)
}
