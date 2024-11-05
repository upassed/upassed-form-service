package form

import (
	"encoding/json"
	event "github.com/upassed/upassed-form-service/internal/messanging/model"
)

func ConvertToFormCreateRequest(messageBody []byte) (*event.FormCreateRequest, error) {
	var request event.FormCreateRequest
	if err := json.Unmarshal(messageBody, &request); err != nil {
		return nil, err
	}

	return &request, nil
}
