package form_test

import (
	"context"
	"errors"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"github.com/upassed/upassed-form-service/internal/config"
	"github.com/upassed/upassed-form-service/internal/logging"
	"github.com/upassed/upassed-form-service/internal/middleware/common/auth"
	domain "github.com/upassed/upassed-form-service/internal/repository/model"
	"github.com/upassed/upassed-form-service/internal/service/form"
	"github.com/upassed/upassed-form-service/internal/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"path/filepath"
	"testing"
)

type mockFormRepository struct {
	mock.Mock
}

func (m *mockFormRepository) ExistsByNameAndTeacherUsername(ctx context.Context, formName, teacherUsername string) (bool, error) {
	args := m.Called(ctx, formName, teacherUsername)
	return args.Bool(0), args.Error(1)
}

func (m *mockFormRepository) Save(ctx context.Context, form *domain.Form) error {
	args := m.Called(ctx, form)
	return args.Error(0)
}

var (
	cfg *config.Config
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

	exitCode := m.Run()
	os.Exit(exitCode)
}

func TestCreate_ErrorCheckingFormExists(t *testing.T) {
	teacherUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, teacherUsername)

	formRepository := new(mockFormRepository)

	formToCreate := util.RandomBusinessForm()
	expectedRepositoryError := errors.New("some repo error")
	formRepository.On(
		"ExistsByNameAndTeacherUsername",
		mock.Anything,
		formToCreate.Name,
		teacherUsername,
	).Return(false, expectedRepositoryError)

	service := form.New(cfg, logging.New(config.EnvTesting), formRepository)
	_, err := service.Create(ctx, formToCreate)
	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, expectedRepositoryError.Error(), convertedError.Message())
	assert.Equal(t, codes.Internal, convertedError.Code())
}

func TestCreate_DuplicateExists(t *testing.T) {
	teacherUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, teacherUsername)

	formRepository := new(mockFormRepository)

	formToCreate := util.RandomBusinessForm()
	formRepository.On(
		"ExistsByNameAndTeacherUsername",
		mock.Anything,
		formToCreate.Name,
		teacherUsername,
	).Return(true, nil)

	service := form.New(cfg, logging.New(config.EnvTesting), formRepository)
	_, err := service.Create(ctx, formToCreate)
	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, "form duplicate found", convertedError.Message())
	assert.Equal(t, codes.AlreadyExists, convertedError.Code())
}

func TestCreate_ErrorSavingForm(t *testing.T) {
	teacherUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, teacherUsername)

	formRepository := new(mockFormRepository)

	formToCreate := util.RandomBusinessForm()
	formRepository.On(
		"ExistsByNameAndTeacherUsername",
		mock.Anything,
		formToCreate.Name,
		teacherUsername,
	).Return(false, nil)

	expectedRepoError := errors.New("some repo error")
	formRepository.On(
		"Save",
		mock.Anything,
		mock.Anything,
	).Return(expectedRepoError)

	service := form.New(cfg, logging.New(config.EnvTesting), formRepository)
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

	formRepository := new(mockFormRepository)

	formToCreate := util.RandomBusinessForm()
	formRepository.On(
		"ExistsByNameAndTeacherUsername",
		mock.Anything,
		formToCreate.Name,
		teacherUsername,
	).Return(false, nil)

	formRepository.On(
		"Save",
		mock.Anything,
		mock.Anything,
	).Return(nil)

	service := form.New(cfg, logging.New(config.EnvTesting), formRepository)
	_, err := service.Create(ctx, formToCreate)
	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, form.ErrFormCreateDeadlineExceeded.Error(), convertedError.Message())
	assert.Equal(t, codes.DeadlineExceeded, convertedError.Code())

	cfg.Timeouts.EndpointExecutionTimeoutMS = oldTimeout
}

func TestCreate_HappyPath(t *testing.T) {
	teacherUsername := gofakeit.Username()
	ctx := context.WithValue(context.Background(), auth.UsernameKey, teacherUsername)

	formRepository := new(mockFormRepository)

	formToCreate := util.RandomBusinessForm()
	formRepository.On(
		"ExistsByNameAndTeacherUsername",
		mock.Anything,
		formToCreate.Name,
		teacherUsername,
	).Return(false, nil)

	formRepository.On(
		"Save",
		mock.Anything,
		mock.Anything,
	).Return(nil)

	service := form.New(cfg, logging.New(config.EnvTesting), formRepository)
	response, err := service.Create(ctx, formToCreate)
	require.NoError(t, err)

	require.NotNil(t, response.CreatedFormID)
}
