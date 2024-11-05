package auth

import (
	"context"
	"github.com/upassed/upassed-form-service/internal/logging"
	"google.golang.org/grpc"
	"log/slog"
)

func (wrapper *ClientWrapper) AuthenticationUnaryServerInterceptor() func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	return func(ctx context.Context, request any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (response any, err error) {
		log := logging.Wrap(
			wrapper.log,
			logging.WithOp(wrapper.AuthenticationUnaryServerInterceptor),
			logging.WithCtx(ctx),
		)

		authenticationFunc, ok := authenticationRules[info.FullMethod]
		if !ok {
			authenticationFunc = wrapper.anyAccountTypeAuthenticationFunc
		}

		enrichedCtx, err := authenticationFunc(ctx)
		if err != nil {
			log.Error("authentication failed", slog.String("err", err.Error()))
			return nil, err
		}

		return handler(enrichedCtx, request)
	}
}
