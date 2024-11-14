package form

import (
	"context"
	"github.com/google/uuid"
	"github.com/upassed/upassed-form-service/internal/config"
	domain "github.com/upassed/upassed-form-service/internal/repository/model"
	"gorm.io/gorm"
	"log/slog"
)

type Repository interface {
	ExistsByNameAndTeacherUsername(ctx context.Context, formName, teacherUsername string) (bool, error)
	Save(ctx context.Context, form *domain.Form) error
	FindByID(ctx context.Context, formID uuid.UUID) (*domain.Form, error)
}

type formRepositoryImpl struct {
	db  *gorm.DB
	cfg *config.Config
	log *slog.Logger
}

func New(db *gorm.DB, cfg *config.Config, log *slog.Logger) Repository {
	return &formRepositoryImpl{
		db:  db,
		cfg: cfg,
		log: log,
	}
}
