package form

import (
	"context"
	"github.com/google/uuid"
	"github.com/upassed/upassed-form-service/pkg/client"
)

func (server *formServerAPI) FindByID(ctx context.Context, request *client.FormFindByIDRequest) (*client.FormFindByIDResponse, error) {
	foundForm, _ := server.service.FindByID(ctx, uuid.MustParse(request.FormId))

	return &client.FormFindByIDResponse{
		Form: &client.FormDTO{
			Id: foundForm.ID.String(),
		},
	}, nil
}
