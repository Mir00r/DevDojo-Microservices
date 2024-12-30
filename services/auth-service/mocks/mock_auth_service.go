// Code generated by MockGen. DO NOT EDIT.
// Source: internal/services/auth_service.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	dtos "github.com/Mir00r/auth-service/internal/models/dtos"
	entities "github.com/Mir00r/auth-service/internal/models/entities"
	services "github.com/Mir00r/auth-service/internal/models/request"
	gomock "github.com/golang/mock/gomock"
)

// MockAuthService is a mock of AuthService interface.
type MockAuthService struct {
	ctrl     *gomock.Controller
	recorder *MockAuthServiceMockRecorder
}

// MockAuthServiceMockRecorder is the mock recorder for MockAuthService.
type MockAuthServiceMockRecorder struct {
	mock *MockAuthService
}

// NewMockAuthService creates a new mock instance.
func NewMockAuthService(ctrl *gomock.Controller) *MockAuthService {
	mock := &MockAuthService{ctrl: ctrl}
	mock.recorder = &MockAuthServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthService) EXPECT() *MockAuthServiceMockRecorder {
	return m.recorder
}

// Authenticate mocks base method.
func (m *MockAuthService) Authenticate(req dtos.LoginRequest) (*dtos.LoginResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Authenticate", req)
	ret0, _ := ret[0].(*dtos.LoginResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Authenticate indicates an expected call of Authenticate.
func (mr *MockAuthServiceMockRecorder) Authenticate(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Authenticate", reflect.TypeOf((*MockAuthService)(nil).Authenticate), req)
}

// GetUserProfile mocks base method.
func (m *MockAuthService) GetUserProfile(userID string) (*entities.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserProfile", userID)
	ret0, _ := ret[0].(*entities.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserProfile indicates an expected call of GetUserProfile.
func (mr *MockAuthServiceMockRecorder) GetUserProfile(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserProfile", reflect.TypeOf((*MockAuthService)(nil).GetUserProfile), userID)
}

// RegisterUser mocks base method.
func (m *MockAuthService) RegisterUser(req services.RegisterRequest) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterUser", req)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterUser indicates an expected call of RegisterUser.
func (mr *MockAuthServiceMockRecorder) RegisterUser(req interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterUser", reflect.TypeOf((*MockAuthService)(nil).RegisterUser), req)
}
