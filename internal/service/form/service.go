package form

import (
	"context"
	"github.com/upassed/upassed-form-service/internal/config"
	business "github.com/upassed/upassed-form-service/internal/service/model"
	"log/slog"
)

type Service interface {
	Create(ctx context.Context, form *business.Form) (*business.FormCreateResponse, error)
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
