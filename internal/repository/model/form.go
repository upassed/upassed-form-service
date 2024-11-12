package domain

import (
	"github.com/google/uuid"
)

type Form struct {
	ID              uuid.UUID
	Name            string
	TeacherUsername string
	Questions       []*Question `gorm:"foreignKey:FormID;references:ID"`
}

func (Form) TableName() string {
	return "form"
}

type Question struct {
	ID      uuid.UUID
	FormID  uuid.UUID
	Text    string
	Answers []*Answer `gorm:"foreignKey:QuestionID;references:ID"`
}

func (Question) TableName() string {
	return "question"
}

type Answer struct {
	ID         uuid.UUID
	QuestionID uuid.UUID
	Text       string
	IsCorrect  bool
}

func (Answer) TableName() string {
	return "answer"
}
