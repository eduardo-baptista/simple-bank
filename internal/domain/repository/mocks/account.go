// Code generated by MockGen. DO NOT EDIT.
// Source: account.go
//
// Generated by this command:
//
//	mockgen -source=account.go -destination=mocks/account.go -package=mocks AccountRepository
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"
	entity "simple-bank/internal/domain/entity"

	gomock "go.uber.org/mock/gomock"
)

// MockAccountRepository is a mock of AccountRepository interface.
type MockAccountRepository struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRepositoryMockRecorder
}

// MockAccountRepositoryMockRecorder is the mock recorder for MockAccountRepository.
type MockAccountRepositoryMockRecorder struct {
	mock *MockAccountRepository
}

// NewMockAccountRepository creates a new mock instance.
func NewMockAccountRepository(ctrl *gomock.Controller) *MockAccountRepository {
	mock := &MockAccountRepository{ctrl: ctrl}
	mock.recorder = &MockAccountRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountRepository) EXPECT() *MockAccountRepositoryMockRecorder {
	return m.recorder
}

// DeleteAllAccounts mocks base method.
func (m *MockAccountRepository) DeleteAllAccounts() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteAllAccounts")
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteAllAccounts indicates an expected call of DeleteAllAccounts.
func (mr *MockAccountRepositoryMockRecorder) DeleteAllAccounts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAllAccounts", reflect.TypeOf((*MockAccountRepository)(nil).DeleteAllAccounts))
}

// GetAccountByID mocks base method.
func (m *MockAccountRepository) GetAccountByID(id string) (*entity.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountByID", id)
	ret0, _ := ret[0].(*entity.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAccountByID indicates an expected call of GetAccountByID.
func (mr *MockAccountRepositoryMockRecorder) GetAccountByID(id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountByID", reflect.TypeOf((*MockAccountRepository)(nil).GetAccountByID), id)
}

// SaveAccount mocks base method.
func (m *MockAccountRepository) SaveAccount(account *entity.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAccount", account)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveAccount indicates an expected call of SaveAccount.
func (mr *MockAccountRepositoryMockRecorder) SaveAccount(account any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAccount", reflect.TypeOf((*MockAccountRepository)(nil).SaveAccount), account)
}

// UpdateAccount mocks base method.
func (m *MockAccountRepository) UpdateAccount(account *entity.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateAccount", account)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateAccount indicates an expected call of UpdateAccount.
func (mr *MockAccountRepositoryMockRecorder) UpdateAccount(account any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAccount", reflect.TypeOf((*MockAccountRepository)(nil).UpdateAccount), account)
}
