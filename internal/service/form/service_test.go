package form_test

import (
	"context"
	"errors"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/upassed/upassed-form-service/internal/config"
	"github.com/upassed/upassed-form-service/internal/logging"
	"github.com/upassed/upassed-form-service/internal/middleware/common/auth"
	formSvc "github.com/upassed/upassed-form-service/internal/service/form"
	"github.com/upassed/upassed-form-service/internal/util"
	"github.com/upassed/upassed-form-service/internal/util/mocks"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"path/filepath"
	"testing"
)

var (
	cfg        *config.Config
	repository *mocks.FormRepository
	service    formSvc.Service
)

func TestMain(m *testing.M) {
	currentDir, _ := os.Getwd()
	projectRoot, err := util.GetProjectRoot(currentDir)
	if err != nil {
		log.Fatal("error to get project root folder: ", err)
	}

	if err := os.Setenv(config.EnvConfigPath, filepath.Join(projectRoot, "config", "test.yml")); err != nil {
		log.Fatal(err)
	}

	cfg, err = config.Load()
	if err != nil {
		log.Fatal("unable to parse config: ", err)
	}

	ctrl := gomock.NewController(nil)
	defer ctrl.Finish()

	repository = mocks.NewFormRepository(ctrl)
	service = formSvc.New(cfg, logging.New(config.EnvTesting), repository)

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCreate_ErrorCheckingFormExists(t *testing.T) {
	teacherUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, teacherUsername)

	formToCreate := util.RandomBusinessForm()
	expectedRepositoryError := errors.New("some repo error")

	repository.EXPECT().
		ExistsByNameAndTeacherUsername(gomock.Any(), formToCreate.Name, teacherUsername).
		Return(false, expectedRepositoryError)

	_, err := service.Create(ctx, formToCreate)
	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, expectedRepositoryError.Error(), convertedError.Message())
	assert.Equal(t, codes.Internal, convertedError.Code())
}

func TestCreate_DuplicateExists(t *testing.T) {
	teacherUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, teacherUsername)

	formToCreate := util.RandomBusinessForm()

	repository.EXPECT().
		ExistsByNameAndTeacherUsername(gomock.Any(), formToCreate.Name, teacherUsername).
		Return(true, nil)

	_, err := service.Create(ctx, formToCreate)
	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, "form duplicate found", convertedError.Message())
	assert.Equal(t, codes.AlreadyExists, convertedError.Code())
}

func TestCreate_ErrorSavingForm(t *testing.T) {
	teacherUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, teacherUsername)

	formToCreate := util.RandomBusinessForm()
	repository.EXPECT().
		ExistsByNameAndTeacherUsername(gomock.Any(), formToCreate.Name, teacherUsername).
		Return(false, nil)

	expectedRepoError := errors.New("some repo error")
	repository.EXPECT().
		Save(gomock.Any(), gomock.Any()).
		Return(expectedRepoError)

	_, err := service.Create(ctx, formToCreate)
	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, expectedRepoError.Error(), convertedError.Message())
	assert.Equal(t, codes.Internal, convertedError.Code())
}

func TestCreate_ErrorDeadlineExceeded(t *testing.T) {
	oldTimeout := cfg.Timeouts.EndpointExecutionTimeoutMS
	cfg.Timeouts.EndpointExecutionTimeoutMS = "0"

	teacherUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, teacherUsername)

	formToCreate := util.RandomBusinessForm()

	repository.EXPECT().
		ExistsByNameAndTeacherUsername(gomock.Any(), formToCreate.Name, teacherUsername).
		Return(false, nil)

	repository.EXPECT().
		Save(gomock.Any(), gomock.Any()).
		Return(nil)

	_, err := service.Create(ctx, formToCreate)
	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, formSvc.ErrFormCreateDeadlineExceeded.Error(), convertedError.Message())
	assert.Equal(t, codes.DeadlineExceeded, convertedError.Code())

	cfg.Timeouts.EndpointExecutionTimeoutMS = oldTimeout
}

func TestCreate_HappyPath(t *testing.T) {
	teacherUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, teacherUsername)

	formToCreate := util.RandomBusinessForm()

	repository.EXPECT().
		ExistsByNameAndTeacherUsername(gomock.Any(), formToCreate.Name, teacherUsername).
		Return(false, nil)

	repository.EXPECT().
		Save(gomock.Any(), gomock.Any()).
		Return(nil)

	response, err := service.Create(ctx, formToCreate)
	require.NoError(t, err)

	require.NotNil(t, response.CreatedFormID)
}

func TestFindByID_RepositoryError(t *testing.T) {
	teacherUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, teacherUsername)

	formID := uuid.New()
	expectedRepositoryError := errors.New("some repo error")

	repository.EXPECT().
		FindByID(gomock.Any(), formID).
		Return(nil, expectedRepositoryError)

	_, err := service.FindByID(ctx, formID)
	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, expectedRepositoryError.Error(), convertedError.Message())
	assert.Equal(t, codes.Internal, convertedError.Code())
}

func TestFindByID_ErrorDeadlineExceeded(t *testing.T) {
	oldTimeout := cfg.Timeouts.EndpointExecutionTimeoutMS
	cfg.Timeouts.EndpointExecutionTimeoutMS = "0"

	teacherUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, teacherUsername)

	formID := uuid.New()

	repository.EXPECT().
		FindByID(gomock.Any(), formID).
		Return(util.RandomDomainForm(), nil)

	_, err := service.FindByID(ctx, formID)
	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, formSvc.ErrFindFormByIDDeadlineExceeded.Error(), convertedError.Message())
	assert.Equal(t, codes.DeadlineExceeded, convertedError.Code())

	cfg.Timeouts.EndpointExecutionTimeoutMS = oldTimeout
}

func TestFindByID_HappyPath(t *testing.T) {
	teacherUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, teacherUsername)

	formID := uuid.New()
	foundForm := util.RandomDomainForm()

	repository.EXPECT().
		FindByID(gomock.Any(), formID).
		Return(foundForm, nil)

	response, err := service.FindByID(ctx, formID)
	require.NoError(t, err)

	require.Equal(t, foundForm.ID, response.ID)
}
