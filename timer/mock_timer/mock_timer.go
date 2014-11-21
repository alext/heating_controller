// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/alext/heating-controller/timer (interfaces: Timer)

package mock_timer

import (
	timer "github.com/alext/heating-controller/timer"
	gomock "code.google.com/p/gomock/gomock"
)

// Mock of Timer interface
type MockTimer struct {
	ctrl     *gomock.Controller
	recorder *_MockTimerRecorder
}

// Recorder for MockTimer (not exported)
type _MockTimerRecorder struct {
	mock *MockTimer
}

func NewMockTimer(ctrl *gomock.Controller) *MockTimer {
	mock := &MockTimer{ctrl: ctrl}
	mock.recorder = &_MockTimerRecorder{mock}
	return mock
}

func (_m *MockTimer) EXPECT() *_MockTimerRecorder {
	return _m.recorder
}

func (_m *MockTimer) AddEvent(_param0 timer.Event) {
	_m.ctrl.Call(_m, "AddEvent", _param0)
}

func (_mr *_MockTimerRecorder) AddEvent(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "AddEvent", arg0)
}

func (_m *MockTimer) NextEvent() *timer.Event {
	ret := _m.ctrl.Call(_m, "NextEvent")
	ret0, _ := ret[0].(*timer.Event)
	return ret0
}

func (_mr *_MockTimerRecorder) NextEvent() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "NextEvent")
}

func (_m *MockTimer) Running() bool {
	ret := _m.ctrl.Call(_m, "Running")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockTimerRecorder) Running() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Running")
}

func (_m *MockTimer) Start() {
	_m.ctrl.Call(_m, "Start")
}

func (_mr *_MockTimerRecorder) Start() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Start")
}

func (_m *MockTimer) Stop() {
	_m.ctrl.Call(_m, "Stop")
}

func (_mr *_MockTimerRecorder) Stop() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Stop")
}
