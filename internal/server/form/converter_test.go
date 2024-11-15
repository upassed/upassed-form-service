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
	assertFormEqual(t, businessForm, convertedForm)
}

func TestConvertToFindByTeacherUsernameResponse(t *testing.T) {
	businessForms := []*business.Form{util.RandomBusinessForm(), util.RandomBusinessForm(), util.RandomBusinessForm()}
	response := form.ConvertToFindByTeacherUsernameResponse(businessForms)

	require.NotNil(t, response.GetFoundForms())

	for idx, convertedForm := range response.GetFoundForms() {
		assertFormEqual(t, businessForms[idx], convertedForm)
	}
}

func assertFormEqual(t *testing.T, businessForm *business.Form, clientForm *client.FormDTO) {
	assert.Equal(t, businessForm.ID.String(), clientForm.GetId())
	assert.Equal(t, businessForm.Name, clientForm.GetName())
	assert.Equal(t, businessForm.TeacherUsername, clientForm.GetTeacherUsername())
	assert.Equal(t, businessForm.Description, clientForm.GetDescription())
	assert.WithinDuration(t, businessForm.TestingBeginDate, clientForm.GetTestingBeginDate().AsTime(), 1*time.Millisecond)
	assert.WithinDuration(t, businessForm.TestingEndDate, clientForm.GetTestingEndDate().AsTime(), 1*time.Millisecond)
	assert.WithinDuration(t, businessForm.CreatedAt, clientForm.GetCreatedAt().AsTime(), 1*time.Millisecond)
	assert.Equal(t, len(businessForm.Questions), len(clientForm.GetQuestions()))

	for idx, question := range businessForm.Questions {
		assertQuestionsEqual(t, question, clientForm.GetQuestions()[idx])
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
