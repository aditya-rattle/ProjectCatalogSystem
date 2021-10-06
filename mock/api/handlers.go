// Code generated by MockGen. DO NOT EDIT.
// Source: ./api/handlers.go

// Package mock_api is a generated GoMock package.
package mock_api

import (
	http "net/http"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockIhandlers is a mock of Ihandlers interface.
type MockIhandlers struct {
	ctrl     *gomock.Controller
	recorder *MockIhandlersMockRecorder
}

// MockIhandlersMockRecorder is the mock recorder for MockIhandlers.
type MockIhandlersMockRecorder struct {
	mock *MockIhandlers
}

// NewMockIhandlers creates a new mock instance.
func NewMockIhandlers(ctrl *gomock.Controller) *MockIhandlers {
	mock := &MockIhandlers{ctrl: ctrl}
	mock.recorder = &MockIhandlersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIhandlers) EXPECT() *MockIhandlersMockRecorder {
	return m.recorder
}

// BuyProduct mocks base method.
func (m *MockIhandlers) BuyProduct(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BuyProduct", w, r)
}

// BuyProduct indicates an expected call of BuyProduct.
func (mr *MockIhandlersMockRecorder) BuyProduct(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuyProduct", reflect.TypeOf((*MockIhandlers)(nil).BuyProduct), w, r)
}

// CreateProduct mocks base method.
func (m *MockIhandlers) CreateProduct(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateProduct", w, r)
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockIhandlersMockRecorder) CreateProduct(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockIhandlers)(nil).CreateProduct), w, r)
}

// DeleteProduct mocks base method.
func (m *MockIhandlers) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteProduct", w, r)
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockIhandlersMockRecorder) DeleteProduct(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockIhandlers)(nil).DeleteProduct), w, r)
}

// ShowProductById mocks base method.
func (m *MockIhandlers) ShowProductById(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ShowProductById", w, r)
}

// ShowProductById indicates an expected call of ShowProductById.
func (mr *MockIhandlersMockRecorder) ShowProductById(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowProductById", reflect.TypeOf((*MockIhandlers)(nil).ShowProductById), w, r)
}

// ShowProducts mocks base method.
func (m *MockIhandlers) ShowProducts(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ShowProducts", w, r)
}

// ShowProducts indicates an expected call of ShowProducts.
func (mr *MockIhandlersMockRecorder) ShowProducts(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowProducts", reflect.TypeOf((*MockIhandlers)(nil).ShowProducts), w, r)
}

// UpdateProduct mocks base method.
func (m *MockIhandlers) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateProduct", w, r)
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockIhandlersMockRecorder) UpdateProduct(w, r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockIhandlers)(nil).UpdateProduct), w, r)
}