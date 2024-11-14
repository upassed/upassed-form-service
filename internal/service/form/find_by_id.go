package form

import (
	"context"
	"errors"
	"github.com/google/uuid"
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
	ErrFindFormByIDDeadlineExceeded = errors.New("find form by id deadline exceeded")
)

func (service *formServiceImpl) FindByID(ctx context.Context, formID uuid.UUID) (*business.Form, error) {
	spanContext, span := otel.Tracer(service.cfg.Tracing.FormTracerName).Start(ctx, "formService#FindByID")
	span.SetAttributes(attribute.String("formID", formID.String()))
	defer span.End()

	log := logging.Wrap(service.log,
		logging.WithOp(service.FindByID),
		logging.WithCtx(ctx),
		logging.WithAny("formID", formID),
	)

	log.Info("started finding form by id")
	timeout := service.cfg.GetEndpointExecutionTimeout()
	foundForm, err := async.ExecuteWithTimeout(spanContext, timeout, func(ctx context.Context) (*business.Form, error) {
		log.Info("finding form data")
		foundForm, err := service.formRepository.FindByID(ctx, formID)
		if err != nil {
			log.Error("error while finding form data by id", logging.Error(err))
			tracing.SetSpanError(span, err)
			return nil, handling.Process(err)
		}

		return ConvertToBusinessForm(foundForm), nil
	})

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Error("find form by id deadline exceeded")
			tracing.SetSpanError(span, err)
			return nil, handling.Wrap(ErrFindFormByIDDeadlineExceeded, handling.WithCode(codes.DeadlineExceeded))
		}

		log.Error("error while finding form by id", logging.Error(err))
		tracing.SetSpanError(span, err)
		return nil, handling.Wrap(err)
	}

	log.Info("form successfully found by id")
	return foundForm, nil
}
