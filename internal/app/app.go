package app

import (
	"github.com/upassed/upassed-form-service/internal/config"
	"github.com/upassed/upassed-form-service/internal/logging"
	"github.com/upassed/upassed-form-service/internal/server"
	"github.com/upassed/upassed-form-service/internal/service/form"
	"log/slog"
)

type App struct {
	Server *server.AppServer
}

func New(config *config.Config, log *slog.Logger) (*App, error) {
	log = logging.Wrap(log, logging.WithOp(New))

	appServer, err := server.New(server.AppServerCreateParams{
		Config:      config,
		Log:         log,
		FormService: form.New(config, log),
	})

	if err != nil {
		log.Error("unable to create new grpc server", logging.Error(err))
		return nil, err
	}

	log.Info("app successfully created")
	return &App{
		Server: appServer,
	}, nil
}
