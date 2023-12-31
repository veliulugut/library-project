// Code generated by MockGen. DO NOT EDIT.
// Source: library/service/book (interfaces: ServiceBook)

// Package book is a generated GoMock package.
package book

import (
	context "context"
	ent "library/ent"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockServiceBook is a mock of ServiceBook interface.
type MockServiceBook struct {
	ctrl     *gomock.Controller
	recorder *MockServiceBookMockRecorder
}

// MockServiceBookMockRecorder is the mock recorder for MockServiceBook.
type MockServiceBookMockRecorder struct {
	mock *MockServiceBook
}

// NewMockServiceBook creates a new mock instance.
func NewMockServiceBook(ctrl *gomock.Controller) *MockServiceBook {
	mock := &MockServiceBook{ctrl: ctrl}
	mock.recorder = &MockServiceBookMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceBook) EXPECT() *MockServiceBookMockRecorder {
	return m.recorder
}

// CreateBook mocks base method.
func (m *MockServiceBook) CreateBook(arg0 context.Context, arg1 *CreateBookModel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateBook", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateBook indicates an expected call of CreateBook.
func (mr *MockServiceBookMockRecorder) CreateBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateBook", reflect.TypeOf((*MockServiceBook)(nil).CreateBook), arg0, arg1)
}

// DeleteBook mocks base method.
func (m *MockServiceBook) DeleteBook(arg0 context.Context, arg1 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteBook", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteBook indicates an expected call of DeleteBook.
func (mr *MockServiceBookMockRecorder) DeleteBook(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteBook", reflect.TypeOf((*MockServiceBook)(nil).DeleteBook), arg0, arg1)
}

// FetchBookData mocks base method.
func (m *MockServiceBook) FetchBookData(arg0, arg1 int, arg2 string, arg3 bool) ([]*ent.Book, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchBookData", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*ent.Book)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchBookData indicates an expected call of FetchBookData.
func (mr *MockServiceBookMockRecorder) FetchBookData(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchBookData", reflect.TypeOf((*MockServiceBook)(nil).FetchBookData), arg0, arg1, arg2, arg3)
}

// GetBookByID mocks base method.
func (m *MockServiceBook) GetBookByID(arg0 context.Context, arg1 int) (*GetBookModel, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookByID", arg0, arg1)
	ret0, _ := ret[0].(*GetBookModel)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByID indicates an expected call of GetBookByID.
func (mr *MockServiceBookMockRecorder) GetBookByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookByID", reflect.TypeOf((*MockServiceBook)(nil).GetBookByID), arg0, arg1)
}

// GetBookByName mocks base method.
func (m *MockServiceBook) GetBookByName(arg0 context.Context, arg1 string) (*GetBookByName, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBookByName", arg0, arg1)
	ret0, _ := ret[0].(*GetBookByName)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBookByName indicates an expected call of GetBookByName.
func (mr *MockServiceBookMockRecorder) GetBookByName(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBookByName", reflect.TypeOf((*MockServiceBook)(nil).GetBookByName), arg0, arg1)
}

// ListBook mocks base method.
func (m *MockServiceBook) ListBook(arg0 context.Context, arg1, arg2 int, arg3 string) ([]*GetBookModel, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBook", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]*GetBookModel)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// ListBook indicates an expected call of ListBook.
func (mr *MockServiceBookMockRecorder) ListBook(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBook", reflect.TypeOf((*MockServiceBook)(nil).ListBook), arg0, arg1, arg2, arg3)
}

// UpdatedBook mocks base method.
func (m *MockServiceBook) UpdatedBook(arg0 context.Context, arg1 int, arg2 *UpdateBookModel) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdatedBook", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdatedBook indicates an expected call of UpdatedBook.
func (mr *MockServiceBookMockRecorder) UpdatedBook(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdatedBook", reflect.TypeOf((*MockServiceBook)(nil).UpdatedBook), arg0, arg1, arg2)
}
