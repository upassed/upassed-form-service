package business

import "github.com/google/uuid"

type Form struct {
	ID        uuid.UUID
	Name      string
	Questions []*Question
}

type Question struct {
	ID      uuid.UUID
	Text    string
	Answers []*Answer
}

type Answer struct {
	ID        uuid.UUID
	Text      string
	IsCorrect bool
}

type FormCreateResponse struct {
	CreatedFormID uuid.UUID
}
