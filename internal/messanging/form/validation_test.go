package form_test

import (
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/require"
	event "github.com/upassed/upassed-form-service/internal/messanging/model"
	"github.com/upassed/upassed-form-service/internal/util"
	"testing"
	"time"
)

func TestFormCreateRequestValidation_EmptyName(t *testing.T) {
	request := util.RandomEventFormCreateRequest()
	request.Name = ""

	err := request.Validate()
	require.Error(t, err)
}

func TestFormCreateRequestValidation_NameTooLong(t *testing.T) {
	request := util.RandomEventFormCreateRequest()
	request.Name = gofakeit.LoremIpsumSentence(1000)

	err := request.Validate()
	require.Error(t, err)
}

func TestFormCreateRequestValidation_EmptyQuestions(t *testing.T) {
	request := util.RandomEventFormCreateRequest()
	request.Questions = nil

	err := request.Validate()
	require.Error(t, err)
}

func TestFormCreateRequestValidation_EmptyQuestionText(t *testing.T) {
	request := util.RandomEventFormCreateRequest()
	request.Questions[0].Text = ""

	err := request.Validate()
	require.Error(t, err)
}

func TestFormCreateRequestValidation_QuestionTextTooLong(t *testing.T) {
	request := util.RandomEventFormCreateRequest()
	request.Questions[0].Text = gofakeit.LoremIpsumSentence(1000)

	err := request.Validate()
	require.Error(t, err)
}

func TestFormCreateRequestValidation_AnswersCountIsZero(t *testing.T) {
	request := util.RandomEventFormCreateRequest()
	request.Questions[0].Answers = nil

	err := request.Validate()
	require.Error(t, err)
}

func TestFormCreateRequestValidation_AnswersCountIsOne(t *testing.T) {
	request := util.RandomEventFormCreateRequest()
	request.Questions[0].Answers = []*event.Answer{util.RandomEventAnswer()}

	err := request.Validate()
	require.Error(t, err)
}

func TestFormCreateRequestValidation_EmptyAnswerText(t *testing.T) {
	request := util.RandomEventFormCreateRequest()
	request.Questions[0].Answers[0].Text = ""

	err := request.Validate()
	require.Error(t, err)
}

func TestFormCreateRequestValidation_AnswerTextTooLong(t *testing.T) {
	request := util.RandomEventFormCreateRequest()
	request.Questions[0].Answers[0].Text = gofakeit.LoremIpsumSentence(1000)

	err := request.Validate()
	require.Error(t, err)
}

func TestFormCreateRequestValidation_NoCorrectAnswerOnTheQuestion(t *testing.T) {
	request := util.RandomEventFormCreateRequest()
	for _, answer := range request.Questions[0].Answers {
		answer.IsCorrect = false
	}

	err := request.Validate()
	require.Error(t, err)
}

func TestFormCreateRequestValidation_InvalidTestingDateRange(t *testing.T) {
	request := util.RandomEventFormCreateRequest()
	request.TestingBeginDate = request.TestingEndDate.Add(1 * time.Hour)

	err := request.Validate()
	require.Error(t, err)
}

func TestFormCreateRequestValidation_HappyPath(t *testing.T) {
	request := util.RandomEventFormCreateRequest()

	err := request.Validate()
	require.NoError(t, err)
}
