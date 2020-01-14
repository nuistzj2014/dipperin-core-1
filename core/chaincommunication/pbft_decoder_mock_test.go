// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/dipperin/dipperin-core/core/chaincommunication (interfaces: P2PMsgDecoder)

// Package chaincommunication is a generated GoMock package.
package chaincommunication

import (
	model "github.com/dipperin/dipperin-core/core/model"
	p2p "github.com/dipperin/dipperin-core/third-party/p2p"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockP2PMsgDecoder is a mock of P2PMsgDecoder interface
type MockP2PMsgDecoder struct {
	ctrl     *gomock.Controller
	recorder *MockP2PMsgDecoderMockRecorder
}

// MockP2PMsgDecoderMockRecorder is the mock recorder for MockP2PMsgDecoder
type MockP2PMsgDecoderMockRecorder struct {
	mock *MockP2PMsgDecoder
}

// NewMockP2PMsgDecoder creates a new mock instance
func NewMockP2PMsgDecoder(ctrl *gomock.Controller) *MockP2PMsgDecoder {
	mock := &MockP2PMsgDecoder{ctrl: ctrl}
	mock.recorder = &MockP2PMsgDecoderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockP2PMsgDecoder) EXPECT() *MockP2PMsgDecoderMockRecorder {
	return m.recorder
}

// DecodeTxMsg mocks base method
func (m *MockP2PMsgDecoder) DecodeTxMsg(arg0 p2p.Msg) (model.AbstractTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecodeTxMsg", arg0)
	ret0, _ := ret[0].(model.AbstractTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecodeTxMsg indicates an expected call of DecodeTxMsg
func (mr *MockP2PMsgDecoderMockRecorder) DecodeTxMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeTxMsg", reflect.TypeOf((*MockP2PMsgDecoder)(nil).DecodeTxMsg), arg0)
}

// DecodeTxsMsg mocks base method
func (m *MockP2PMsgDecoder) DecodeTxsMsg(arg0 p2p.Msg) ([]model.AbstractTransaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecodeTxsMsg", arg0)
	ret0, _ := ret[0].([]model.AbstractTransaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecodeTxsMsg indicates an expected call of DecodeTxsMsg
func (mr *MockP2PMsgDecoderMockRecorder) DecodeTxsMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecodeTxsMsg", reflect.TypeOf((*MockP2PMsgDecoder)(nil).DecodeTxsMsg), arg0)
}

// DecoderBlockMsg mocks base method
func (m *MockP2PMsgDecoder) DecoderBlockMsg(arg0 p2p.Msg) (model.AbstractBlock, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DecoderBlockMsg", arg0)
	ret0, _ := ret[0].(model.AbstractBlock)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DecoderBlockMsg indicates an expected call of DecoderBlockMsg
func (mr *MockP2PMsgDecoderMockRecorder) DecoderBlockMsg(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DecoderBlockMsg", reflect.TypeOf((*MockP2PMsgDecoder)(nil).DecoderBlockMsg), arg0)
}