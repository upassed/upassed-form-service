package auth

import (
	"context"
	"errors"
	"github.com/upassed/upassed-authentication-service/pkg/client"
	"github.com/upassed/upassed-form-service/internal/handling"
	"github.com/upassed/upassed-form-service/internal/logging"
	"google.golang.org/grpc/codes"
	"log/slog"
)

func (wrapper *ClientWrapper) AnyAccountTypeAuthenticationFunc(ctx context.Context, token string) (context.Context, error) {
	log := logging.Wrap(
		wrapper.log,
		logging.WithOp(wrapper.AnyAccountTypeAuthenticationFunc),
		logging.WithCtx(ctx),
	)

	timeout := wrapper.cfg.GetEndpointExecutionTimeout()
	ctxWithTimeout, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	response, err := wrapper.authenticationServiceClient.Validate(ctxWithTimeout, &client.TokenValidateRequest{
		AccessToken: token,
	})

	if err != nil {
		log.Error("error while validating token on an authentication service", slog.String("err", err.Error()))
		return nil, handling.Wrap(errors.New("validate token error"), handling.WithCode(codes.Unauthenticated))
	}

	enrichedContext := context.WithValue(ctx, usernameKey, response.GetUsername())
	return enrichedContext, nil
}