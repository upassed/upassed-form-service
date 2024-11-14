package form

import (
	"github.com/google/uuid"
	domain "github.com/upassed/upassed-form-service/internal/repository/model"
	business "github.com/upassed/upassed-form-service/internal/service/model"
)

func ConvertToDomainForm(form *business.Form) *domain.Form {
	return &domain.Form{
		ID:               form.ID,
		Name:             form.Name,
		TeacherUsername:  form.TeacherUsername,
		Questions:        convertToDomainQuestions(form.Questions, form.ID),
		Description:      form.Description,
		TestingBeginDate: form.TestingBeginDate,
		TestingEndDate:   form.TestingEndDate,
		CreatedAt:        form.CreatedAt,
	}
}

func convertToDomainQuestions(questions []*business.Question, formID uuid.UUID) []*domain.Question {
	domainQuestions := make([]*domain.Question, 0, len(questions))
	for _, question := range questions {
		domainQuestions = append(domainQuestions, &domain.Question{
			ID:      question.ID,
			FormID:  formID,
			Text:    question.Text,
			Answers: convertToDomainAnswers(question.Answers, question.ID),
		})
	}

	return domainQuestions
}

func convertToDomainAnswers(answers []*business.Answer, questionID uuid.UUID) []*domain.Answer {
	domainAnswers := make([]*domain.Answer, 0, len(answers))
	for _, answer := range answers {
		domainAnswers = append(domainAnswers, &domain.Answer{
			ID:         answer.ID,
			QuestionID: questionID,
			Text:       answer.Text,
			IsCorrect:  answer.IsCorrect,
		})
	}

	return domainAnswers
}
