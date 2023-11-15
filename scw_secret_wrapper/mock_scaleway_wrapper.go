// Code generated by MockGen. DO NOT EDIT.
// Source: scaleway_wrapper.go
//
// Generated by this command:
//
//	mockgen -source scaleway_wrapper.go -destination mock_scaleway_wrapper.go -package scw_secret_manager
//
// Package scw_secret_manager is a generated GoMock package.
package scw_secret_manager

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockIScaleWayWrapper is a mock of IScaleWayWrapper interface.
type MockIScaleWayWrapper struct {
	ctrl     *gomock.Controller
	recorder *MockIScaleWayWrapperMockRecorder
}

// MockIScaleWayWrapperMockRecorder is the mock recorder for MockIScaleWayWrapper.
type MockIScaleWayWrapperMockRecorder struct {
	mock *MockIScaleWayWrapper
}

// NewMockIScaleWayWrapper creates a new mock instance.
func NewMockIScaleWayWrapper(ctrl *gomock.Controller) *MockIScaleWayWrapper {
	mock := &MockIScaleWayWrapper{ctrl: ctrl}
	mock.recorder = &MockIScaleWayWrapperMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIScaleWayWrapper) EXPECT() *MockIScaleWayWrapperMockRecorder {
	return m.recorder
}

// CreateNewSecretVersion mocks base method.
func (m *MockIScaleWayWrapper) CreateNewSecretVersion(secret Secret, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewSecretVersion", secret, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNewSecretVersion indicates an expected call of CreateNewSecretVersion.
func (mr *MockIScaleWayWrapperMockRecorder) CreateNewSecretVersion(secret, data any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewSecretVersion", reflect.TypeOf((*MockIScaleWayWrapper)(nil).CreateNewSecretVersion), secret, data)
}

// DeleteSecret mocks base method.
func (m *MockIScaleWayWrapper) DeleteSecret(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecret", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecret indicates an expected call of DeleteSecret.
func (mr *MockIScaleWayWrapperMockRecorder) DeleteSecret(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecret", reflect.TypeOf((*MockIScaleWayWrapper)(nil).DeleteSecret), id)
}

// DeleteSecretVersion mocks base method.
func (m *MockIScaleWayWrapper) DeleteSecretVersion(id, revision string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSecretVersion", id, revision)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteSecretVersion indicates an expected call of DeleteSecretVersion.
func (mr *MockIScaleWayWrapperMockRecorder) DeleteSecretVersion(id, revision any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSecretVersion", reflect.TypeOf((*MockIScaleWayWrapper)(nil).DeleteSecretVersion), id, revision)
}

// GetSecretData mocks base method.
func (m *MockIScaleWayWrapper) GetSecretData(secretName, revision string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSecretData", secretName, revision)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSecretData indicates an expected call of GetSecretData.
func (mr *MockIScaleWayWrapperMockRecorder) GetSecretData(secretName, revision any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSecretData", reflect.TypeOf((*MockIScaleWayWrapper)(nil).GetSecretData), secretName, revision)
}

// ListSecretVersions mocks base method.
func (m *MockIScaleWayWrapper) ListSecretVersions(secretID string) (SecretVersionHolder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListSecretVersions", secretID)
	ret0, _ := ret[0].(SecretVersionHolder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecretVersions indicates an expected call of ListSecretVersions.
func (mr *MockIScaleWayWrapperMockRecorder) ListSecretVersions(secretID any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecretVersions", reflect.TypeOf((*MockIScaleWayWrapper)(nil).ListSecretVersions), secretID)
}

// ListSecrets mocks base method.
func (m *MockIScaleWayWrapper) ListSecrets(names ...string) (SecretHolder, error) {
	m.ctrl.T.Helper()
	varargs := []any{}
	for _, a := range names {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListSecrets", varargs...)
	ret0, _ := ret[0].(SecretHolder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListSecrets indicates an expected call of ListSecrets.
func (mr *MockIScaleWayWrapperMockRecorder) ListSecrets(names ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListSecrets", reflect.TypeOf((*MockIScaleWayWrapper)(nil).ListSecrets), names...)
}

// SetSecret mocks base method.
func (m *MockIScaleWayWrapper) SetSecret(secretName string, secretValue []byte) (Secret, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSecret", secretName, secretValue)
	ret0, _ := ret[0].(Secret)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SetSecret indicates an expected call of SetSecret.
func (mr *MockIScaleWayWrapperMockRecorder) SetSecret(secretName, secretValue any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSecret", reflect.TypeOf((*MockIScaleWayWrapper)(nil).SetSecret), secretName, secretValue)
}
