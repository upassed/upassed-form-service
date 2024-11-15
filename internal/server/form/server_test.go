package form_test

import (
	"context"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/upassed/upassed-form-service/internal/config"
	"github.com/upassed/upassed-form-service/internal/handling"
	"github.com/upassed/upassed-form-service/internal/logging"
	"github.com/upassed/upassed-form-service/internal/middleware/common/auth"
	"github.com/upassed/upassed-form-service/internal/server"
	business "github.com/upassed/upassed-form-service/internal/service/model"
	"github.com/upassed/upassed-form-service/internal/util"
	"github.com/upassed/upassed-form-service/internal/util/mocks"
	"github.com/upassed/upassed-form-service/pkg/client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	"log"
	"os"
	"path/filepath"
	"testing"
)

var (
	formClient client.FormClient
	service    *mocks.FormService
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
		log.Fatal("cfg load error: ", err)
	}

	logger := logging.New(cfg.Env)
	ctrl := gomock.NewController(nil)
	defer ctrl.Finish()

	service = mocks.NewFormService(ctrl)
	authClient := mocks.NewAuthClientMW(ctrl)
	authClient.EXPECT().AuthenticationUnaryServerInterceptor().Return(emptyAuthMiddleware())

	formServer, err := server.New(server.AppServerCreateParams{
		Config:      cfg,
		Log:         logger,
		FormService: service,
		AuthClient:  authClient,
	})

	if err != nil {
		log.Fatal("server create error: ", err)
	}

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	cc, err := grpc.NewClient(fmt.Sprintf(":%s", cfg.GrpcServer.Port), opts...)
	if err != nil {
		log.Fatal("error creating client connection: ", err)
	}

	formClient = client.NewFormClient(cc)
	go func() {
		if err := formServer.Run(); err != nil {
			os.Exit(1)
		}
	}()

	exitCode := m.Run()
	formServer.GracefulStop()
	os.Exit(exitCode)
}

func TestFindByID_InvalidRequest(t *testing.T) {
	request := client.FormFindByIDRequest{
		FormId: "invalid_uuid",
	}

	_, err := formClient.FindByID(context.Background(), &request)
	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, codes.InvalidArgument, convertedError.Code())
}

func TestFindByID_ServiceError(t *testing.T) {
	request := client.FormFindByIDRequest{
		FormId: uuid.NewString(),
	}

	expectedServiceError := handling.Wrap(errors.New("some service error"), handling.WithCode(codes.NotFound))
	service.EXPECT().FindByID(gomock.Any(), uuid.MustParse(request.GetFormId())).Return(nil, expectedServiceError)

	_, err := formClient.FindByID(context.Background(), &request)
	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, "some service error", convertedError.Message())
	assert.Equal(t, codes.NotFound, convertedError.Code())
}

func TestFindByID_HappyPath(t *testing.T) {
	request := client.FormFindByIDRequest{
		FormId: uuid.NewString(),
	}

	foundForm := util.RandomBusinessForm()
	service.EXPECT().FindByID(gomock.Any(), uuid.MustParse(request.GetFormId())).Return(foundForm, nil)

	response, err := formClient.FindByID(context.Background(), &request)
	require.NoError(t, err)

	assert.Equal(t, foundForm.ID.String(), response.GetForm().GetId())
}

func TestFindByTeacherUsername_ServiceError(t *testing.T) {
	request := client.FormFindByTeacherUsernameRequest{}

	expectedServiceError := handling.Wrap(errors.New("some service error"), handling.WithCode(codes.NotFound))
	service.EXPECT().FindByTeacherUsername(gomock.Any(), gomock.Any()).Return(nil, expectedServiceError)

	_, err := formClient.FindByTeacherUsername(context.Background(), &request)
	require.Error(t, err)

	convertedError := status.Convert(err)
	assert.Equal(t, "some service error", convertedError.Message())
	assert.Equal(t, codes.NotFound, convertedError.Code())
}

func TestFindByTeacherUsername_HappyPath(t *testing.T) {
	request := client.FormFindByTeacherUsernameRequest{}

	foundForms := []*business.Form{util.RandomBusinessForm(), util.RandomBusinessForm()}
	service.EXPECT().FindByTeacherUsername(gomock.Any(), gomock.Any()).Return(foundForms, nil)

	response, err := formClient.FindByTeacherUsername(context.Background(), &request)
	require.NoError(t, err)

	for idx, form := range foundForms {
		assert.Equal(t, form.ID.String(), response.GetFoundForms()[idx].GetId())
	}
}

func emptyAuthMiddleware() func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		ctxWithUsername := context.WithValue(ctx, auth.UsernameKey, "teacherUsername")

		return handler(ctxWithUsername, req)
	}
}
