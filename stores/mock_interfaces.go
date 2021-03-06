// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go

// Package stores is a generated GoMock package.
package stores

import (
	gomock "github.com/golang/mock/gomock"
	models "goApi/models"
	reflect "reflect"
)

// MockProduct is a mock of Product interface
type MockProduct struct {
	ctrl     *gomock.Controller
	recorder *MockProductMockRecorder
}

// MockProductMockRecorder is the mock recorder for MockProduct
type MockProductMockRecorder struct {
	mock *MockProduct
}

// NewMockProduct creates a new mock instance
func NewMockProduct(ctrl *gomock.Controller) *MockProduct {
	mock := &MockProduct{ctrl: ctrl}
	mock.recorder = &MockProductMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockProduct) EXPECT() *MockProductMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockProduct) Create(product models.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", product)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create
func (mr *MockProductMockRecorder) Create(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProduct)(nil).Create), product)
}

// Read mocks base method
func (m *MockProduct) Read() ([]models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read")
	ret0, _ := ret[0].([]models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read
func (mr *MockProductMockRecorder) Read() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockProduct)(nil).Read))
}

// Update mocks base method
func (m *MockProduct) Update(price string, id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", price, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockProductMockRecorder) Update(price, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockProduct)(nil).Update), price, id)
}

// Delete mocks base method
func (m *MockProduct) Delete(id int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockProductMockRecorder) Delete(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockProduct)(nil).Delete), id)
}
