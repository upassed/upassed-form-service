package form

import (
	"context"
	"github.com/google/uuid"
	"github.com/upassed/upassed-form-service/internal/logging"
	"github.com/upassed/upassed-form-service/internal/middleware/requestid"
	"github.com/wagslane/go-rabbitmq"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"log/slog"
)

func (client *rabbitClient) CreateQueueConsumer() func(d rabbitmq.Delivery) rabbitmq.Action {
	return func(delivery rabbitmq.Delivery) rabbitmq.Action {
		requestID := uuid.New().String()
		ctx := context.WithValue(context.Background(), requestid.ContextKey, requestID)

		log := logging.Wrap(client.log,
			logging.WithOp(client.CreateQueueConsumer),
			logging.WithCtx(ctx),
		)

		log.Info("consumed form create message", slog.String("messageBody", string(delivery.Body)))
		_, span := otel.Tracer(client.cfg.Tracing.FormTracerName).Start(ctx, "form#Create")
		span.SetAttributes(attribute.String(string(requestid.ContextKey), requestid.GetRequestIDFromContext(ctx)))
		defer span.End()

		log.Info("converting message body to form create request struct")

		log.Info("successfully created form", slog.Any("createdFormID", uuid.New()))
		return rabbitmq.Ack
	}
}
