package form

import (
	"encoding/json"
	"github.com/google/uuid"
	event "github.com/upassed/upassed-form-service/internal/messanging/model"
	business "github.com/upassed/upassed-form-service/internal/service/model"
	"time"
)

func ConvertToFormCreateRequest(messageBody []byte) (*event.FormCreateRequest, error) {
	var request event.FormCreateRequest
	if err := json.Unmarshal(messageBody, &request); err != nil {
		return nil, err
	}

	return &request, nil
}

func ConvertToBusinessForm(request *event.FormCreateRequest, teacherUsername string) *business.Form {
	questions := make([]*business.Question, 0, len(request.Questions))
	for _, question := range request.Questions {
		questions = append(questions, convertToQuestion(question))
	}

	return &business.Form{
		ID:               uuid.New(),
		Name:             request.Name,
		TeacherUsername:  teacherUsername,
		Questions:        questions,
		Description:      request.Description,
		TestingBeginDate: request.TestingBeginDate,
		TestingEndDate:   request.TestingEndDate,
		TestingDuration:  request.TestingDuration,
		CreatedAt:        time.Now(),
	}
}

func convertToQuestion(question *event.Question) *business.Question {
	answers := make([]*business.Answer, 0, len(question.Answers))
	for _, answer := range question.Answers {
		answers = append(answers, convertToAnswer(answer))
	}

	return &business.Question{
		ID:      uuid.New(),
		Text:    question.Text,
		Answers: answers,
	}
}

func convertToAnswer(answer *event.Answer) *business.Answer {
	return &business.Answer{
		ID:        uuid.New(),
		Text:      answer.Text,
		IsCorrect: answer.IsCorrect,
	}
}
