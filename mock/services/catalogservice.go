// Code generated by MockGen. DO NOT EDIT.
// Source: ./services/catalogservice.go

// Package mock_services is a generated GoMock package.
package mock_services

import (
	reflect "reflect"

	models "github.com/aditya/ProjectCatalog/models"
	gomock "github.com/golang/mock/gomock"
)

// MockICatalogService is a mock of ICatalogService interface.
type MockICatalogService struct {
	ctrl     *gomock.Controller
	recorder *MockICatalogServiceMockRecorder
}

// MockICatalogServiceMockRecorder is the mock recorder for MockICatalogService.
type MockICatalogServiceMockRecorder struct {
	mock *MockICatalogService
}

// NewMockICatalogService creates a new mock instance.
func NewMockICatalogService(ctrl *gomock.Controller) *MockICatalogService {
	mock := &MockICatalogService{ctrl: ctrl}
	mock.recorder = &MockICatalogServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockICatalogService) EXPECT() *MockICatalogServiceMockRecorder {
	return m.recorder
}

// BuyProduct mocks base method.
func (m *MockICatalogService) BuyProduct(productId, productQuantity int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BuyProduct", productId, productQuantity)
	ret0, _ := ret[0].(error)
	return ret0
}

// BuyProduct indicates an expected call of BuyProduct.
func (mr *MockICatalogServiceMockRecorder) BuyProduct(productId, productQuantity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BuyProduct", reflect.TypeOf((*MockICatalogService)(nil).BuyProduct), productId, productQuantity)
}

// CreateProduct mocks base method.
func (m *MockICatalogService) CreateProduct(product models.UserProduct) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProduct", product)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProduct indicates an expected call of CreateProduct.
func (mr *MockICatalogServiceMockRecorder) CreateProduct(product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProduct", reflect.TypeOf((*MockICatalogService)(nil).CreateProduct), product)
}

// DeleteProduct mocks base method.
func (m *MockICatalogService) DeleteProduct(productId int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", productId)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockICatalogServiceMockRecorder) DeleteProduct(productId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockICatalogService)(nil).DeleteProduct), productId)
}

// ShowProductById mocks base method.
func (m *MockICatalogService) ShowProductById(productId int64) models.DbProduct {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowProductById", productId)
	ret0, _ := ret[0].(models.DbProduct)
	return ret0
}

// ShowProductById indicates an expected call of ShowProductById.
func (mr *MockICatalogServiceMockRecorder) ShowProductById(productId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowProductById", reflect.TypeOf((*MockICatalogService)(nil).ShowProductById), productId)
}

// ShowProducts mocks base method.
func (m *MockICatalogService) ShowProducts() []models.DbProduct {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ShowProducts")
	ret0, _ := ret[0].([]models.DbProduct)
	return ret0
}

// ShowProducts indicates an expected call of ShowProducts.
func (mr *MockICatalogServiceMockRecorder) ShowProducts() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ShowProducts", reflect.TypeOf((*MockICatalogService)(nil).ShowProducts))
}

// TopProduct mocks base method.
func (m *MockICatalogService) TopProduct() []models.Transaction {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TopProduct")
	ret0, _ := ret[0].([]models.Transaction)
	return ret0
}

// TopProduct indicates an expected call of TopProduct.
func (mr *MockICatalogServiceMockRecorder) TopProduct() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TopProduct", reflect.TypeOf((*MockICatalogService)(nil).TopProduct))
}

// UpdateProduct mocks base method.
func (m *MockICatalogService) UpdateProduct(arg0 models.DbProduct, arg1 int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProduct", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProduct indicates an expected call of UpdateProduct.
func (mr *MockICatalogServiceMockRecorder) UpdateProduct(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProduct", reflect.TypeOf((*MockICatalogService)(nil).UpdateProduct), arg0, arg1)
}
