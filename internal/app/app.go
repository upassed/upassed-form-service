package app

import (
	"github.com/upassed/upassed-form-service/internal/config"
	"github.com/upassed/upassed-form-service/internal/logging"
	"github.com/upassed/upassed-form-service/internal/messanging"
	formRabbit "github.com/upassed/upassed-form-service/internal/messanging/form"
	"github.com/upassed/upassed-form-service/internal/middleware/common/auth"
	"github.com/upassed/upassed-form-service/internal/repository"
	formRepo "github.com/upassed/upassed-form-service/internal/repository/form"
	"github.com/upassed/upassed-form-service/internal/server"
	"github.com/upassed/upassed-form-service/internal/service/form"
	"log/slog"
)

type App struct {
	Server *server.AppServer
}

func New(config *config.Config, log *slog.Logger) (*App, error) {
	log = logging.Wrap(log, logging.WithOp(New))

	db, err := repository.OpenGormDbConnection(config, log)
	if err != nil {
		return nil, err
	}

	rabbit, err := messanging.OpenRabbitConnection(config, log)
	if err != nil {
		return nil, err
	}

	formRepository := formRepo.New(db, config, log)

	formService := form.New(config, log, formRepository)
	authClient, err := auth.NewClient(config, log)
	if err != nil {
		return nil, err
	}

	formRabbit.Initialize(authClient, formService, rabbit, config, log)

	appServer, err := server.New(server.AppServerCreateParams{
		Config:      config,
		Log:         log,
		FormService: formService,
		AuthClient:  authClient,
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
