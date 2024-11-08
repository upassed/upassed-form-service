package requestid

import (
	"context"
	"github.com/google/uuid"
	"github.com/upassed/upassed-form-service/internal/messanging"
	"github.com/upassed/upassed-form-service/internal/middleware/amqp"
	"github.com/wagslane/go-rabbitmq"
)

type contextKey string

const ContextKey = contextKey("requestID")

func Middleware() amqp.Middleware {
	return func(ctx context.Context, next messanging.HandlerWithContext) messanging.HandlerWithContext {
		return func(ctx context.Context, d rabbitmq.Delivery) (action rabbitmq.Action) {
			requestID := uuid.New().String()
			ctxWithRequestID := context.WithValue(ctx, ContextKey, requestID)
			return next(ctxWithRequestID, d)
		}
	}
}

func GetRequestIDFromContext(ctx context.Context) string {
	requestID, ok := ctx.Value(ContextKey).(string)
	if !ok {
		return ""
	}

	return requestID
}
