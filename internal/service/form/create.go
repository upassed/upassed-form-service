package form

import (
	"context"
	"errors"
	"github.com/upassed/upassed-form-service/internal/async"
	"github.com/upassed/upassed-form-service/internal/handling"
	"github.com/upassed/upassed-form-service/internal/logging"
	"github.com/upassed/upassed-form-service/internal/middleware/common/auth"
	business "github.com/upassed/upassed-form-service/internal/service/model"
	"github.com/upassed/upassed-form-service/internal/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc/codes"
	"log/slog"
)

var (
	ErrFormCreateDeadlineExceeded = errors.New("form create deadline exceeded")
)

func (service *formServiceImpl) Create(ctx context.Context, form *business.Form) (*business.FormCreateResponse, error) {
	spanContext, span := otel.Tracer(service.cfg.Tracing.FormTracerName).Start(ctx, "formService#Create")
	span.SetAttributes(attribute.String("formName", form.Name))
	defer span.End()

	log := logging.Wrap(service.log,
		logging.WithOp(service.Create),
		logging.WithCtx(ctx),
		logging.WithAny("formName", form.Name),
	)

	log.Info("started creating form")
	timeout := service.cfg.GetEndpointExecutionTimeout()
	teacherUsername := ctx.Value(auth.UsernameKey).(string)

	formCreateResponse, err := async.ExecuteWithTimeout(spanContext, timeout, func(ctx context.Context) (*business.FormCreateResponse, error) {
		log.Info("checking form duplicates")
		formExists, err := service.formRepository.ExistsByNameAndTeacherUsername(ctx, form.Name, teacherUsername)
		if err != nil {
			log.Error("unable to check form duplicates", logging.Error(err))
			tracing.SetSpanError(span, err)
			return nil, err
		}

		if formExists {
			log.Error("teacher have already created a form with this name")
			tracing.SetSpanError(span, errors.New("form duplicate found"))
			return nil, handling.Wrap(errors.New("form duplicate found"), handling.WithCode(codes.AlreadyExists))
		}

		log.Info("saving form data to the database")
		domainForm := ConvertToDomainForm(form)
		if err := service.formRepository.Save(ctx, domainForm); err != nil {
			log.Error("unable to save form data to the database", logging.Error(err))
			tracing.SetSpanError(span, err)
			return nil, err
		}

		return &business.FormCreateResponse{
			CreatedFormID: domainForm.ID,
		}, nil
	})

	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Error("form creating deadline exceeded")
			tracing.SetSpanError(span, err)
			return nil, handling.Wrap(ErrFormCreateDeadlineExceeded, handling.WithCode(codes.DeadlineExceeded))
		}

		log.Error("error while creating form", logging.Error(err))
		tracing.SetSpanError(span, err)
		return nil, handling.Process(err)
	}

	log.Info("form successfully created", slog.Any("createdFormID", formCreateResponse.CreatedFormID))
	return formCreateResponse, nil
}
