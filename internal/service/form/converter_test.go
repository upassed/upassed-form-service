package form_test

import (
	"github.com/stretchr/testify/assert"
	domain "github.com/upassed/upassed-form-service/internal/repository/model"
	"github.com/upassed/upassed-form-service/internal/service/form"
	business "github.com/upassed/upassed-form-service/internal/service/model"
	"github.com/upassed/upassed-form-service/internal/util"
	"testing"
	"time"
)

func TestConvertToDomainForm(t *testing.T) {
	businessForm := util.RandomBusinessForm()
	domainForm := form.ConvertToDomainForm(businessForm)

	assert.Equal(t, businessForm.ID, domainForm.ID)
	assert.Equal(t, businessForm.Name, domainForm.Name)
	assert.Equal(t, businessForm.Description, domainForm.Description)
	assert.Equal(t, businessForm.TeacherUsername, domainForm.TeacherUsername)
	assert.WithinDuration(t, businessForm.TestingBeginDate, domainForm.TestingBeginDate, 1*time.Millisecond)
	assert.WithinDuration(t, businessForm.TestingEndDate, domainForm.TestingEndDate, 1*time.Millisecond)
	assert.WithinDuration(t, businessForm.CreatedAt, domainForm.CreatedAt, 1*time.Millisecond)
	assert.Equal(t, int64(businessForm.TestingDuration.Seconds()), domainForm.TestingDurationInSeconds)
	assert.Equal(t, len(businessForm.Questions), len(domainForm.Questions))

	for idx, question := range businessForm.Questions {
		assert.NotNil(t, question.ID)
		assert.Equal(t, businessForm.ID, domainForm.Questions[idx].FormID)
		assertQuestionsEqual(t, businessForm.Questions[idx], domainForm.Questions[idx])
	}
}

func TestConvertToBusinessForm(t *testing.T) {
	domainForm := util.RandomDomainForm()
	businessForm := form.ConvertToBusinessForm(domainForm)

	assert.Equal(t, businessForm.ID, domainForm.ID)
	assert.Equal(t, businessForm.Name, domainForm.Name)
	assert.Equal(t, businessForm.Description, domainForm.Description)
	assert.Equal(t, businessForm.TeacherUsername, domainForm.TeacherUsername)
	assert.WithinDuration(t, businessForm.TestingBeginDate, domainForm.TestingBeginDate, 1*time.Millisecond)
	assert.WithinDuration(t, businessForm.TestingEndDate, domainForm.TestingEndDate, 1*time.Millisecond)
	assert.WithinDuration(t, businessForm.CreatedAt, domainForm.CreatedAt, 1*time.Millisecond)
	assert.Equal(t, int64(businessForm.TestingDuration.Seconds()), domainForm.TestingDurationInSeconds)
	assert.Equal(t, len(businessForm.Questions), len(domainForm.Questions))

	for idx, question := range businessForm.Questions {
		assert.NotNil(t, question.ID)
		assert.Equal(t, businessForm.ID, domainForm.Questions[idx].FormID)
		assertQuestionsEqual(t, businessForm.Questions[idx], domainForm.Questions[idx])
	}
}

func assertQuestionsEqual(t *testing.T, businessQuestion *business.Question, domainQuestion *domain.Question) {
	assert.Equal(t, businessQuestion.ID, domainQuestion.ID)
	assert.Equal(t, businessQuestion.Text, domainQuestion.Text)
	assert.Equal(t, len(businessQuestion.Answers), len(domainQuestion.Answers))

	for idx, answer := range businessQuestion.Answers {
		assert.NotNil(t, answer.ID)
		assert.Equal(t, businessQuestion.ID, domainQuestion.Answers[idx].QuestionID)
		assertAnswersEqual(t, businessQuestion.Answers[idx], domainQuestion.Answers[idx])
	}
}

func assertAnswersEqual(t *testing.T, businessAnswer *business.Answer, domainAnswer *domain.Answer) {
	assert.Equal(t, businessAnswer.ID, domainAnswer.ID)
	assert.Equal(t, businessAnswer.Text, domainAnswer.Text)
	assert.Equal(t, businessAnswer.IsCorrect, domainAnswer.IsCorrect)
}
