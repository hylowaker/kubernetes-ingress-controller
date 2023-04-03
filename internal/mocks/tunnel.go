// Code generated by MockGen. DO NOT EDIT.
// Source: golang.ngrok.com/ngrok (interfaces: Tunnel)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	ngrok "golang.ngrok.com/ngrok"
	net "net"
	reflect "reflect"
)

// MockTunnel is a mock of Tunnel interface
type MockTunnel struct {
	ctrl     *gomock.Controller
	recorder *MockTunnelMockRecorder
}

// MockTunnelMockRecorder is the mock recorder for MockTunnel
type MockTunnelMockRecorder struct {
	mock *MockTunnel
}

// NewMockTunnel creates a new mock instance
func NewMockTunnel(ctrl *gomock.Controller) *MockTunnel {
	mock := &MockTunnel{ctrl: ctrl}
	mock.recorder = &MockTunnelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTunnel) EXPECT() *MockTunnelMockRecorder {
	return m.recorder
}

// Accept mocks base method
func (m *MockTunnel) Accept() (net.Conn, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Accept")
	ret0, _ := ret[0].(net.Conn)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Accept indicates an expected call of Accept
func (mr *MockTunnelMockRecorder) Accept() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Accept", reflect.TypeOf((*MockTunnel)(nil).Accept))
}

// Addr mocks base method
func (m *MockTunnel) Addr() net.Addr {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Addr")
	ret0, _ := ret[0].(net.Addr)
	return ret0
}

// Addr indicates an expected call of Addr
func (mr *MockTunnelMockRecorder) Addr() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Addr", reflect.TypeOf((*MockTunnel)(nil).Addr))
}

// Close mocks base method
func (m *MockTunnel) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockTunnelMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockTunnel)(nil).Close))
}

// CloseWithContext mocks base method
func (m *MockTunnel) CloseWithContext(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CloseWithContext", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// CloseWithContext indicates an expected call of CloseWithContext
func (mr *MockTunnelMockRecorder) CloseWithContext(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CloseWithContext", reflect.TypeOf((*MockTunnel)(nil).CloseWithContext), arg0)
}

// ForwardsTo mocks base method
func (m *MockTunnel) ForwardsTo() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ForwardsTo")
	ret0, _ := ret[0].(string)
	return ret0
}

// ForwardsTo indicates an expected call of ForwardsTo
func (mr *MockTunnelMockRecorder) ForwardsTo() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ForwardsTo", reflect.TypeOf((*MockTunnel)(nil).ForwardsTo))
}

// ID mocks base method
func (m *MockTunnel) ID() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(string)
	return ret0
}

// ID indicates an expected call of ID
func (mr *MockTunnelMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockTunnel)(nil).ID))
}

// Labels mocks base method
func (m *MockTunnel) Labels() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Labels")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}

// Labels indicates an expected call of Labels
func (mr *MockTunnelMockRecorder) Labels() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Labels", reflect.TypeOf((*MockTunnel)(nil).Labels))
}

// Metadata mocks base method
func (m *MockTunnel) Metadata() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Metadata")
	ret0, _ := ret[0].(string)
	return ret0
}

// Metadata indicates an expected call of Metadata
func (mr *MockTunnelMockRecorder) Metadata() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Metadata", reflect.TypeOf((*MockTunnel)(nil).Metadata))
}

// Proto mocks base method
func (m *MockTunnel) Proto() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Proto")
	ret0, _ := ret[0].(string)
	return ret0
}

// Proto indicates an expected call of Proto
func (mr *MockTunnelMockRecorder) Proto() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Proto", reflect.TypeOf((*MockTunnel)(nil).Proto))
}

// Session mocks base method
func (m *MockTunnel) Session() ngrok.Session {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Session")
	ret0, _ := ret[0].(ngrok.Session)
	return ret0
}

// Session indicates an expected call of Session
func (mr *MockTunnelMockRecorder) Session() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Session", reflect.TypeOf((*MockTunnel)(nil).Session))
}

// URL mocks base method
func (m *MockTunnel) URL() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "URL")
	ret0, _ := ret[0].(string)
	return ret0
}

// URL indicates an expected call of URL
func (mr *MockTunnelMockRecorder) URL() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "URL", reflect.TypeOf((*MockTunnel)(nil).URL))
}
