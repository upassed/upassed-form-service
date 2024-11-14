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

func (wrapper *clientImpl) TeacherAccountTypeAuthenticationFunc(ctx context.Context, token string) (context.Context, error) {
	log := logging.Wrap(
		wrapper.log,
		logging.WithOp(wrapper.TeacherAccountTypeAuthenticationFunc),
		logging.WithCtx(ctx),
	)

	timeout := wrapper.cfg.GetEndpointExecutionTimeout()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	response, err := wrapper.authenticationServiceClient.Validate(ctx, &client.TokenValidateRequest{
		AccessToken: token,
	})

	if err != nil {
		log.Error("error while validating token on an authentication service", slog.String("err", err.Error()))
		return nil, handling.Wrap(errors.New("validate token error"), handling.WithCode(codes.Unauthenticated))
	}

	enrichedContext := context.WithValue(ctx, UsernameKey, response.GetUsername())
	if !(response.GetAccountType() == "TEACHER") {
		log.Error("account type is not equal to teacher", slog.String("accountType", response.GetAccountType()))
		return nil, handling.Wrap(errors.New("required teacher account type"), handling.WithCode(codes.PermissionDenied))
	}

	return enrichedContext, nil
}
