// Code generated by MockGen. DO NOT EDIT.
// Source: library/cmd/api/handler/v1/book (interfaces: Handler)

// Package book is a generated GoMock package.
package book

import (
	reflect "reflect"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockHandler is a mock of Handler interface.
type MockHandler struct {
	ctrl     *gomock.Controller
	recorder *MockHandlerMockRecorder
}

// MockHandlerMockRecorder is the mock recorder for MockHandler.
type MockHandlerMockRecorder struct {
	mock *MockHandler
}

// NewMockHandler creates a new mock instance.
func NewMockHandler(ctrl *gomock.Controller) *MockHandler {
	mock := &MockHandler{ctrl: ctrl}
	mock.recorder = &MockHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHandler) EXPECT() *MockHandlerMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockHandler) CreateBook(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "CreateBook", arg0)
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockHandlerMockRecorder) CreateBook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockHandler)(nil).CreateBook), arg0)
}

// DeleteBook mocks base method.
func (m *MockHandler) DeleteBook(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteBook", arg0)
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockHandlerMockRecorder) DeleteBook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockHandler)(nil).DeleteBook), arg0)
}

// GetBookByID mocks base method.
func (m *MockHandler) GetBookByID(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetBookByID", arg0)
}

// GetBookByID indicates an expected call of GetBookByID.
func (mr *MockHandlerMockRecorder) GetBookByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookByID", reflect.TypeOf((*MockHandler)(nil).GetBookByID), arg0)
}

// GetBookByName mocks base method.
func (m *MockHandler) GetBookByName(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "GetBookByName", arg0)
}

// GetBookByName indicates an expected call of GetBookByName.
func (mr *MockHandlerMockRecorder) GetBookByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookByName", reflect.TypeOf((*MockHandler)(nil).GetBookByName), arg0)
}

// ListBook mocks base method.
func (m *MockHandler) ListBook(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ListBook", arg0)
}

// ListBook indicates an expected call of ListBook.
func (mr *MockHandlerMockRecorder) ListBook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBook", reflect.TypeOf((*MockHandler)(nil).ListBook), arg0)
}

// ServeExcel mocks base method.
func (m *MockHandler) ServeExcel(arg0 *gin.Context, arg1, arg2 int, arg3 string, arg4 bool) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ServeExcel", arg0, arg1, arg2, arg3, arg4)
}

// ServeExcel indicates an expected call of ServeExcel.
func (mr *MockHandlerMockRecorder) ServeExcel(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ServeExcel", reflect.TypeOf((*MockHandler)(nil).ServeExcel), arg0, arg1, arg2, arg3, arg4)
}

// UpdatedBook mocks base method.
func (m *MockHandler) UpdatedBook(arg0 *gin.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdatedBook", arg0)
}

// UpdatedBook indicates an expected call of UpdatedBook.
func (mr *MockHandlerMockRecorder) UpdatedBook(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatedBook", reflect.TypeOf((*MockHandler)(nil).UpdatedBook), arg0)
}
