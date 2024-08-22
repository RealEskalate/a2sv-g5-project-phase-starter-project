// Code generated by mockery v2.45.0. DO NOT EDIT.

package mocks

import (
	domain "github.com/RealEskalate/blogpost/domain"
	mock "github.com/stretchr/testify/mock"
)

// StateRepository is an autogenerated mock type for the StateRepository type
type StateRepository struct {
	mock.Mock
}

// DeleteState provides a mock function with given fields: _a0
func (_m *StateRepository) DeleteState(_a0 string) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for DeleteState")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetState provides a mock function with given fields: _a0
func (_m *StateRepository) GetState(_a0 string) (*domain.State, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for GetState")
	}

	var r0 *domain.State
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.State, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.State); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.State)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// InsertState provides a mock function with given fields: _a0
func (_m *StateRepository) InsertState(_a0 domain.State) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for InsertState")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.State) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewStateRepository creates a new instance of StateRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStateRepository(t interface {
	mock.TestingT
	Cleanup(func())
}) *StateRepository {
	mock := &StateRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
