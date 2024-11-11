package form

import (
	domain "github.com/upassed/upassed-form-service/internal/repository/model"
	business "github.com/upassed/upassed-form-service/internal/service/model"
)

func ConvertToDomainForm(form *business.Form) *domain.Form {
	return &domain.Form{
		ID:        form.ID,
		Name:      form.Name,
		Questions: convertToDomainQuestions(form.Questions),
	}
}

func convertToDomainQuestions(questions []*business.Question) []*domain.Question {
	domainQuestions := make([]*domain.Question, 0, len(questions))
	for _, question := range questions {
		domainQuestions = append(domainQuestions, &domain.Question{
			ID:      question.ID,
			Text:    question.Text,
			Answers: convertToDomainAnswers(question.Answers),
		})
	}

	return domainQuestions
}

func convertToDomainAnswers(answers []*business.Answer) []*domain.Answer {
	domainAnswers := make([]*domain.Answer, 0, len(answers))
	for _, answer := range answers {
		domainAnswers = append(domainAnswers, &domain.Answer{
			ID:        answer.ID,
			Text:      answer.Text,
			IsCorrect: answer.IsCorrect,
		})
	}

	return domainAnswers
}
