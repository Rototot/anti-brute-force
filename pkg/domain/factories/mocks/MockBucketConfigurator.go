// Code generated by MockGen. DO NOT EDIT.
// Source: BucketConfigurator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBucketConfigurator is a mock of BucketConfigurator interface
type MockBucketConfigurator struct {
	ctrl     *gomock.Controller
	recorder *MockBucketConfiguratorMockRecorder
}

// MockBucketConfiguratorMockRecorder is the mock recorder for MockBucketConfigurator
type MockBucketConfiguratorMockRecorder struct {
	mock *MockBucketConfigurator
}

// NewMockBucketConfigurator creates a new mock instance
func NewMockBucketConfigurator(ctrl *gomock.Controller) *MockBucketConfigurator {
	mock := &MockBucketConfigurator{ctrl: ctrl}
	mock.recorder = &MockBucketConfiguratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBucketConfigurator) EXPECT() *MockBucketConfiguratorMockRecorder {
	return m.recorder
}

// IpBucketCapacity mocks base method
func (m *MockBucketConfigurator) IPBucketCapacity() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IPBucketCapacity")
	ret0, _ := ret[0].(int)
	return ret0
}

// IpBucketCapacity indicates an expected call of IpBucketCapacity
func (mr *MockBucketConfiguratorMockRecorder) IpBucketCapacity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IPBucketCapacity", reflect.TypeOf((*MockBucketConfigurator)(nil).IPBucketCapacity))
}

// LoginBucketCapacity mocks base method
func (m *MockBucketConfigurator) LoginBucketCapacity() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginBucketCapacity")
	ret0, _ := ret[0].(int)
	return ret0
}

// LoginBucketCapacity indicates an expected call of LoginBucketCapacity
func (mr *MockBucketConfiguratorMockRecorder) LoginBucketCapacity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginBucketCapacity", reflect.TypeOf((*MockBucketConfigurator)(nil).LoginBucketCapacity))
}

// PasswordBucketCapacity mocks base method
func (m *MockBucketConfigurator) PasswordBucketCapacity() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PasswordBucketCapacity")
	ret0, _ := ret[0].(int)
	return ret0
}

// PasswordBucketCapacity indicates an expected call of PasswordBucketCapacity
func (mr *MockBucketConfiguratorMockRecorder) PasswordBucketCapacity() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PasswordBucketCapacity", reflect.TypeOf((*MockBucketConfigurator)(nil).PasswordBucketCapacity))
}
