package form

import (
	"context"
	"github.com/google/uuid"
	"github.com/upassed/upassed-form-service/internal/config"
	domain "github.com/upassed/upassed-form-service/internal/repository/model"
	business "github.com/upassed/upassed-form-service/internal/service/model"
	"log/slog"
)

type Service interface {
	Create(ctx context.Context, form *business.Form) (*business.FormCreateResponse, error)
	FindByID(ctx context.Context, formID uuid.UUID) (*business.Form, error)
}

type formServiceImpl struct {
	cfg            *config.Config
	log            *slog.Logger
	formRepository formRepository
}

type formRepository interface {
	ExistsByNameAndTeacherUsername(ctx context.Context, formName, teacherUsername string) (bool, error)
	Save(ctx context.Context, form *domain.Form) error
}

func New(cfg *config.Config, log *slog.Logger, formRepository formRepository) Service {
	return &formServiceImpl{
		cfg:            cfg,
		log:            log,
		formRepository: formRepository,
	}
}
