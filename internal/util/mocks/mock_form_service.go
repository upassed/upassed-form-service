// Code generated by MockGen. DO NOT EDIT.
// Source: internal/service/form/service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	domain "github.com/upassed/upassed-form-service/internal/repository/model"
	business "github.com/upassed/upassed-form-service/internal/service/model"
)

// FormService is a mock of Service interface.
type FormService struct {
	ctrl     *gomock.Controller
	recorder *FormServiceMockRecorder
}

// FormServiceMockRecorder is the mock recorder for FormService.
type FormServiceMockRecorder struct {
	mock *FormService
}

// NewFormService creates a new mock instance.
func NewFormService(ctrl *gomock.Controller) *FormService {
	mock := &FormService{ctrl: ctrl}
	mock.recorder = &FormServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *FormService) EXPECT() *FormServiceMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *FormService) Create(ctx context.Context, form *business.Form) (*business.FormCreateResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, form)
	ret0, _ := ret[0].(*business.FormCreateResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *FormServiceMockRecorder) Create(ctx, form interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*FormService)(nil).Create), ctx, form)
}

// FindByID mocks base method.
func (m *FormService) FindByID(ctx context.Context, formID uuid.UUID) (*business.Form, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, formID)
	ret0, _ := ret[0].(*business.Form)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *FormServiceMockRecorder) FindByID(ctx, formID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*FormService)(nil).FindByID), ctx, formID)
}

// MockformRepository is a mock of formRepository interface.
type MockformRepository struct {
	ctrl     *gomock.Controller
	recorder *MockformRepositoryMockRecorder
}

// MockformRepositoryMockRecorder is the mock recorder for MockformRepository.
type MockformRepositoryMockRecorder struct {
	mock *MockformRepository
}

// NewMockformRepository creates a new mock instance.
func NewMockformRepository(ctrl *gomock.Controller) *MockformRepository {
	mock := &MockformRepository{ctrl: ctrl}
	mock.recorder = &MockformRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockformRepository) EXPECT() *MockformRepositoryMockRecorder {
	return m.recorder
}

// ExistsByNameAndTeacherUsername mocks base method.
func (m *MockformRepository) ExistsByNameAndTeacherUsername(ctx context.Context, formName, teacherUsername string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExistsByNameAndTeacherUsername", ctx, formName, teacherUsername)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ExistsByNameAndTeacherUsername indicates an expected call of ExistsByNameAndTeacherUsername.
func (mr *MockformRepositoryMockRecorder) ExistsByNameAndTeacherUsername(ctx, formName, teacherUsername interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExistsByNameAndTeacherUsername", reflect.TypeOf((*MockformRepository)(nil).ExistsByNameAndTeacherUsername), ctx, formName, teacherUsername)
}

// FindByID mocks base method.
func (m *MockformRepository) FindByID(ctx context.Context, formID uuid.UUID) (*domain.Form, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByID", ctx, formID)
	ret0, _ := ret[0].(*domain.Form)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByID indicates an expected call of FindByID.
func (mr *MockformRepositoryMockRecorder) FindByID(ctx, formID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByID", reflect.TypeOf((*MockformRepository)(nil).FindByID), ctx, formID)
}

// Save mocks base method.
func (m *MockformRepository) Save(ctx context.Context, form *domain.Form) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", ctx, form)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockformRepositoryMockRecorder) Save(ctx, form interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockformRepository)(nil).Save), ctx, form)
}