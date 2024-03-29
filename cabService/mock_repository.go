// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package cabService is a generated GoMock package.
package cabService

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockCabRepository is a mock of CabRepository interface
type MockCabRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCabRepositoryMockRecorder
}

// MockCabRepositoryMockRecorder is the mock recorder for MockCabRepository
type MockCabRepositoryMockRecorder struct {
	mock *MockCabRepository
}

// NewMockCabRepository creates a new mock instance
func NewMockCabRepository(ctrl *gomock.Controller) *MockCabRepository {
	mock := &MockCabRepository{ctrl: ctrl}
	mock.recorder = &MockCabRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCabRepository) EXPECT() *MockCabRepositoryMockRecorder {
	return m.recorder
}

// GetNumberOfTripsByMedallionIds mocks base method
func (m *MockCabRepository) GetNumberOfTripsByMedallionIds(ctx context.Context, ids []string, date time.Time) ([]map[string]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNumberOfTripsByMedallionIds", ctx, ids, date)
	ret0, _ := ret[0].([]map[string]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNumberOfTripsByMedallionIds indicates an expected call of GetNumberOfTripsByMedallionIds
func (mr *MockCabRepositoryMockRecorder) GetNumberOfTripsByMedallionIds(ctx, ids, date interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNumberOfTripsByMedallionIds", reflect.TypeOf((*MockCabRepository)(nil).GetNumberOfTripsByMedallionIds), ctx, ids, date)
}
