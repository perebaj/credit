// Code generated by MockGen. DO NOT EDIT.
// Source: http.go
//
// Generated by this command:
//
//	mockgen -source http.go -destination ../mock/http_mock.go -package mock
//

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	credit "github.com/perebaj/credit"
	gomock "go.uber.org/mock/gomock"
)

// MockCompanyService is a mock of CompanyService interface.
type MockCompanyService struct {
	ctrl     *gomock.Controller
	recorder *MockCompanyServiceMockRecorder
}

// MockCompanyServiceMockRecorder is the mock recorder for MockCompanyService.
type MockCompanyServiceMockRecorder struct {
	mock *MockCompanyService
}

// NewMockCompanyService creates a new mock instance.
func NewMockCompanyService(ctrl *gomock.Controller) *MockCompanyService {
	mock := &MockCompanyService{ctrl: ctrl}
	mock.recorder = &MockCompanyServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCompanyService) EXPECT() *MockCompanyServiceMockRecorder {
	return m.recorder
}

// SaveCompany mocks base method.
func (m *MockCompanyService) SaveCompany(ctx context.Context, company credit.Company) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveCompany", ctx, company)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveCompany indicates an expected call of SaveCompany.
func (mr *MockCompanyServiceMockRecorder) SaveCompany(ctx, company any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveCompany", reflect.TypeOf((*MockCompanyService)(nil).SaveCompany), ctx, company)
}