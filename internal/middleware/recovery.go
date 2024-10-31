package middleware

import (
	"context"
	"log/slog"
	"reflect"
	"runtime"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func PanicRecoveryMiddlewareInterceptor(log *slog.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		defer func() {
			if r := recover(); r != nil {
				op := runtime.FuncForPC(reflect.ValueOf(PanicRecoveryMiddlewareInterceptor).Pointer()).Name()

				log := log.With(
					slog.String("op", op),
				)

				log.Error("panic recovered",
					slog.String("requestID", GetRequestIDFromContext(ctx)),
					slog.Any("message", r),
					slog.String("method", info.FullMethod),
				)

				err = status.Errorf(codes.Internal, "internal server error")
			}
		}()

		return handler(ctx, req)
	}
}
