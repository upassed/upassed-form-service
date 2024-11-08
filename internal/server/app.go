package server

import (
	"errors"
	"fmt"
	"github.com/upassed/upassed-form-service/internal/config"
	"github.com/upassed/upassed-form-service/internal/logging"
	"github.com/upassed/upassed-form-service/internal/middleware/auth"
	loggingMiddleware "github.com/upassed/upassed-form-service/internal/middleware/grpc/logging"
	"github.com/upassed/upassed-form-service/internal/middleware/grpc/recovery"
	"github.com/upassed/upassed-form-service/internal/middleware/grpc/request_id"
	formSvc "github.com/upassed/upassed-form-service/internal/service/form"
	"google.golang.org/grpc"
	"log/slog"
	"net"
)

var (
	errStartingTcpConnection = errors.New("unable to start tcp connection")
	errStartingServer        = errors.New("unable to start gRPC server")
)

type AppServer struct {
	config *config.Config
	log    *slog.Logger
	server *grpc.Server
}

type AppServerCreateParams struct {
	Config      *config.Config
	Log         *slog.Logger
	FormService formSvc.Service
}

func New(params AppServerCreateParams) (*AppServer, error) {
	authenticationClient, err := auth.NewClient(params.Config, params.Log)
	if err != nil {
		return nil, err
	}

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			requestid.MiddlewareInterceptor(),
			recovery.MiddlewareInterceptor(params.Log),
			loggingMiddleware.MiddlewareInterceptor(params.Log),
			authenticationClient.AuthenticationUnaryServerInterceptor(),
		),
	)

	return &AppServer{
		config: params.Config,
		log:    params.Log,
		server: server,
	}, nil
}

func (server *AppServer) Run() error {
	log := logging.Wrap(server.log,
		logging.WithOp(server.Run),
	)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%s", server.config.GrpcServer.Port))
	if err != nil {
		return errStartingTcpConnection
	}

	log.Info("gRPC server is running", slog.String("address", listener.Addr().String()))
	if err := server.server.Serve(listener); err != nil {
		return errStartingServer
	}

	return nil
}

func (server *AppServer) GracefulStop() {
	log := logging.Wrap(server.log,
		logging.WithOp(server.GracefulStop),
	)

	log.Info("gracefully stopping gRPC server...")
	server.server.GracefulStop()
}
