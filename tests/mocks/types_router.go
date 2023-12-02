// Code generated by MockGen. DO NOT EDIT.
// Source: types/router.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	types "github.com/opzlabs/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)

// MockRouter is a mock of Router interface.
type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMockRecorder
}

// MockRouterMockRecorder is the mock recorder for MockRouter.
type MockRouterMockRecorder struct {
	mock *MockRouter
}

// NewMockRouter creates a new mock instance.
func NewMockRouter(ctrl *gomock.Controller) *MockRouter {
	mock := &MockRouter{ctrl: ctrl}
	mock.recorder = &MockRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRouter) EXPECT() *MockRouterMockRecorder {
	return m.recorder
}

// AddRoute mocks base method.
func (m *MockRouter) AddRoute(r types.Route) types.Router {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRoute", r)
	ret0, _ := ret[0].(types.Router)
	return ret0
}

// AddRoute indicates an expected call of AddRoute.
func (mr *MockRouterMockRecorder) AddRoute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoute", reflect.TypeOf((*MockRouter)(nil).AddRoute), r)
}

// Route mocks base method.
func (m *MockRouter) Route(ctx types.Context, path string) types.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Route", ctx, path)
	ret0, _ := ret[0].(types.Handler)
	return ret0
}

// Route indicates an expected call of Route.
func (mr *MockRouterMockRecorder) Route(ctx, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Route", reflect.TypeOf((*MockRouter)(nil).Route), ctx, path)
}

// MockQueryRouter is a mock of QueryRouter interface.
type MockQueryRouter struct {
	ctrl     *gomock.Controller
	recorder *MockQueryRouterMockRecorder
}

// MockQueryRouterMockRecorder is the mock recorder for MockQueryRouter.
type MockQueryRouterMockRecorder struct {
	mock *MockQueryRouter
}

// NewMockQueryRouter creates a new mock instance.
func NewMockQueryRouter(ctrl *gomock.Controller) *MockQueryRouter {
	mock := &MockQueryRouter{ctrl: ctrl}
	mock.recorder = &MockQueryRouterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockQueryRouter) EXPECT() *MockQueryRouterMockRecorder {
	return m.recorder
}

// AddRoute mocks base method.
func (m *MockQueryRouter) AddRoute(r string, h types.Querier) types.QueryRouter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRoute", r, h)
	ret0, _ := ret[0].(types.QueryRouter)
	return ret0
}

// AddRoute indicates an expected call of AddRoute.
func (mr *MockQueryRouterMockRecorder) AddRoute(r, h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoute", reflect.TypeOf((*MockQueryRouter)(nil).AddRoute), r, h)
}

// Route mocks base method.
func (m *MockQueryRouter) Route(path string) types.Querier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Route", path)
	ret0, _ := ret[0].(types.Querier)
	return ret0
}

// Route indicates an expected call of Route.
func (mr *MockQueryRouterMockRecorder) Route(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Route", reflect.TypeOf((*MockQueryRouter)(nil).Route), path)
}
