// Code generated by MockGen. DO NOT EDIT.
// Source: BucketRepository.go

// Package mocks is a generated GoMock package.
package mocks

import (
	entities "github.com/Rototot/anti-brute-force/pkg/domain/entities"
	valueObjects "github.com/Rototot/anti-brute-force/pkg/domain/valueObjects"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBucketRepository is a mock of BucketRepository interface
type MockBucketRepository struct {
	ctrl     *gomock.Controller
	recorder *MockBucketRepositoryMockRecorder
}

// MockBucketRepositoryMockRecorder is the mock recorder for MockBucketRepository
type MockBucketRepositoryMockRecorder struct {
	mock *MockBucketRepository
}

// NewMockBucketRepository creates a new mock instance
func NewMockBucketRepository(ctrl *gomock.Controller) *MockBucketRepository {
	mock := &MockBucketRepository{ctrl: ctrl}
	mock.recorder = &MockBucketRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBucketRepository) EXPECT() *MockBucketRepositoryMockRecorder {
	return m.recorder
}

// FindOneByID mocks base method
func (m *MockBucketRepository) FindOneByID(id valueObjects.BucketID) (*entities.Bucket, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOneByID", id)
	ret0, _ := ret[0].(*entities.Bucket)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOneByID indicates an expected call of FindOneByID
func (mr *MockBucketRepositoryMockRecorder) FindOneByID(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOneByID", reflect.TypeOf((*MockBucketRepository)(nil).FindOneByID), id)
}

// Add mocks base method
func (m *MockBucketRepository) Add(bucket *entities.Bucket) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", bucket)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockBucketRepositoryMockRecorder) Add(bucket interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockBucketRepository)(nil).Add), bucket)
}

// Update mocks base method
func (m *MockBucketRepository) Update(bucket *entities.Bucket) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", bucket)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockBucketRepositoryMockRecorder) Update(bucket interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockBucketRepository)(nil).Update), bucket)
}

// Remove mocks base method
func (m *MockBucketRepository) Remove(bucket *entities.Bucket) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Remove", bucket)
	ret0, _ := ret[0].(error)
	return ret0
}

// Remove indicates an expected call of Remove
func (mr *MockBucketRepositoryMockRecorder) Remove(bucket interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockBucketRepository)(nil).Remove), bucket)
}
