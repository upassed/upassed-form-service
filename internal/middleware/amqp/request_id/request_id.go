package requestid

import (
	"context"
	"github.com/google/uuid"
	"github.com/upassed/upassed-form-service/internal/messanging"
	"github.com/upassed/upassed-form-service/internal/middleware/amqp"
	"github.com/upassed/upassed-form-service/internal/middleware/common/request_id"
	"github.com/wagslane/go-rabbitmq"
)

func Middleware() amqp.Middleware {
	return func(ctx context.Context, next messanging.HandlerWithContext) messanging.HandlerWithContext {
		return func(ctx context.Context, d rabbitmq.Delivery) (action rabbitmq.Action) {
			requestID := uuid.New().String()
			ctxWithRequestID := context.WithValue(ctx, requestid.ContextKey, requestID)
			return next(ctxWithRequestID, d)
		}
	}
}
