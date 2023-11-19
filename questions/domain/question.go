package domain

import "fmt"

type Question struct {
	QuestionID string
	Body       string
}

type QList []Question

func CreateQuestion(body string) (*Question, error) {
	if body == "" {
		return nil, fmt.Errorf("Question body must not be empty")
	}

	return &Question{Body: body}, nil
}
