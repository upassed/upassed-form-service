package event

import (
	"errors"
	"github.com/go-playground/validator/v10"
)

var (
	errQuestionsSliceIsEmpty       = errors.New("number of questions should be > 0")
	errInsufficientNumberOfAnswers = errors.New("number of answers should be > 1")
	errNoOneAnswerIsCorrect        = errors.New("no one answer is correct")
)

type FormCreateRequest struct {
	Name      string      `json:"name,omitempty" validate:"required,min=4,max=120"`
	Questions []*Question `json:"questions,omitempty" validate:"required,dive"`
}

type Question struct {
	Text    string    `json:"text,omitempty" validate:"required,min=4,max=250"`
	Answers []*Answer `json:"answers,omitempty" validate:"required,dive"`
}

type Answer struct {
	Text      string `json:"text,omitempty" validate:"required,min=2,max=120"`
	IsCorrect bool   `json:"is_correct,omitempty" validate:"boolean"`
}

func (request *FormCreateRequest) Validate() error {
	validate := validator.New()
	if err := validate.Struct(*request); err != nil {
		return err
	}

	if len(request.Questions) == 0 {
		return errQuestionsSliceIsEmpty
	}

	for _, question := range request.Questions {
		if len(question.Answers) <= 1 {
			return errInsufficientNumberOfAnswers
		}

		correctAnswersCount := 0
		for _, answer := range question.Answers {
			if answer.IsCorrect {
				correctAnswersCount++
			}
		}

		if correctAnswersCount == 0 {
			return errNoOneAnswerIsCorrect
		}
	}

	return nil
}
