package port

import (
	reflect "reflect"

	domain "github.com/hender14/app/domain"

	gomock "github.com/golang/mock/gomock"
)

// MockUserRepository is a mock of UserRepository interface.
type MockUserRepository struct {
	ctrl     *gomock.Controller
	recorder *MockUserRepositoryMockRecorder
}

// MockUserRepositoryMockRecorder is the mock recorder for MockUserRepository.
type MockUserRepositoryMockRecorder struct {
	mock *MockUserRepository
}

// NewMockUserRepository creates a new mock instance.
func NewMockUserRepository(ctrl *gomock.Controller) *MockUserRepository {
	mock := &MockUserRepository{ctrl: ctrl}
	mock.recorder = &MockUserRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserRepository) EXPECT() *MockUserRepositoryMockRecorder {
	return m.recorder
}

// QueryEmail mocks base method.
func (m *MockUserRepository) QueryEmail(arg0 *domain.InUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QueryEmail", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// QueryEmail indicates an expected call of QueryEmail.
func (mr *MockUserRepositoryMockRecorder) QueryEmail(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QueryEmail", reflect.TypeOf((*MockUserRepository)(nil).QueryEmail), arg0)
}

// RegisterAccoount mocks base method.
func (m *MockUserRepository) RegisterAccoount(arg0 *domain.SignUser) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RegisterAccoount", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// RegisterAccoount indicates an expected call of RegisterAccoount.
func (mr *MockUserRepositoryMockRecorder) RegisterAccoount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterAccoount", reflect.TypeOf((*MockUserRepository)(nil).RegisterAccoount), arg0)
}

// MockUserInputPort is a mock of UserInputPort interface.
type MockUserInputPort struct {
	ctrl     *gomock.Controller
	recorder *MockUserInputPortMockRecorder
}

// MockUserInputPortMockRecorder is the mock recorder for MockUserInputPort.
type MockUserInputPortMockRecorder struct {
	mock *MockUserInputPort
}

// NewMockUserInputPort creates a new mock instance.
func NewMockUserInputPort(ctrl *gomock.Controller) *MockUserInputPort {
	mock := &MockUserInputPort{ctrl: ctrl}
	mock.recorder = &MockUserInputPortMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUserInputPort) EXPECT() *MockUserInputPortMockRecorder {
	return m.recorder
}

// Sign mocks base method.
func (m *MockUserInputPort) Sign(input *domain.InUser) (interface{}, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Sign", input)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Sign indicates an expected call of Sign.
func (mr *MockUserInputPortMockRecorder) Sign(input interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Sign", reflect.TypeOf((*MockUserInputPort)(nil).Sign), input)
}
