package form_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/upassed/upassed-form-service/internal/config"
	"github.com/upassed/upassed-form-service/internal/logging"
	"github.com/upassed/upassed-form-service/internal/middleware/common/auth"
	"github.com/upassed/upassed-form-service/internal/repository"
	"github.com/upassed/upassed-form-service/internal/repository/form"
	"github.com/upassed/upassed-form-service/internal/util"
	"github.com/upassed/upassed-form-service/testcontainer"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"testing"
	"time"
)

var (
	studentRepository form.Repository
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

	cfg, err := config.Load()
	if err != nil {
		log.Fatal("unable to parse cfg: ", err)
	}

	ctx := context.Background()
	postgresTestcontainer, err := testcontainer.NewPostgresTestcontainer(ctx)
	if err != nil {
		log.Fatal("unable to create a testcontainer: ", err)
	}

	port, err := postgresTestcontainer.Start(ctx)
	if err != nil {
		log.Fatal("unable to get a postgres testcontainer real port: ", err)
	}

	cfg.Storage.Port = strconv.Itoa(port)
	logger := logging.New(cfg.Env)
	if err := postgresTestcontainer.Migrate(cfg, logger); err != nil {
		log.Fatal("unable to run migrations: ", err)
	}

	db, err := repository.OpenGormDbConnection(cfg, logger)
	if err != nil {
		log.Fatal("unable to open connection to postgres: ", err)
	}

	studentRepository = form.New(db, cfg, logger)
	exitCode := m.Run()
	if err := postgresTestcontainer.Stop(ctx); err != nil {
		log.Println("unable to stop postgres testcontainer: ", err)
	}

	os.Exit(exitCode)
}

func TestSave_DeadlineExceededWhileSaving(t *testing.T) {
	formToSave := util.RandomDomainForm()

	ctx, cancel := context.WithTimeout(
		context.WithValue(
			context.Background(), auth.UsernameKey, formToSave.TeacherUsername,
		),
		1*time.Nanosecond,
	)
	defer cancel()

	err := studentRepository.Save(ctx, formToSave)
	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, form.ErrSavingForm.Error(), convertedError.Message())
	assert.Equal(t, codes.Internal, convertedError.Code())
}

func TestSave_HappyPath(t *testing.T) {
	formToSave := util.RandomDomainForm()

	ctx := context.WithValue(context.Background(), auth.UsernameKey, formToSave.TeacherUsername)
	err := studentRepository.Save(ctx, formToSave)
	require.NoError(t, err)
}