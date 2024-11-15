package form

import (
	"context"
	"github.com/upassed/upassed-form-service/internal/handling"
	"github.com/upassed/upassed-form-service/internal/logging"
	domain "github.com/upassed/upassed-form-service/internal/repository/model"
	"github.com/upassed/upassed-form-service/internal/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc/codes"
)

func (repository *formRepositoryImpl) FindByTeacherUsername(ctx context.Context, teacherUsername string) ([]*domain.Form, error) {
	_, span := otel.Tracer(repository.cfg.Tracing.FormTracerName).Start(ctx, "formRepository#FindByTeacherUsername")
	span.SetAttributes(attribute.String("teacherUsername", teacherUsername))
	defer span.End()

	log := logging.Wrap(repository.log,
		logging.WithOp(repository.FindByTeacherUsername),
		logging.WithCtx(ctx),
		logging.WithAny("teacherUsername", teacherUsername),
	)

	log.Info("started searching forms by teacher username in a database")
	foundForms := make([]*domain.Form, 0)
	searchResult := repository.db.WithContext(ctx).Preload("Questions.Answers").Where("teacher_username = ?", teacherUsername).Find(&foundForms)
	if err := searchResult.Error; err != nil {
		log.Error("error while searching forms by teacher username in the database", logging.Error(err))
		tracing.SetSpanError(span, err)
		return nil, handling.New(errSearchingFormByID.Error(), codes.Internal)
	}

	log.Info("forms by teacher username were successfully found in a database")
	return foundForms, nil
}
