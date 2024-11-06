package util

import (
	"github.com/brianvoe/gofakeit/v7"
	event "github.com/upassed/upassed-form-service/internal/messanging/model"
)

func RandomEventFormCreateRequest() *event.FormCreateRequest {
	questionsCount := gofakeit.IntRange(30, 50)
	questions := make([]*event.Question, 0, questionsCount)
	for i := 0; i < questionsCount; i++ {
		questions = append(questions, RandomEventQuestion())
	}

	return &event.FormCreateRequest{
		Name:      gofakeit.Slogan(),
		Questions: questions,
	}
}

func RandomEventQuestion() *event.Question {
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
		IsCorrect: gofakeit.Bool(),
	}
}
