// Code generated by MockGen. DO NOT EDIT.
// Source: interfaces.go
//
// Generated by this command:
//
//	mockgen -source=interfaces.go -destination=mock_interfaces.go -package=services
//

// Package services is a generated GoMock package.
package services

import (
	reflect "reflect"

	models "github.com/aditi-zs/test-repo/models"
	gomock "go.uber.org/mock/gomock"
	gofr "gofr.dev/pkg/gofr"
)

// MockProduct is a mock of Product interface.
type MockProduct struct {
	ctrl     *gomock.Controller
	recorder *MockProductMockRecorder
}

// MockProductMockRecorder is the mock recorder for MockProduct.
type MockProductMockRecorder struct {
	mock *MockProduct
}

// NewMockProduct creates a new mock instance.
func NewMockProduct(ctrl *gomock.Controller) *MockProduct {
	mock := &MockProduct{ctrl: ctrl}
	mock.recorder = &MockProductMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProduct) EXPECT() *MockProductMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockProduct) Create(ctx *gofr.Context, product *models.Product) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// Create indicates an expected call of Create.
func (mr *MockProductMockRecorder) Create(ctx, product any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockProduct)(nil).Create), ctx, product)
}

// GetAll mocks base method.
func (m *MockProduct) GetAll(ctx *gofr.Context, f string) ([]*models.Product, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, f)
	ret0, _ := ret[0].([]*models.Product)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockProductMockRecorder) GetAll(ctx, f any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockProduct)(nil).GetAll), ctx, f)
}
