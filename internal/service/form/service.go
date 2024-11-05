package form

import (
	"github.com/upassed/upassed-form-service/internal/config"
	"log/slog"
)

type Service interface {
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
