package util

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	event "github.com/upassed/upassed-form-service/internal/messanging/model"
	domain "github.com/upassed/upassed-form-service/internal/repository/model"
	business "github.com/upassed/upassed-form-service/internal/service/model"
)

func RandomEventFormCreateRequest() *event.FormCreateRequest {
	questionsCount := gofakeit.IntRange(30, 50)
	questions := make([]*event.Question, 0, questionsCount)
	for i := 0; i < questionsCount; i++ {
		questions = append(questions, randomEventQuestion())
	}

	return &event.FormCreateRequest{
		Name:      gofakeit.Slogan(),
		Questions: questions,
	}
}

func randomEventQuestion() *event.Question {
	answersCount := gofakeit.IntRange(3, 8)
	answers := make([]*event.Answer, 0, answersCount)
	for i := 0; i < answersCount; i++ {
		answers = append(answers, RandomEventAnswer())
	}

	return &event.Question{
		Text:    gofakeit.Slogan(),
		Answers: answers,
	}
}

func RandomEventAnswer() *event.Answer {
	return &event.Answer{
		Text:      gofakeit.Slogan(),
		IsCorrect: true,
	}
}

func RandomBusinessForm() *business.Form {
	questionsCount := gofakeit.IntRange(30, 50)
	questions := make([]*business.Question, 0, questionsCount)
	for i := 0; i < questionsCount; i++ {
		questions = append(questions, randomBusinessQuestion())
	}

	return &business.Form{
		ID:        uuid.New(),
		Name:      gofakeit.Slogan(),
		Questions: questions,
	}
}

func randomBusinessQuestion() *business.Question {
	answersCount := gofakeit.IntRange(3, 8)
	answers := make([]*business.Answer, 0, answersCount)
	for i := 0; i < answersCount; i++ {
		answers = append(answers, randomBusinessAnswer())
	}

	return &business.Question{
		ID:      uuid.New(),
		Text:    gofakeit.Slogan(),
		Answers: answers,
	}
}

func randomBusinessAnswer() *business.Answer {
	return &business.Answer{
		ID:        uuid.New(),
		Text:      gofakeit.Slogan(),
		IsCorrect: true,
	}
}

func RandomDomainForm() *domain.Form {
	domainFormID := uuid.New()

	questionsCount := gofakeit.IntRange(30, 50)
	questions := make([]*domain.Question, 0, questionsCount)
	for i := 0; i < questionsCount; i++ {
		questions = append(questions, randomDomainQuestion(domainFormID))
	}

	return &domain.Form{
		ID:              domainFormID,
		Name:            gofakeit.Slogan(),
		TeacherUsername: gofakeit.Username(),
		Questions:       questions,
	}
}

func randomDomainQuestion(domainFormID uuid.UUID) *domain.Question {
	domainQuestionID := uuid.New()

	answersCount := gofakeit.IntRange(3, 8)
	answers := make([]*domain.Answer, 0, answersCount)
	for i := 0; i < answersCount; i++ {
		answers = append(answers, randomDomainAnswer(domainQuestionID))
	}

	return &domain.Question{
		ID:      domainQuestionID,
		FormID:  domainFormID,
		Text:    gofakeit.Slogan(),
		Answers: answers,
	}
}

func randomDomainAnswer(domainQuestionID uuid.UUID) *domain.Answer {
	return &domain.Answer{
		ID:         uuid.New(),
		QuestionID: domainQuestionID,
		Text:       gofakeit.Slogan(),
		IsCorrect:  true,
	}
}
