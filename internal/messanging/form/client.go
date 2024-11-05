package form

import (
	"github.com/upassed/upassed-form-service/internal/config"
	"github.com/upassed/upassed-form-service/internal/logging"
	"github.com/upassed/upassed-form-service/internal/service/form"
	"github.com/wagslane/go-rabbitmq"
	"log/slog"
)

type rabbitClient struct {
	service          form.Service
	rabbitConnection *rabbitmq.Conn
	cfg              *config.Config
	log              *slog.Logger
}

func Initialize(service form.Service, rabbitConnection *rabbitmq.Conn, cfg *config.Config, log *slog.Logger) {
	log = logging.Wrap(log,
		logging.WithOp(Initialize),
	)

	client := &rabbitClient{
		service:          service,
		rabbitConnection: rabbitConnection,
		cfg:              cfg,
		log:              log,
	}

	go func() {
		if err := InitializeCreateQueueConsumer(client); err != nil {
			log.Error("error while initializing form queue consumer", logging.Error(err))
			return
		}
	}()
}
