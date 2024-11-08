package auth

import (
	"context"
	"errors"
	"github.com/upassed/upassed-authentication-service/pkg/client"
	"github.com/upassed/upassed-form-service/internal/config"
	"github.com/upassed/upassed-form-service/internal/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"net"
)

var (
	errCreatingAuthServiceConn = errors.New("unable to create authentication service connection")
)

const (
	authenticationHeaderKey = "authentication"
	usernameKey             = "username"
)

type tokenAuthFunc func(ctx context.Context, token string) (context.Context, error)

type ClientWrapper struct {
	cfg                         *config.Config
	log                         *slog.Logger
	authenticationServiceClient client.TokenClient
}

func NewClient(cfg *config.Config, log *slog.Logger) (*ClientWrapper, error) {
	authenticationServiceUrl := net.JoinHostPort(
		cfg.Services.Authentication.Host,
		cfg.Services.Authentication.Port,
	)

	log = logging.Wrap(
		log,
		logging.WithOp(NewClient),
		logging.WithAny("authentication-service-url", authenticationServiceUrl),
	)

	authenticationServiceConnection, err := grpc.NewClient(authenticationServiceUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Error("unable to create authentication client connection", slog.String("err", err.Error()))
		return nil, errCreatingAuthServiceConn
	}

	return &ClientWrapper{
		cfg:                         cfg,
		log:                         log,
		authenticationServiceClient: client.NewTokenClient(authenticationServiceConnection),
	}, nil
}
