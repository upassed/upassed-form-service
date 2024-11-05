package auth

import (
	"context"
	"errors"
	"github.com/upassed/upassed-authentication-service/pkg/client"
	"github.com/upassed/upassed-form-service/internal/handling"
	"github.com/upassed/upassed-form-service/internal/logging"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"log/slog"
)

func (wrapper *ClientWrapper) teacherAccountTypeAuthenticationFunc(ctx context.Context) (context.Context, error) {
	log := logging.Wrap(
		wrapper.log,
		logging.WithOp(wrapper.teacherAccountTypeAuthenticationFunc),
		logging.WithCtx(ctx),
	)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		log.Error("unable to extract metadata from incoming context")
		return nil, handling.Wrap(errors.New("unable to extract metadata"), handling.WithCode(codes.Internal))
	}

	token, ok := md[authenticationHeaderKey]
	if !ok || len(token) != 1 {
		log.Error("missing authentication header in request metadata")
		return nil, handling.Wrap(errors.New("unable to extract authentication header with jwt token"), handling.WithCode(codes.Unauthenticated))
	}

	timeout := wrapper.cfg.GetEndpointExecutionTimeout()
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	response, err := wrapper.authenticationServiceClient.Validate(ctx, &client.TokenValidateRequest{
		AccessToken: token[0],
	})

	if err != nil {
		log.Error("error while validating token on an authentication service", slog.String("err", err.Error()))
		return nil, handling.Wrap(errors.New("validate token error"), handling.WithCode(codes.Unauthenticated))
	}

	enrichedContext := context.WithValue(ctx, usernameKey, response.GetUsername())
	if !(response.GetAccountType() == "TEACHER") {
		log.Error("account type is not equal to teacher", slog.String("accountType", response.GetAccountType()))
		return nil, handling.Wrap(errors.New("required teacher account type"), handling.WithCode(codes.PermissionDenied))
	}

	return enrichedContext, nil
}
