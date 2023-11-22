package grpc

import (
	"context"

	"github.com/Hamada92/Quest/questions/internal/application"
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
