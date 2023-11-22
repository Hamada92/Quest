package domain

import "fmt"

type Question struct {
	ID   string
	Body string
}

type QList []Question

func CreateQuestion(id string, body string) (*Question, error) {
	if body == "" {
		return nil, fmt.Errorf("Question body must not be empty")
	}

	return &Question{ID: id, Body: body}, nil
}
