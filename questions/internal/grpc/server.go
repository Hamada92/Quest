package grpc

import (
	"context"

	"github.com/Hamada92/Quest/questions/internal/application"
	"github.com/Hamada92/Quest/questions/internal/domain"
	"github.com/Hamada92/Quest/questions/questionspb"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

type server struct {
	app application.App
	questionspb.UnimplementedQuestionsServiceServer
}

var _ questionspb.QuestionsServiceServer = (*server)(nil)

func RegisterServer(app application.App, registrar grpc.ServiceRegistrar) error {
	questionspb.RegisterQuestionsServiceServer(registrar, server{app: app})
	return nil
}

func (s server) CreateQuestion(ctx context.Context, request *questionspb.CreateQuestionRequest) (*questionspb.CreateQuestionResponse, error) {
	id := uuid.New().String()
	body := request.GetBody()

	err := s.app.CreateQuestion(ctx, application.CreateQuestion{ID: id, Body: body})

	return &questionspb.CreateQuestionResponse{Id: id, Body: body}, err
}

func (s server) GetQuestion(ctx context.Context, request *questionspb.GetQuestionRequest) (*questionspb.GetQuestionResponse, error) {
	id := request.GetId()

	question, err := s.app.GetQuestion(ctx, application.GetQuestion{ID: id})

	if err != nil {
		return nil, err
	}

	answers, err := s.app.GetAnswers(ctx, application.GetQuestion{ID: id})

	return &questionspb.GetQuestionResponse{
		Question: s.qustionFromDomain(question),
		Answers:  s.answersFromDomain(answers),
	}, err
}

func (s server) qustionFromDomain(question *domain.Question) *questionspb.Question {
	return &questionspb.Question{
		Id:   question.ID,
		Body: question.Body,
	}
}

func (s server) answersFromDomain(answers []*domain.Answer) []*questionspb.Answer {
	var res []*questionspb.Answer
	for _, a := range answers {
		w := &questionspb.Answer{
			Id:         a.ID,
			Body:       a.Body,
			QuestionId: a.QuestionID,
		}
		res = append(res, w)
	}
	return res
}
