package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Hamada92/Quest/answers/internal/domain"
)

type AnswersRepository struct {
	db        *sql.DB
	tableName string
}

func NewAnswersRepository(db *sql.DB, tableName string) *AnswersRepository {
	return &AnswersRepository{db: db, tableName: tableName}
}
func (r *AnswersRepository) Save(ctx context.Context, answer *domain.Answer) error {
	query := "INSERT INTO %s (id, body, question_id) VALUES ($1, $2, $3)"

	res, err := r.db.ExecContext(ctx, r.table(query), answer.ID, answer.Body, answer.QuestionId)

	if err != nil {
		return err
	}

	rows, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rows != 1 {
		return fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}

	return nil

}

func (r *AnswersRepository) GetAnswers(ctx context.Context, questionId string) (domain.Answers, error) {
	query := "select * from %s where question_id = $1"

	rows, err := r.db.QueryContext(ctx, r.table(query), questionId)

	if err != nil {
		return nil, err
	}

	answers := domain.Answers{}

	for rows.Next() {
		var answer domain.Answer
		err = rows.Scan(&answer.ID, &answer.Body, &answer.QuestionId)

		if err != nil {
			return nil, err
		}

		answers = append(answers, answer)
	}

	if closeErr := rows.Close(); closeErr != nil {
		return nil, closeErr
	}

	// Check for row scan error.
	if err != nil {
		return nil, err
	}

	return answers, nil

}

func (r *AnswersRepository) table(query string) string {
	return fmt.Sprintf(query, r.tableName)
}
