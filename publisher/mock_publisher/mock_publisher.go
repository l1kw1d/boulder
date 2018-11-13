// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/letsencrypt/boulder/core (interfaces: Publisher)

// Package mock_publisher is a generated GoMock package.
package mock_publisher

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	proto "github.com/letsencrypt/boulder/publisher/proto"
)

// MockPublisher is a mock of Publisher interface
type MockPublisher struct {
	ctrl     *gomock.Controller
	recorder *MockPublisherMockRecorder
}

// MockPublisherMockRecorder is the mock recorder for MockPublisher
type MockPublisherMockRecorder struct {
	mock *MockPublisher
}

// NewMockPublisher creates a new mock instance
func NewMockPublisher(ctrl *gomock.Controller) *MockPublisher {
	mock := &MockPublisher{ctrl: ctrl}
	mock.recorder = &MockPublisherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPublisher) EXPECT() *MockPublisherMockRecorder {
	return m.recorder
}

// SubmitToSingleCTWithResult mocks base method
func (m *MockPublisher) SubmitToSingleCTWithResult(arg0 context.Context, arg1 *proto.Request) (*proto.Result, error) {
	ret := m.ctrl.Call(m, "SubmitToSingleCTWithResult", arg0, arg1)
	ret0, _ := ret[0].(*proto.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitToSingleCTWithResult indicates an expected call of SubmitToSingleCTWithResult
func (mr *MockPublisherMockRecorder) SubmitToSingleCTWithResult(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitToSingleCTWithResult", reflect.TypeOf((*MockPublisher)(nil).SubmitToSingleCTWithResult), arg0, arg1)
}
