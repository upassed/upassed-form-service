package form_test

import (
	"encoding/json"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/upassed/upassed-form-service/internal/messanging/form"
	event "github.com/upassed/upassed-form-service/internal/messanging/model"
	business "github.com/upassed/upassed-form-service/internal/service/model"
	"github.com/upassed/upassed-form-service/internal/util"
	"testing"
	"time"
)

func TestConvertToFormCreateRequest_InvalidBytes(t *testing.T) {
	invalidBytes := make([]byte, 10)
	_, err := form.ConvertToFormCreateRequest(invalidBytes)
	require.Error(t, err)
}

func TestConvertToFormCreateRequest_ValidBytes(t *testing.T) {
	initialRequest := util.RandomEventFormCreateRequest()
	initialRequestBytes, err := json.Marshal(initialRequest)
	require.NoError(t, err)

	convertedRequest, err := form.ConvertToFormCreateRequest(initialRequestBytes)
	require.NoError(t, err)

	assert.Equal(t, initialRequest.Name, convertedRequest.Name)
	assert.Equal(t, initialRequest.Description, convertedRequest.Description)
	assert.WithinDuration(t, initialRequest.TestingBeginDate, convertedRequest.TestingBeginDate, 1*time.Millisecond)
	assert.WithinDuration(t, initialRequest.TestingEndDate, convertedRequest.TestingEndDate, 1*time.Millisecond)
	assert.Equal(t, initialRequest.Questions, convertedRequest.Questions)
}

func TestConvertToBusinessForm(t *testing.T) {
	eventForm := util.RandomEventFormCreateRequest()
	teacherUsername := gofakeit.Username()
	businessForm := form.ConvertToBusinessForm(eventForm, teacherUsername)

	assert.NotNil(t, businessForm.ID)
	assert.NotNil(t, businessForm.Name)
	assert.Equal(t, len(eventForm.Questions), len(businessForm.Questions))
	assert.Equal(t, teacherUsername, businessForm.TeacherUsername)
	for idx, question := range businessForm.Questions {
		assert.NotNil(t, question.ID)
		assertQuestionsEqual(t, eventForm.Questions[idx], businessForm.Questions[idx])
	}
}

func assertQuestionsEqual(t *testing.T, eventQuestion *event.Question, businessQuestion *business.Question) {
	assert.Equal(t, eventQuestion.Text, businessQuestion.Text)
	assert.Equal(t, len(eventQuestion.Answers), len(businessQuestion.Answers))
	for idx, answer := range businessQuestion.Answers {
		assert.NotNil(t, answer.ID)
		assertAnswersEqual(t, eventQuestion.Answers[idx], businessQuestion.Answers[idx])
	}
}

func assertAnswersEqual(t *testing.T, eventAnswer *event.Answer, businessAnswer *business.Answer) {
	assert.Equal(t, eventAnswer.Text, businessAnswer.Text)
	assert.Equal(t, eventAnswer.IsCorrect, businessAnswer.IsCorrect)
}
