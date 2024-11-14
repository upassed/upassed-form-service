package form

import (
	"context"
	"errors"
	"github.com/upassed/upassed-form-service/internal/handling"
	"github.com/upassed/upassed-form-service/internal/logging"
	domain "github.com/upassed/upassed-form-service/internal/repository/model"
	"github.com/upassed/upassed-form-service/internal/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc/codes"
	"log/slog"
)

var (
	errCountingDuplicateForms = errors.New("error while counting form duplicates")
)

func (repository *formRepositoryImpl) ExistsByNameAndTeacherUsername(ctx context.Context, formName, teacherUsername string) (bool, error) {
	_, span := otel.Tracer(repository.cfg.Tracing.FormTracerName).Start(ctx, "formRepository#ExistsByNameAndTeacherUsername")
	span.SetAttributes(
		attribute.String("formName", formName),
		attribute.String("teacherUsername", teacherUsername),
	)
	defer span.End()

	log := logging.Wrap(repository.log,
		logging.WithOp(repository.ExistsByNameAndTeacherUsername),
		logging.WithCtx(ctx),
		logging.WithAny("formName", formName),
		logging.WithAny("teacherUsername", teacherUsername),
	)

	log.Info("started checking form duplicates")
	var formCount int64
	countResult := repository.db.WithContext(ctx).Model(&domain.Form{}).Where("name = ?", formName).Where("teacher_username = ?", teacherUsername).Count(&formCount)
	if err := countResult.Error; err != nil {
		log.Error("error while counting forms with name and teacherUsername in database", logging.Error(err))
		tracing.SetSpanError(span, err)
		return false, handling.New(errCountingDuplicateForms.Error(), codes.Internal)
	}

	if formCount > 0 {
		log.Info("found form duplicates in database", slog.Int64("formDuplicatesCount", formCount))
		return true, nil
	}

	log.Info("form duplicates not found in database")
	return false, nil
}
