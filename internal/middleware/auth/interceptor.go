package auth

import (
	"context"
	"github.com/upassed/upassed-form-service/internal/middleware"
	"google.golang.org/grpc"
	"log/slog"
	"reflect"
	"runtime"
)

func (wrapper *ClientWrapper) AuthenticationUnaryServerInterceptor() func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	op := runtime.FuncForPC(reflect.ValueOf(wrapper.AuthenticationUnaryServerInterceptor).Pointer()).Name()

	return func(ctx context.Context, request any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (response any, err error) {
		log := wrapper.log.With(
			slog.String("op", op),
			slog.String(string(middleware.RequestIDKey), middleware.GetRequestIDFromContext(ctx)),
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
