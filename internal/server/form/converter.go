package form

import (
	business "github.com/upassed/upassed-form-service/internal/service/model"
	"github.com/upassed/upassed-form-service/pkg/client"
	"google.golang.org/protobuf/types/known/durationpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertToFindByIdResponse(form *business.Form) *client.FormFindByIDResponse {
	return &client.FormFindByIDResponse{
		Form: convertToClientForm(form),
	}
}

func convertToClientForm(businessForm *business.Form) *client.FormDTO {
	return &client.FormDTO{
		Id:               businessForm.ID.String(),
		Name:             businessForm.Name,
		TeacherUsername:  businessForm.TeacherUsername,
		Description:      businessForm.Description,
		TestingBeginDate: timestamppb.New(businessForm.TestingBeginDate),
		TestingEndDate:   timestamppb.New(businessForm.TestingEndDate),
		TestingDuration:  durationpb.New(businessForm.TestingDuration),
		CreatedAt:        timestamppb.New(businessForm.CreatedAt),
		Questions:        convertToClientQuestions(businessForm.Questions),
	}
}

func convertToClientQuestions(questions []*business.Question) []*client.QuestionDTO {
	clientQuestions := make([]*client.QuestionDTO, 0, len(questions))
	for _, question := range questions {
		clientQuestions = append(clientQuestions, &client.QuestionDTO{
			Id:      question.ID.String(),
			Text:    question.Text,
			Answers: convertToClientAnswers(question.Answers),
		})
	}

	return clientQuestions
}

func convertToClientAnswers(answers []*business.Answer) []*client.AnswerDTO {
	clientAnswers := make([]*client.AnswerDTO, 0, len(answers))
	for _, answer := range answers {
		clientAnswers = append(clientAnswers, &client.AnswerDTO{
			Id:        answer.ID.String(),
			Text:      answer.Text,
			IsCorrect: answer.IsCorrect,
		})
	}

	return clientAnswers
}

func ConvertToFindByTeacherUsernameResponse(businessForms []*business.Form) *client.FormFindByTeacherUsernameResponse {
	clientForms := make([]*client.FormDTO, 0, len(businessForms))
	for _, businessForm := range businessForms {
		clientForms = append(clientForms, convertToClientForm(businessForm))
	}

	return &client.FormFindByTeacherUsernameResponse{
		FoundForms: clientForms,
	}
}
