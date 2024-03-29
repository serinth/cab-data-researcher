// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package cacheService is a generated GoMock package.
package cacheService

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCacheRepository is a mock of CacheRepository interface
type MockCacheRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCacheRepositoryMockRecorder
}

// MockCacheRepositoryMockRecorder is the mock recorder for MockCacheRepository
type MockCacheRepositoryMockRecorder struct {
	mock *MockCacheRepository
}

// NewMockCacheRepository creates a new mock instance
func NewMockCacheRepository(ctrl *gomock.Controller) *MockCacheRepository {
	mock := &MockCacheRepository{ctrl: ctrl}
	mock.recorder = &MockCacheRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCacheRepository) EXPECT() *MockCacheRepositoryMockRecorder {
	return m.recorder
}

// SaveExpiringKeyValue mocks base method
func (m *MockCacheRepository) SaveExpiringKeyValue(ctx context.Context, key string, value, expireSeconds int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveExpiringKeyValue", ctx, key, value, expireSeconds)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveExpiringKeyValue indicates an expected call of SaveExpiringKeyValue
func (mr *MockCacheRepositoryMockRecorder) SaveExpiringKeyValue(ctx, key, value, expireSeconds interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveExpiringKeyValue", reflect.TypeOf((*MockCacheRepository)(nil).SaveExpiringKeyValue), ctx, key, value, expireSeconds)
}

// GetKey mocks base method
func (m *MockCacheRepository) GetKey(ctx context.Context, key string) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetKey", ctx, key)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetKey indicates an expected call of GetKey
func (mr *MockCacheRepositoryMockRecorder) GetKey(ctx, key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetKey", reflect.TypeOf((*MockCacheRepository)(nil).GetKey), ctx, key)
}
