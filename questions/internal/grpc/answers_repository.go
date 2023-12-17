package grpc

import (
	"context"

	"github.com/Hamada92/Quest/answers/answerspb"
	"github.com/Hamada92/Quest/questions/internal/domain"
	"google.golang.org/grpc"
)

type AnswersRepository struct {
	client answerspb.AnswersServiceClient
}

var _ domain.AnswersRepository = (*AnswersRepository)(nil)

func NewAnswersRepository(conn *grpc.ClientConn) AnswersRepository {
	return AnswersRepository{client: answerspb.NewAnswersServiceClient(conn)}
}

func (a AnswersRepository) GetAnswers(ctx context.Context, questionID string) ([]*domain.Answer, error) {
	resp, err := a.client.GetAnswers(ctx, &answerspb.GetAnswersRequest{QuestionId: questionID})

	if err != nil {
		return nil, err
	}

	return a.answersToDomain(resp.Answers), nil
}

func (a *AnswersRepository) answersToDomain(answers []*answerspb.Answer) []*domain.Answer {
	var res []*domain.Answer

	for _, a := range answers {
		an := &domain.Answer{ID: a.Id, Body: a.Body, QuestionID: a.QuestionId}

		res = append(res, an)
	}

	return res
}
