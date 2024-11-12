package form

import (
	"context"
	"github.com/upassed/upassed-form-service/internal/logging"
	"github.com/upassed/upassed-form-service/internal/middleware/amqp"
	loggingMiddleware "github.com/upassed/upassed-form-service/internal/middleware/amqp/logging"
	"github.com/upassed/upassed-form-service/internal/middleware/amqp/recovery"
	requestidMiddleware "github.com/upassed/upassed-form-service/internal/middleware/amqp/request_id"
	"github.com/upassed/upassed-form-service/internal/middleware/common/auth"
	"github.com/upassed/upassed-form-service/internal/middleware/common/request_id"
	"github.com/upassed/upassed-form-service/internal/tracing"
	"github.com/wagslane/go-rabbitmq"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"log/slog"
)

func (client *rabbitClient) CreateQueueConsumer() rabbitmq.Handler {
	baseHandler := func(ctx context.Context, delivery rabbitmq.Delivery) rabbitmq.Action {
		log := logging.Wrap(client.log,
			logging.WithOp(client.CreateQueueConsumer),
			logging.WithCtx(ctx),
		)

		log.Info("consumed form create message", slog.String("messageBody", string(delivery.Body)))
		spanContext, span := otel.Tracer(client.cfg.Tracing.FormTracerName).Start(ctx, "form#Create")
		span.SetAttributes(attribute.String(string(requestid.ContextKey), requestid.GetRequestIDFromContext(ctx)))
		defer span.End()

		log.Info("converting message body to form create request struct")
		request, err := ConvertToFormCreateRequest(delivery.Body)
		if err != nil {
			log.Error("unable to convert message body to create request struct", logging.Error(err))
			tracing.SetSpanError(span, err)
			return rabbitmq.NackDiscard
		}

		teacherUsername := ctx.Value(auth.UsernameKey).(string)
		span.SetAttributes(
			attribute.String("formName", request.Name),
			attribute.String(auth.UsernameKey, teacherUsername),
		)

		log.Info("validating form create request")
		if err := request.Validate(); err != nil {
			log.Error("form create request is invalid", logging.Error(err))
			tracing.SetSpanError(span, err)
			return rabbitmq.NackDiscard
		}

		log.Info("creating form")
		response, err := client.service.Create(spanContext, ConvertToBusinessForm(request))
		if err != nil {
			log.Error("unable to create form", logging.Error(err))
			tracing.SetSpanError(span, err)
			return rabbitmq.NackDiscard
		}

		log.Info("successfully created form", slog.Any("createdFormID", response.CreatedFormID))
		return rabbitmq.Ack
	}

	handlerWithMiddleware := amqp.ChainMiddleware(
		baseHandler,
		requestidMiddleware.Middleware(),
		loggingMiddleware.Middleware(client.log),
		client.authClient.AmqpMiddleware(client.log),
		recovery.Middleware(client.log),
	)

	return func(d rabbitmq.Delivery) (action rabbitmq.Action) {
		ctx := context.Background()
		return handlerWithMiddleware(ctx, d)
	}
}
