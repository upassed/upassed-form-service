package form

import (
	"context"
	"errors"
	"github.com/upassed/upassed-form-service/internal/async"
	"github.com/upassed/upassed-form-service/internal/handling"
	"github.com/upassed/upassed-form-service/internal/logging"
	business "github.com/upassed/upassed-form-service/internal/service/model"
	"github.com/upassed/upassed-form-service/internal/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc/codes"
)

var (
	ErrFindFormsByTeacherUsernameDeadlineExceeded = errors.New("forms find by teacher username deadline exceeded")
)

func (service *formServiceImpl) FindByTeacherUsername(ctx context.Context, teacherUsername string) ([]*business.Form, error) {
	spanContext, span := otel.Tracer(service.cfg.Tracing.FormTracerName).Start(ctx, "formService#FindByTeacherUsername")
	span.SetAttributes(attribute.String("teacherUsername", teacherUsername))
	defer span.End()

	log := logging.Wrap(service.log,
		logging.WithOp(service.FindByTeacherUsername),
		logging.WithCtx(ctx),
		logging.WithAny("teacherUsername", teacherUsername),
	)

	log.Info("started finding forms by teacher username")
	timeout := service.cfg.GetEndpointExecutionTimeout()
	foundForms, err := async.ExecuteWithTimeout(spanContext, timeout, func(ctx context.Context) ([]*business.Form, error) {
		log.Info("finding forms data")
		foundForms, err := service.formRepository.FindByTeacherUsername(ctx, teacherUsername)
		if err != nil {
			log.Error("error while finding forms data by teacherUsername", logging.Error(err))
			tracing.SetSpanError(span, err)
			return nil, handling.Process(err)
		}

		return ConvertToBusinessForms(foundForms), nil
	})

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Error("find forms by teacher username deadline exceeded")
			tracing.SetSpanError(span, err)
			return nil, handling.Wrap(ErrFindFormsByTeacherUsernameDeadlineExceeded, handling.WithCode(codes.DeadlineExceeded))
		}

		log.Error("error while finding forms by teacher username", logging.Error(err))
		tracing.SetSpanError(span, err)
		return nil, handling.Wrap(err)
	}

	log.Info("forms successfully found by teacher username")
	return foundForms, nil
}
