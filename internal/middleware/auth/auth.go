package auth

import (
	"errors"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/upassed/upassed-authentication-service/pkg/client"
	"github.com/upassed/upassed-form-service/internal/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log/slog"
	"net"
	"reflect"
	"runtime"
)

var (
	errCreatingAuthServiceConn = errors.New("unable to create authentication service connection")
)

const (
	authenticationHeaderKey = "authentication"
	usernameKey             = "username"
)

var authenticationRules = map[string]auth.AuthFunc{}

type ClientWrapper struct {
	cfg                         *config.Config
	log                         *slog.Logger
	authenticationServiceClient client.TokenClient
}

func NewClient(cfg *config.Config, log *slog.Logger) (*ClientWrapper, error) {
	op := runtime.FuncForPC(reflect.ValueOf(NewClient).Pointer()).Name()

	authenticationServiceUrl := net.JoinHostPort(
		cfg.Services.Authentication.Host,
		cfg.Services.Authentication.Port,
	)

	log = log.With(
		slog.String("op", op),
		slog.String("authentication-service-url", authenticationServiceUrl),
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
