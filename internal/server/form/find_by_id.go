package form

import (
	"context"
	"github.com/google/uuid"
	"github.com/upassed/upassed-form-service/internal/handling"
	requestid "github.com/upassed/upassed-form-service/internal/middleware/common/request_id"
	"github.com/upassed/upassed-form-service/internal/tracing"
	"github.com/upassed/upassed-form-service/pkg/client"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc/codes"
)

func (server *formServerAPI) FindByID(ctx context.Context, request *client.FormFindByIDRequest) (*client.FormFindByIDResponse, error) {
	spanContext, span := otel.Tracer(server.cfg.Tracing.FormTracerName).Start(ctx, "form#FindByID")
	span.SetAttributes(
		attribute.String(string(requestid.ContextKey), requestid.GetRequestIDFromContext(ctx)),
		attribute.String("formID", request.GetFormId()),
	)
	defer span.End()

	if err := request.Validate(); err != nil {
		tracing.SetSpanError(span, err)
		return nil, handling.Wrap(err, handling.WithCode(codes.InvalidArgument))
	}

	form, err := server.service.FindByID(spanContext, uuid.MustParse(request.GetFormId()))
	if err != nil {
		tracing.SetSpanError(span, err)
		return nil, err
	}

	return ConvertToFindByIdResponse(form), nil
}
