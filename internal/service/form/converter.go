package form

import (
	"github.com/google/uuid"
	domain "github.com/upassed/upassed-form-service/internal/repository/model"
	business "github.com/upassed/upassed-form-service/internal/service/model"
	"time"
)

func ConvertToDomainForm(businessForm *business.Form) *domain.Form {
	return &domain.Form{
		ID:                       businessForm.ID,
		Name:                     businessForm.Name,
		TeacherUsername:          businessForm.TeacherUsername,
		Questions:                convertToDomainQuestions(businessForm.Questions, businessForm.ID),
		Description:              businessForm.Description,
		TestingBeginDate:         businessForm.TestingBeginDate,
		TestingDurationInSeconds: int64(businessForm.TestingDuration.Seconds()),
		TestingEndDate:           businessForm.TestingEndDate,
		CreatedAt:                businessForm.CreatedAt,
	}
}

func convertToDomainQuestions(businessQuestions []*business.Question, formID uuid.UUID) []*domain.Question {
	domainQuestions := make([]*domain.Question, 0, len(businessQuestions))
	for _, question := range businessQuestions {
		domainQuestions = append(domainQuestions, &domain.Question{
			ID:      question.ID,
			FormID:  formID,
			Text:    question.Text,
			Answers: convertToDomainAnswers(question.Answers, question.ID),
		})
	}

	return domainQuestions
}

func convertToDomainAnswers(businessAnswers []*business.Answer, questionID uuid.UUID) []*domain.Answer {
	domainAnswers := make([]*domain.Answer, 0, len(businessAnswers))
	for _, answer := range businessAnswers {
		domainAnswers = append(domainAnswers, &domain.Answer{
			ID:         answer.ID,
			QuestionID: questionID,
			Text:       answer.Text,
			IsCorrect:  answer.IsCorrect,
		})
	}

	return domainAnswers
}

func ConvertToBusinessForm(domainForm *domain.Form) *business.Form {
	return &business.Form{
		ID:               domainForm.ID,
		Name:             domainForm.Name,
		TeacherUsername:  domainForm.TeacherUsername,
		Questions:        convertToBusinessQuestions(domainForm.Questions),
		Description:      domainForm.Description,
		TestingBeginDate: domainForm.TestingBeginDate,
		TestingEndDate:   domainForm.TestingEndDate,
		TestingDuration:  time.Duration(domainForm.TestingDurationInSeconds) * time.Second,
		CreatedAt:        domainForm.CreatedAt,
	}
}

func ConvertToBusinessForms(domainForms []*domain.Form) []*business.Form {
	businessForms := make([]*business.Form, 0, len(domainForms))
	for _, domainForm := range domainForms {
		businessForms = append(businessForms, ConvertToBusinessForm(domainForm))
	}

	return businessForms
}

func convertToBusinessQuestions(domainQuestions []*domain.Question) []*business.Question {
	businessQuestions := make([]*business.Question, 0, len(domainQuestions))
	for _, question := range domainQuestions {
		businessQuestions = append(businessQuestions, &business.Question{
			ID:      question.ID,
			Text:    question.Text,
			Answers: convertToBusinessAnswers(question.Answers),
		})
	}

	return businessQuestions
}

func convertToBusinessAnswers(domainAnswers []*domain.Answer) []*business.Answer {
	businessAnswers := make([]*business.Answer, 0, len(domainAnswers))
	for _, answer := range domainAnswers {
		businessAnswers = append(businessAnswers, &business.Answer{
			ID:        answer.ID,
			Text:      answer.Text,
			IsCorrect: answer.IsCorrect,
		})
	}

	return businessAnswers
}
