package form

import (
	"context"
	errors "errors"
	"github.com/google/uuid"
	"github.com/upassed/upassed-form-service/internal/handling"
	"github.com/upassed/upassed-form-service/internal/logging"
	domain "github.com/upassed/upassed-form-service/internal/repository/model"
	"github.com/upassed/upassed-form-service/internal/tracing"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"google.golang.org/grpc/codes"
	"gorm.io/gorm"
)

var (
	ErrFormNotFoundByID  = errors.New("form was not found by id in the database")
	errSearchingFormByID = errors.New("error while searching form by id")
)

func (repository *formRepositoryImpl) FindByID(ctx context.Context, formID uuid.UUID) (*domain.Form, error) {
	_, span := otel.Tracer(repository.cfg.Tracing.FormTracerName).Start(ctx, "formRepository#FindByID")
	span.SetAttributes(attribute.String("formID", formID.String()))
	defer span.End()

	log := logging.Wrap(repository.log,
		logging.WithOp(repository.FindByID),
		logging.WithCtx(ctx),
		logging.WithAny("formID", formID),
	)

	log.Info("started searching form by id in a database")
	foundForm := domain.Form{}
	searchResult := repository.db.WithContext(ctx).Preload("Questions.Answers").First(&foundForm, formID)
	if err := searchResult.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Error("form by id was not found in the database", logging.Error(err))
			tracing.SetSpanError(span, err)
			return nil, handling.New(ErrFormNotFoundByID.Error(), codes.NotFound)
		}

		log.Error("error while searching form by id in the database", logging.Error(err))
		tracing.SetSpanError(span, err)
		return nil, handling.New(errSearchingFormByID.Error(), codes.Internal)
	}

	log.Info("form by id was successfully found in a database")
	return &foundForm, nil
}
