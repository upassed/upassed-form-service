package domain

import "github.com/google/uuid"

type Form struct {
	ID        uuid.UUID
	Name      string
	Questions []*Question
}

func (Form) TableName() string {
	return "form"
}

type Question struct {
	ID      uuid.UUID
	Text    string
	Answers []*Answer
}

func (Question) TableName() string {
	return "question"
}

type Answer struct {
	ID        uuid.UUID
	Text      string
	IsCorrect bool
}

func (Answer) TableName() string {
	return "answer"
}
