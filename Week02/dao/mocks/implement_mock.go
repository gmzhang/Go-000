// Code generated by MockGen. DO NOT EDIT.
// Source: dao/interface.go

// Package mocks is a generated GoMock package.
package mocks

import (
	model "github.com/gmzhang/Go-000/Week02/model"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockDao is a mock of Dao interface
type MockDao struct {
	ctrl     *gomock.Controller
	recorder *MockDaoMockRecorder
}

// MockDaoMockRecorder is the mock recorder for MockDao
type MockDaoMockRecorder struct {
	mock *MockDao
}

// NewMockDao creates a new mock instance
func NewMockDao(ctrl *gomock.Controller) *MockDao {
	mock := &MockDao{ctrl: ctrl}
	mock.recorder = &MockDaoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockDao) EXPECT() *MockDaoMockRecorder {
	return m.recorder
}

// GetUserById mocks base method
func (m *MockDao) GetUserById(id uint) (model.User, error) {
	ret := m.ctrl.Call(m, "GetUserById", id)
	ret0, _ := ret[0].(model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserById indicates an expected call of GetUserById
func (mr *MockDaoMockRecorder) GetUserById(id interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserById", reflect.TypeOf((*MockDao)(nil).GetUserById), id)
}
