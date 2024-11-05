package form_test

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/upassed/upassed-form-service/internal/messanging/form"
	"github.com/upassed/upassed-form-service/internal/util"
	"testing"
)

func TestConvertToFormCreateRequest_InvalidBytes(t *testing.T) {
	invalidBytes := make([]byte, 10)
	_, err := form.ConvertToFormCreateRequest(invalidBytes)
	require.NotNil(t, err)
}

func TestConvertToFormCreateRequest_ValidBytes(t *testing.T) {
	initialRequest := util.RandomEventFormCreateRequest()
	initialRequestBytes, err := json.Marshal(initialRequest)
	require.Nil(t, err)

	convertedRequest, err := form.ConvertToFormCreateRequest(initialRequestBytes)
	require.Nil(t, err)

	assert.Equal(t, *initialRequest, *convertedRequest)
}
