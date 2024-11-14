package form_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/upassed/upassed-form-service/internal/server/form"
	business "github.com/upassed/upassed-form-service/internal/service/model"
	"github.com/upassed/upassed-form-service/internal/util"
	"github.com/upassed/upassed-form-service/pkg/client"
	"testing"
	"time"
)

func TestConvertToFindByIdResponse(t *testing.T) {
	businessForm := util.RandomBusinessForm()
	response := form.ConvertToFindByIdResponse(businessForm)

	require.NotNil(t, response.GetForm())

	convertedForm := response.GetForm()
	assert.Equal(t, businessForm.ID.String(), convertedForm.GetId())
	assert.Equal(t, businessForm.Name, convertedForm.GetName())
	assert.Equal(t, businessForm.TeacherUsername, convertedForm.GetTeacherUsername())
	assert.Equal(t, businessForm.Description, convertedForm.GetDescription())
	assert.WithinDuration(t, businessForm.TestingBeginDate, convertedForm.GetTestingBeginDate().AsTime(), 1*time.Millisecond)
	assert.WithinDuration(t, businessForm.TestingEndDate, convertedForm.GetTestingEndDate().AsTime(), 1*time.Millisecond)
	assert.WithinDuration(t, businessForm.CreatedAt, convertedForm.GetCreatedAt().AsTime(), 1*time.Millisecond)
	assert.Equal(t, len(businessForm.Questions), len(convertedForm.GetQuestions()))

	for idx, question := range businessForm.Questions {
		assertQuestionsEqual(t, question, convertedForm.GetQuestions()[idx])
	}
}

func assertQuestionsEqual(t *testing.T, businessQuestion *business.Question, clientQuestion *client.QuestionDTO) {
	assert.Equal(t, businessQuestion.ID.String(), clientQuestion.GetId())
	assert.Equal(t, businessQuestion.Text, clientQuestion.GetText())

	for idx, answer := range businessQuestion.Answers {
		assertAnswersEqual(t, answer, clientQuestion.GetAnswers()[idx])
	}
}

func assertAnswersEqual(t *testing.T, businessAnswer *business.Answer, clientAnswer *client.AnswerDTO) {
	assert.Equal(t, businessAnswer.ID.String(), clientAnswer.GetId())
	assert.Equal(t, businessAnswer.Text, clientAnswer.GetText())
	assert.Equal(t, businessAnswer.IsCorrect, clientAnswer.GetIsCorrect())
}
