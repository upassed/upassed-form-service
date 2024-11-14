package form

import (
	"context"
	"github.com/google/uuid"
	domain "github.com/upassed/upassed-form-service/internal/repository/model"
)

func (repository *formRepositoryImpl) FindByID(ctx context.Context, formID uuid.UUID) (*domain.Form, error) {
	//TODO implement me
	panic("implement me")
}
