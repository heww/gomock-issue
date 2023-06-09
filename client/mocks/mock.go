// Code generated by MockGen. DO NOT EDIT.
// Source: ./client.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	client "github.com/heww/gomock-issue/client"
	types "github.com/heww/gomock-issue/types"
)

// MockClient is a mock of Client interface.
type MockClient[T any, P types.Model[client.T]] struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder[T, P]
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder[T any, P types.Model[client.T]] struct {
	mock *MockClient[T, P]
}

// NewMockClient creates a new mock instance.
func NewMockClient[T any, P types.Model[client.T]](ctrl *gomock.Controller) *MockClient[T, P] {
	mock := &MockClient[T, P]{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder[T, P]{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient[T, P]) EXPECT() *MockClientMockRecorder[T, P] {
	return m.recorder
}

// Create mocks base method.
func (m *MockClient[T, P]) Create(arg0 *P) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockClientMockRecorder[T, P]) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockClient[T, P])(nil).Create), arg0)
}
