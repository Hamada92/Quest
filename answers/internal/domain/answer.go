package domain

import "fmt"

type Answer struct {
	ID         string
	QuestionId string
	Body       string
}

type Answers []Answer

func CreateAnswer(id string, body string, questionId string) (*Answer, error) {
	if body == "" {
		return nil, fmt.Errorf("body can't be empty")
	}

	if questionId == "" {
		return nil, fmt.Errorf("question id can't be empty")
	}

	answer := &Answer{ID: id, QuestionId: questionId, Body: body}

	return answer, nil
}
