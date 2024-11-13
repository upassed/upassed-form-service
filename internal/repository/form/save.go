package form

import (
	"context"
	"errors"
	"github.com/upassed/upassed-form-service/internal/handling"
	"github.com/upassed/upassed-form-service/internal/logging"
	"github.com/upassed/upassed-form-service/internal/middleware/common/auth"
	domain "github.com/upassed/upassed-form-service/internal/repository/model"
	"github.com/upassed/upassed-form-service/internal/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc/codes"
)

var (
	ErrSavingForm = errors.New("error while saving form")
)

func (repository *formRepositoryImpl) Save(ctx context.Context, form *domain.Form) error {
	teacherUsername := ctx.Value(auth.UsernameKey).(string)

	_, span := otel.Tracer(repository.cfg.Tracing.FormTracerName).Start(ctx, "formRepository#Save")
	span.SetAttributes(
		attribute.String("formName", form.Name),
		attribute.String(auth.UsernameKey, teacherUsername),
	)
	defer span.End()

	log := logging.Wrap(repository.log,
		logging.WithOp(repository.Save),
		logging.WithCtx(ctx),
		logging.WithAny("formName", form.Name),
	)

	log.Info("started saving form to a database")
	form.TeacherUsername = teacherUsername

	saveResult := repository.db.WithContext(ctx).Create(form)

	var form2 domain.Form
	repository.db.WithContext(ctx).First(&form2)

	if err := saveResult.Error; err != nil || saveResult.RowsAffected != 1 {
		log.Error("error while saving form data to a database", logging.Error(err))
		tracing.SetSpanError(span, err)
		return handling.New(ErrSavingForm.Error(), codes.Internal)
	}

	log.Info("form was successfully inserted into a database")
	return nil
}
