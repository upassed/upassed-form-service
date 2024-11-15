package form

import (
	"context"
	"github.com/upassed/upassed-form-service/internal/handling"
	"github.com/upassed/upassed-form-service/internal/middleware/common/auth"
	requestid "github.com/upassed/upassed-form-service/internal/middleware/common/request_id"
	"github.com/upassed/upassed-form-service/internal/tracing"
	"github.com/upassed/upassed-form-service/pkg/client"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc/codes"
)

func (server *formServerAPI) FindByTeacherUsername(ctx context.Context, request *client.FormFindByTeacherUsernameRequest) (*client.FormFindByTeacherUsernameResponse, error) {
	teacherUsername := ctx.Value(auth.UsernameKey).(string)

	spanContext, span := otel.Tracer(server.cfg.Tracing.FormTracerName).Start(ctx, "form#FindByTeacherUsername")
	span.SetAttributes(
		attribute.String(string(requestid.ContextKey), requestid.GetRequestIDFromContext(ctx)),
		attribute.String("teacherUsername", teacherUsername),
	)
	defer span.End()

	if err := request.Validate(); err != nil {
		tracing.SetSpanError(span, err)
		return nil, handling.Wrap(err, handling.WithCode(codes.InvalidArgument))
	}

	form, err := server.service.FindByTeacherUsername(spanContext, teacherUsername)
	if err != nil {
		tracing.SetSpanError(span, err)
		return nil, err
	}

	return ConvertToFindByTeacherUsernameResponse(form), nil
}
