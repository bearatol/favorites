// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_service is a generated GoMock package.
package mock_service

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPharmacies is a mock of Pharmacies interface.
type MockPharmacies struct {
	ctrl     *gomock.Controller
	recorder *MockPharmaciesMockRecorder
}

// MockPharmaciesMockRecorder is the mock recorder for MockPharmacies.
type MockPharmaciesMockRecorder struct {
	mock *MockPharmacies
}

// NewMockPharmacies creates a new mock instance.
func NewMockPharmacies(ctrl *gomock.Controller) *MockPharmacies {
	mock := &MockPharmacies{ctrl: ctrl}
	mock.recorder = &MockPharmaciesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPharmacies) EXPECT() *MockPharmaciesMockRecorder {
	return m.recorder
}

// DeletePharmacy mocks base method.
func (m *MockPharmacies) DeletePharmacy(ctx context.Context, user, pharmacy int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePharmacy", ctx, user, pharmacy)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePharmacy indicates an expected call of DeletePharmacy.
func (mr *MockPharmaciesMockRecorder) DeletePharmacy(ctx, user, pharmacy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePharmacy", reflect.TypeOf((*MockPharmacies)(nil).DeletePharmacy), ctx, user, pharmacy)
}

// SetPharmacy mocks base method.
func (m *MockPharmacies) SetPharmacy(ctx context.Context, user, pharmacy int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPharmacy", ctx, user, pharmacy)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPharmacy indicates an expected call of SetPharmacy.
func (mr *MockPharmaciesMockRecorder) SetPharmacy(ctx, user, pharmacy interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPharmacy", reflect.TypeOf((*MockPharmacies)(nil).SetPharmacy), ctx, user, pharmacy)
}

// MockProducts is a mock of Products interface.
type MockProducts struct {
	ctrl     *gomock.Controller
	recorder *MockProductsMockRecorder
}

// MockProductsMockRecorder is the mock recorder for MockProducts.
type MockProductsMockRecorder struct {
	mock *MockProducts
}

// NewMockProducts creates a new mock instance.
func NewMockProducts(ctrl *gomock.Controller) *MockProducts {
	mock := &MockProducts{ctrl: ctrl}
	mock.recorder = &MockProductsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockProducts) EXPECT() *MockProductsMockRecorder {
	return m.recorder
}

// DeleteProduct mocks base method.
func (m *MockProducts) DeleteProduct(ctx context.Context, user, product int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteProduct", ctx, user, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteProduct indicates an expected call of DeleteProduct.
func (mr *MockProductsMockRecorder) DeleteProduct(ctx, user, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteProduct", reflect.TypeOf((*MockProducts)(nil).DeleteProduct), ctx, user, product)
}

// SetProduct mocks base method.
func (m *MockProducts) SetProduct(ctx context.Context, user, product int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetProduct", ctx, user, product)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetProduct indicates an expected call of SetProduct.
func (mr *MockProductsMockRecorder) SetProduct(ctx, user, product interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetProduct", reflect.TypeOf((*MockProducts)(nil).SetProduct), ctx, user, product)
}
