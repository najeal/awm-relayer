// Code generated by MockGen. DO NOT EDIT.
// Source: apprequest_network_signer_test.go
//
// Generated by this command:
//
//	mockgen -source=apprequest_network_signer_test.go -destination=mocks/mock_inboundmessage.go -package=mocks -mock_names=InboundMessage=MockInboundMessage
//

// Package mocks is a generated GoMock package.
package mocks

import (
	fmt "fmt"
	reflect "reflect"
	time "time"

	ids "github.com/ava-labs/avalanchego/ids"
	message "github.com/ava-labs/avalanchego/message"
	gomock "go.uber.org/mock/gomock"
)

// MockInboundMessage is a mock of InboundMessage interface.
type MockInboundMessage struct {
	ctrl     *gomock.Controller
	recorder *MockInboundMessageMockRecorder
}

// MockInboundMessageMockRecorder is the mock recorder for MockInboundMessage.
type MockInboundMessageMockRecorder struct {
	mock *MockInboundMessage
}

// NewMockInboundMessage creates a new mock instance.
func NewMockInboundMessage(ctrl *gomock.Controller) *MockInboundMessage {
	mock := &MockInboundMessage{ctrl: ctrl}
	mock.recorder = &MockInboundMessageMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInboundMessage) EXPECT() *MockInboundMessageMockRecorder {
	return m.recorder
}

// BytesSavedCompression mocks base method.
func (m *MockInboundMessage) BytesSavedCompression() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BytesSavedCompression")
	ret0, _ := ret[0].(int)
	return ret0
}

// BytesSavedCompression indicates an expected call of BytesSavedCompression.
func (mr *MockInboundMessageMockRecorder) BytesSavedCompression() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BytesSavedCompression", reflect.TypeOf((*MockInboundMessage)(nil).BytesSavedCompression))
}

// Expiration mocks base method.
func (m *MockInboundMessage) Expiration() time.Time {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Expiration")
	ret0, _ := ret[0].(time.Time)
	return ret0
}

// Expiration indicates an expected call of Expiration.
func (mr *MockInboundMessageMockRecorder) Expiration() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Expiration", reflect.TypeOf((*MockInboundMessage)(nil).Expiration))
}

// Message mocks base method.
func (m *MockInboundMessage) Message() fmt.Stringer {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Message")
	ret0, _ := ret[0].(fmt.Stringer)
	return ret0
}

// Message indicates an expected call of Message.
func (mr *MockInboundMessageMockRecorder) Message() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Message", reflect.TypeOf((*MockInboundMessage)(nil).Message))
}

// NodeID mocks base method.
func (m *MockInboundMessage) NodeID() ids.NodeID {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NodeID")
	ret0, _ := ret[0].(ids.NodeID)
	return ret0
}

// NodeID indicates an expected call of NodeID.
func (mr *MockInboundMessageMockRecorder) NodeID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NodeID", reflect.TypeOf((*MockInboundMessage)(nil).NodeID))
}

// OnFinishedHandling mocks base method.
func (m *MockInboundMessage) OnFinishedHandling() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OnFinishedHandling")
}

// OnFinishedHandling indicates an expected call of OnFinishedHandling.
func (mr *MockInboundMessageMockRecorder) OnFinishedHandling() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OnFinishedHandling", reflect.TypeOf((*MockInboundMessage)(nil).OnFinishedHandling))
}

// Op mocks base method.
func (m *MockInboundMessage) Op() message.Op {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Op")
	ret0, _ := ret[0].(message.Op)
	return ret0
}

// Op indicates an expected call of Op.
func (mr *MockInboundMessageMockRecorder) Op() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Op", reflect.TypeOf((*MockInboundMessage)(nil).Op))
}

// String mocks base method.
func (m *MockInboundMessage) String() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "String")
	ret0, _ := ret[0].(string)
	return ret0
}

// String indicates an expected call of String.
func (mr *MockInboundMessageMockRecorder) String() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "String", reflect.TypeOf((*MockInboundMessage)(nil).String))
}