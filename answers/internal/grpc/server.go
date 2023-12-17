package grpc

import (
	"context"

	"github.com/Hamada92/Quest/answers/answerspb"
	"github.com/Hamada92/Quest/answers/internal/application"
	"github.com/Hamada92/Quest/answers/internal/domain"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	answerspb.UnimplementedAnswersServiceServer
}

var _ answerspb.AnswersServiceServer = (*server)(nil)

func RegisterServer(app application.App, registrar grpc.ServiceRegistrar) error {
	answerspb.RegisterAnswersServiceServer(registrar, server{app: app})

	return nil
}

func (s server) CreateAnswer(ctx context.Context, request *answerspb.CreateAnswerRequest) (*answerspb.CreateAnswerResponse, error) {
	id := uuid.New().String()
	body := request.GetBody()
	questionId := request.GetQuestionId()

	err := s.app.CreateAnswer(ctx, application.CreateAnswer{Id: id, Body: body, QuestionId: questionId})
	return &answerspb.CreateAnswerResponse{Id: id, Body: body, QuestionId: questionId}, err
}

func (s server) GetAnswers(ctx context.Context, request *answerspb.GetAnswersRequest) (*answerspb.GetAnswersResponse, error) {
	questionId := request.GetQuestionId()
	answers, err := s.app.GetAnswers(ctx, application.GetAnswers{QuestionId: questionId})

	if err != nil {
		return nil, err
	}

	return &answerspb.GetAnswersResponse{Answers: AnswersFromDomainAnswers(answers)}, nil
}

func AnswersFromDomainAnswers(domainAnswers domain.Answers) []*answerspb.Answer {
	answers := []*answerspb.Answer{}
	for _, a := range domainAnswers {
		a := &answerspb.Answer{Id: a.ID, Body: a.Body, QuestionId: a.QuestionId}
		answers = append(answers, a)
	}

	return answers
}
