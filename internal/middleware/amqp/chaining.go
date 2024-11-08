package amqp

import (
	"context"
	"github.com/upassed/upassed-form-service/internal/messanging"
	"github.com/wagslane/go-rabbitmq"
)

type Middleware func(ctx context.Context, handler messanging.HandlerWithContext) messanging.HandlerWithContext

func ChainMiddleware(ctx context.Context, handler messanging.HandlerWithContext, middlewares ...Middleware) rabbitmq.Handler {
	for i := len(middlewares) - 1; i >= 0; i-- {
		handler = middlewares[i](ctx, handler)
	}

	return func(d rabbitmq.Delivery) (action rabbitmq.Action) {
		return handler(ctx, d)
	}
}
