package form

import (
	"errors"
	"github.com/upassed/upassed-form-service/internal/logging"
	"github.com/wagslane/go-rabbitmq"
)

var (
	errCreatingFormCreateQueueConsumer = errors.New("unable to create form queue consumer")
	errRunningFormCreateQueueConsumer  = errors.New("unable to run form queue consumer")
)

func InitializeCreateQueueConsumer(client *rabbitClient) error {
	log := logging.Wrap(client.log,
		logging.WithOp(InitializeCreateQueueConsumer),
	)

	log.Info("started crating form create queue consumer")
	formCreateGroupConsumer, err := rabbitmq.NewConsumer(
		client.rabbitConnection,
		client.cfg.Rabbit.Queues.FormCreate.Name,
		rabbitmq.WithConsumerOptionsRoutingKey(client.cfg.Rabbit.Queues.FormCreate.RoutingKey),
		rabbitmq.WithConsumerOptionsExchangeName(client.cfg.Rabbit.Exchange.Name),
		rabbitmq.WithConsumerOptionsExchangeDeclare,
	)

	if err != nil {
		log.Error("unable to create form queue consumer", logging.Error(err))
		return errCreatingFormCreateQueueConsumer
	}

	defer formCreateGroupConsumer.Close()
	if err := formCreateGroupConsumer.Run(client.CreateQueueConsumer()); err != nil {
		log.Error("unable to run form queue consumer")
		return errRunningFormCreateQueueConsumer
	}

	log.Info("form queue consumer successfully initialized")
	return nil
}
