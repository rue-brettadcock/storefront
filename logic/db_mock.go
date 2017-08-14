package logic

import (
	gomock "github.com/golang/mock/gomock"
)

type MockMyDb struct {
	ctrl *gomock.Controller
	mydb *MockMyDbRecorder
}

type MockMyDbRecorder struct {
	mock *MockMyDb
}

func NewMockDB(ctrl *gomock.Controller) *MockMyDb {
	mock := &MockMyDb{ctrl: ctrl}
	mock.mydb = &MockMyDbRecorder{mock}
	return mock
}
