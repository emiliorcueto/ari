// Automatically generated by MockGen. DO NOT EDIT!
// Source: github.com/CyCoreSystems/ari (interfaces: StoredRecording)

package mock

import (
	ari "github.com/CyCoreSystems/ari"
	gomock "github.com/golang/mock/gomock"
)

// Mock of StoredRecording interface
type MockStoredRecording struct {
	ctrl     *gomock.Controller
	recorder *_MockStoredRecordingRecorder
}

// Recorder for MockStoredRecording (not exported)
type _MockStoredRecordingRecorder struct {
	mock *MockStoredRecording
}

func NewMockStoredRecording(ctrl *gomock.Controller) *MockStoredRecording {
	mock := &MockStoredRecording{ctrl: ctrl}
	mock.recorder = &_MockStoredRecordingRecorder{mock}
	return mock
}

func (_m *MockStoredRecording) EXPECT() *_MockStoredRecordingRecorder {
	return _m.recorder
}

func (_m *MockStoredRecording) Copy(_param0 string, _param1 string) (*ari.StoredRecordingHandle, error) {
	ret := _m.ctrl.Call(_m, "Copy", _param0, _param1)
	ret0, _ := ret[0].(*ari.StoredRecordingHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStoredRecordingRecorder) Copy(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Copy", arg0, arg1)
}

func (_m *MockStoredRecording) Data(_param0 string) (ari.StoredRecordingData, error) {
	ret := _m.ctrl.Call(_m, "Data", _param0)
	ret0, _ := ret[0].(ari.StoredRecordingData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStoredRecordingRecorder) Data(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Data", arg0)
}

func (_m *MockStoredRecording) Delete(_param0 string) error {
	ret := _m.ctrl.Call(_m, "Delete", _param0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStoredRecordingRecorder) Delete(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0)
}

func (_m *MockStoredRecording) Get(_param0 string) *ari.StoredRecordingHandle {
	ret := _m.ctrl.Call(_m, "Get", _param0)
	ret0, _ := ret[0].(*ari.StoredRecordingHandle)
	return ret0
}

func (_mr *_MockStoredRecordingRecorder) Get(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0)
}

func (_m *MockStoredRecording) List() ([]*ari.StoredRecordingHandle, error) {
	ret := _m.ctrl.Call(_m, "List")
	ret0, _ := ret[0].([]*ari.StoredRecordingHandle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStoredRecordingRecorder) List() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "List")
}
