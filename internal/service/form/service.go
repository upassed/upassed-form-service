package form

import (
	"context"
	"github.com/google/uuid"
	"github.com/upassed/upassed-form-service/internal/config"
	business "github.com/upassed/upassed-form-service/internal/service/model"
	"log/slog"
)

type Service interface {
	FindByID(context.Context, uuid.UUID) (*business.Form, error)
}

type formServiceImpl struct {
	cfg *config.Config
	log *slog.Logger
}

func New(cfg *config.Config, log *slog.Logger) Service {
	return &formServiceImpl{
		cfg: cfg,
		log: log,
	}
}
