package form

import (
	business "github.com/upassed/upassed-form-service/internal/service/model"
	"github.com/upassed/upassed-form-service/pkg/client"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertToFindByIdResponse(form *business.Form) *client.FormFindByIDResponse {
	return &client.FormFindByIDResponse{
		Form: &client.FormDTO{
			Id:               form.ID.String(),
			Name:             form.Name,
			TeacherUsername:  form.TeacherUsername,
			Description:      form.Description,
			TestingBeginDate: timestamppb.New(form.TestingBeginDate),
			TestingEndDate:   timestamppb.New(form.TestingEndDate),
			CreatedAt:        timestamppb.New(form.CreatedAt),
			Questions:        convertToClientQuestions(form.Questions),
		},
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
