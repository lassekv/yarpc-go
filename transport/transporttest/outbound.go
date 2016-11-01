// Copyright (c) 2016 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Automatically generated by MockGen. DO NOT EDIT!
// Source: go.uber.org/yarpc/transport (interfaces: Outbound)

package transporttest

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	transport "go.uber.org/yarpc/transport"
)

// Mock of Outbound interface
type MockOutbound struct {
	ctrl     *gomock.Controller
	recorder *_MockOutboundRecorder
}

// Recorder for MockOutbound (not exported)
type _MockOutboundRecorder struct {
	mock *MockOutbound
}

func NewMockOutbound(ctrl *gomock.Controller) *MockOutbound {
	mock := &MockOutbound{ctrl: ctrl}
	mock.recorder = &_MockOutboundRecorder{mock}
	return mock
}

func (_m *MockOutbound) EXPECT() *_MockOutboundRecorder {
	return _m.recorder
}

func (_m *MockOutbound) Call(_param0 context.Context, _param1 *transport.Request) (*transport.Response, error) {
	ret := _m.ctrl.Call(_m, "Call", _param0, _param1)
	ret0, _ := ret[0].(*transport.Response)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockOutboundRecorder) Call(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Call", arg0, arg1)
}

func (_m *MockOutbound) Start(_param0 transport.Deps) error {
	ret := _m.ctrl.Call(_m, "Start", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockOutboundRecorder) Start(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start", arg0)
}

func (_m *MockOutbound) Stop() error {
	ret := _m.ctrl.Call(_m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockOutboundRecorder) Stop() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop")
}
