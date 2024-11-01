package form

import (
	"context"
	"github.com/google/uuid"
	business "github.com/upassed/upassed-form-service/internal/service/model"
)

func (service *formServiceImpl) FindByID(ctx context.Context, formID uuid.UUID) (*business.Form, error) {
	return &business.Form{
		ID: formID,
	}, nil
}
