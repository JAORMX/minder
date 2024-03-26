// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stacklok/minder/internal/marketplaces/subscriptions (interfaces: SubscriptionService)
//
// Generated by this command:
//
//	mockgen -package mocksubscription -destination internal/marketplaces/subscriptions/mock/subscription.go github.com/stacklok/minder/internal/marketplaces/subscriptions SubscriptionService
//

// Package mocksubscription is a generated GoMock package.
package mocksubscription

import (
	context "context"
	reflect "reflect"

	db "github.com/stacklok/minder/internal/db"
	types "github.com/stacklok/minder/internal/marketplaces/types"
	reader "github.com/stacklok/minder/pkg/mindpak/reader"
	gomock "go.uber.org/mock/gomock"
)

// MockSubscriptionService is a mock of SubscriptionService interface.
type MockSubscriptionService struct {
	ctrl     *gomock.Controller
	recorder *MockSubscriptionServiceMockRecorder
}

// MockSubscriptionServiceMockRecorder is the mock recorder for MockSubscriptionService.
type MockSubscriptionServiceMockRecorder struct {
	mock *MockSubscriptionService
}

// NewMockSubscriptionService creates a new mock instance.
func NewMockSubscriptionService(ctrl *gomock.Controller) *MockSubscriptionService {
	mock := &MockSubscriptionService{ctrl: ctrl}
	mock.recorder = &MockSubscriptionServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSubscriptionService) EXPECT() *MockSubscriptionServiceMockRecorder {
	return m.recorder
}

// CreateProfile mocks base method.
func (m *MockSubscriptionService) CreateProfile(arg0 context.Context, arg1 types.ProjectContext, arg2 reader.BundleReader, arg3 string, arg4 db.ExtendQuerier) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateProfile", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateProfile indicates an expected call of CreateProfile.
func (mr *MockSubscriptionServiceMockRecorder) CreateProfile(arg0, arg1, arg2, arg3, arg4 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateProfile", reflect.TypeOf((*MockSubscriptionService)(nil).CreateProfile), arg0, arg1, arg2, arg3, arg4)
}

// Subscribe mocks base method.
func (m *MockSubscriptionService) Subscribe(arg0 context.Context, arg1 types.ProjectContext, arg2 reader.BundleReader, arg3 db.ExtendQuerier) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockSubscriptionServiceMockRecorder) Subscribe(arg0, arg1, arg2, arg3 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockSubscriptionService)(nil).Subscribe), arg0, arg1, arg2, arg3)
}
