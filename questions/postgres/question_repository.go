package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Hamada92/Quest/questions/domain"
)

type QuestionRepository struct {
	tableName string
	db        *sql.DB
}

func NewQuestionRepository(db *sql.DB, tableName string) QuestionRepository {
	return QuestionRepository{
		tableName: tableName,
		db:        db,
	}
}

func (r QuestionRepository) Find(ctx context.Context, questionID string) (*domain.Question, error) {
	const query = "SELECT body, FROM %s WHERE id = $1 LIMIT 1"
	question := &domain.Question{QuestionID: questionID}

	err := r.db.QueryRowContext(ctx, r.table(query), questionID).Scan(&question.Body)

	return question, err
}

func (r QuestionRepository) List(ctx context.Context, n int, limit int) (*domain.QList, error) {
	const query = "SELECT *, FROM %s WHERE id < $1 ORDER BY id DESC LIMIT $2"

	rows, err := r.db.QueryContext(ctx, r.table(query), n, limit)
	if err != nil {
		return nil, err
	}

	qlist := domain.QList{}

	for rows.Next() {
		var q domain.Question
		err := rows.Scan(&q)
		if err != nil {
			break
		}
		qlist = append(qlist, q)
	}

	if closeErr := rows.Close(); closeErr != nil {
		return nil, closeErr
	}

	// Check for row scan error.
	if err != nil {
		return nil, err
	}

	return &qlist, err
}

func (r QuestionRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}