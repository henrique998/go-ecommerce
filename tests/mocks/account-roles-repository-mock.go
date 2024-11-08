// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/contracts/account-roles-repository.go
//
// Generated by this command:
//
//	mockgen -source=internal/app/contracts/account-roles-repository.go -destination=tests/mocks/account-roles-repository-mock.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	errors "github.com/henrique998/go-ecommerce/internal/app/errors"
	gomock "go.uber.org/mock/gomock"
)

// MockAccountRolesRepository is a mock of AccountRolesRepository interface.
type MockAccountRolesRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRolesRepositoryMockRecorder
}

// MockAccountRolesRepositoryMockRecorder is the mock recorder for MockAccountRolesRepository.
type MockAccountRolesRepositoryMockRecorder struct {
	mock *MockAccountRolesRepository
}

// NewMockAccountRolesRepository creates a new mock instance.
func NewMockAccountRolesRepository(ctrl *gomock.Controller) *MockAccountRolesRepository {
	mock := &MockAccountRolesRepository{ctrl: ctrl}
	mock.recorder = &MockAccountRolesRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRolesRepository) EXPECT() *MockAccountRolesRepositoryMockRecorder {
	return m.recorder
}

// FindByAccountId mocks base method.
func (m *MockAccountRolesRepository) FindByAccountId(id string) ([]string, errors.AppErr) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByAccountId", id)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(errors.AppErr)
	return ret0, ret1
}

// FindByAccountId indicates an expected call of FindByAccountId.
func (mr *MockAccountRolesRepositoryMockRecorder) FindByAccountId(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByAccountId", reflect.TypeOf((*MockAccountRolesRepository)(nil).FindByAccountId), id)
}
