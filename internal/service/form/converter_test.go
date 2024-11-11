package form_test

import (
	"github.com/stretchr/testify/assert"
	domain "github.com/upassed/upassed-form-service/internal/repository/model"
	"github.com/upassed/upassed-form-service/internal/service/form"
	business "github.com/upassed/upassed-form-service/internal/service/model"
	"github.com/upassed/upassed-form-service/internal/util"
	"testing"
)

func TestConvertToDomainForm(t *testing.T) {
	businessForm := util.RandomBusinessForm()
	domainForm := form.ConvertToDomainForm(businessForm)

	assert.NotNil(t, domainForm.ID)
	assert.NotNil(t, domainForm.Name)
	assert.Equal(t, len(businessForm.Questions), len(domainForm.Questions))
	for idx, question := range businessForm.Questions {
		assert.NotNil(t, question.ID)
		assertQuestionsEqual(t, businessForm.Questions[idx], domainForm.Questions[idx])
	}
}

func assertQuestionsEqual(t *testing.T, businessQuestion *business.Question, domainQuestion *domain.Question) {
	assert.Equal(t, businessQuestion.ID, domainQuestion.ID)
	assert.Equal(t, businessQuestion.Text, domainQuestion.Text)
	assert.Equal(t, len(businessQuestion.Answers), len(domainQuestion.Answers))
	for idx, answer := range businessQuestion.Answers {
		assert.NotNil(t, answer.ID)
		assertAnswersEqual(t, businessQuestion.Answers[idx], domainQuestion.Answers[idx])
	}
}

func assertAnswersEqual(t *testing.T, businessAnswer *business.Answer, domainAnswer *domain.Answer) {
	assert.Equal(t, businessAnswer.ID, domainAnswer.ID)
	assert.Equal(t, businessAnswer.Text, domainAnswer.Text)
	assert.Equal(t, businessAnswer.IsCorrect, domainAnswer.IsCorrect)
}
