// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/oxyno-zeta/s3-proxy/pkg/s3-proxy/s3client (interfaces: Manager)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	s3client "github.com/oxyno-zeta/s3-proxy/pkg/s3-proxy/s3client"
)

// MockManager is a mock of Manager interface.
type MockManager struct {
	ctrl     *gomock.Controller
	recorder *MockManagerMockRecorder
}

// MockManagerMockRecorder is the mock recorder for MockManager.
type MockManagerMockRecorder struct {
	mock *MockManager
}

// NewMockManager creates a new mock instance.
func NewMockManager(ctrl *gomock.Controller) *MockManager {
	mock := &MockManager{ctrl: ctrl}
	mock.recorder = &MockManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockManager) EXPECT() *MockManagerMockRecorder {
	return m.recorder
}

// GetClientForTarget mocks base method.
func (m *MockManager) GetClientForTarget(arg0 string) s3client.Client {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetClientForTarget", arg0)
	ret0, _ := ret[0].(s3client.Client)
	return ret0
}

// GetClientForTarget indicates an expected call of GetClientForTarget.
func (mr *MockManagerMockRecorder) GetClientForTarget(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetClientForTarget", reflect.TypeOf((*MockManager)(nil).GetClientForTarget), arg0)
}

// Load mocks base method.
func (m *MockManager) Load() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load")
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load.
func (mr *MockManagerMockRecorder) Load() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockManager)(nil).Load))
}
