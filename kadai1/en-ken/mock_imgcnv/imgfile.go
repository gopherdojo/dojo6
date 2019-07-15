// Code generated by MockGen. DO NOT EDIT.
// Source: ./imgcnv/imgfile.go

// Package mock_imgcnv is a generated GoMock package.
package mock_imgcnv

import (
	gomock "github.com/golang/mock/gomock"
	imgcnv "github.com/gopherdojo/dojo6/kadai1/en-ken/imgcnv"
	reflect "reflect"
)

// MockImageFile is a mock of ImageFile interface
type MockImageFile struct {
	ctrl     *gomock.Controller
	recorder *MockImageFileMockRecorder
}

// MockImageFileMockRecorder is the mock recorder for MockImageFile
type MockImageFileMockRecorder struct {
	mock *MockImageFile
}

// NewMockImageFile creates a new mock instance
func NewMockImageFile(ctrl *gomock.Controller) *MockImageFile {
	mock := &MockImageFile{ctrl: ctrl}
	mock.recorder = &MockImageFileMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImageFile) EXPECT() *MockImageFileMockRecorder {
	return m.recorder
}

// AbsPath mocks base method
func (m *MockImageFile) AbsPath() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AbsPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// AbsPath indicates an expected call of AbsPath
func (mr *MockImageFileMockRecorder) AbsPath() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AbsPath", reflect.TypeOf((*MockImageFile)(nil).AbsPath))
}

// SaveAs mocks base method
func (m *MockImageFile) SaveAs(path string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveAs", path)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveAs indicates an expected call of SaveAs
func (mr *MockImageFileMockRecorder) SaveAs(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveAs", reflect.TypeOf((*MockImageFile)(nil).SaveAs), path)
}

// MockImageFileFactory is a mock of ImageFileFactory interface
type MockImageFileFactory struct {
	ctrl     *gomock.Controller
	recorder *MockImageFileFactoryMockRecorder
}

// MockImageFileFactoryMockRecorder is the mock recorder for MockImageFileFactory
type MockImageFileFactoryMockRecorder struct {
	mock *MockImageFileFactory
}

// NewMockImageFileFactory creates a new mock instance
func NewMockImageFileFactory(ctrl *gomock.Controller) *MockImageFileFactory {
	mock := &MockImageFileFactory{ctrl: ctrl}
	mock.recorder = &MockImageFileFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockImageFileFactory) EXPECT() *MockImageFileFactoryMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockImageFileFactory) Create(path string) (imgcnv.ImageFile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", path)
	ret0, _ := ret[0].(imgcnv.ImageFile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockImageFileFactoryMockRecorder) Create(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockImageFileFactory)(nil).Create), path)
}
