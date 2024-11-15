package form

import (
	"context"
	"github.com/google/uuid"
	"github.com/upassed/upassed-form-service/internal/config"
	business "github.com/upassed/upassed-form-service/internal/service/model"
	"github.com/upassed/upassed-form-service/pkg/client"
	"google.golang.org/grpc"
)

type formServerAPI struct {
	client.UnimplementedFormServer
	cfg     *config.Config
	service formService
}

type formService interface {
	FindByID(ctx context.Context, formID uuid.UUID) (*business.Form, error)
	FindByTeacherUsername(ctx context.Context, teacherUsername string) ([]*business.Form, error)
}

func Register(gRPC *grpc.Server, cfg *config.Config, service formService) {
	client.RegisterFormServer(gRPC, &formServerAPI{
		cfg:     cfg,
		service: service,
	})
}
