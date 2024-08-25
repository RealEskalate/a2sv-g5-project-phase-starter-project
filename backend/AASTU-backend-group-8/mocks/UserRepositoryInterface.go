// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "meleket/domain"

	mock "github.com/stretchr/testify/mock"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
)

// UserRepositoryInterface is an autogenerated mock type for the UserRepositoryInterface type
type UserRepositoryInterface struct {
	mock.Mock
}

// Create provides a mock function with given fields: user
func (_m *UserRepositoryInterface) Create(user *domain.User) error {
	ret := _m.Called(user)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: id
func (_m *UserRepositoryInterface) DeleteUser(id primitive.ObjectID) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAllUsers provides a mock function with given fields:
func (_m *UserRepositoryInterface) GetAllUsers() ([]*domain.User, error) {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for GetAllUsers")
	}

	var r0 []*domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func() ([]*domain.User, error)); ok {
		return rf()
	}
	if rf, ok := ret.Get(0).(func() []*domain.User); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByEmail provides a mock function with given fields: email
func (_m *UserRepositoryInterface) GetUserByEmail(email string) (*domain.User, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByEmail")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.User, error)); ok {
		return rf(email)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(email)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByID provides a mock function with given fields: id
func (_m *UserRepositoryInterface) GetUserByID(id primitive.ObjectID) (*domain.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByID")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) (*domain.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(primitive.ObjectID) *domain.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(primitive.ObjectID) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetUserByUsername provides a mock function with given fields: username
func (_m *UserRepositoryInterface) GetUserByUsername(username string) (*domain.User, error) {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for GetUserByUsername")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.User, error)); ok {
		return rf(username)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(username)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(username)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateRole provides a mock function with given fields: username, role
func (_m *UserRepositoryInterface) UpdateRole(username string, role string) error {
	ret := _m.Called(username, role)

	if len(ret) == 0 {
		panic("no return value specified for UpdateRole")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string) error); ok {
		r0 = rf(username, role)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: username, user
func (_m *UserRepositoryInterface) UpdateUser(username string, user *domain.User) error {
	ret := _m.Called(username, user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, *domain.User) error); ok {
		r0 = rf(username, user)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserRepositoryInterface creates a new instance of UserRepositoryInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserRepositoryInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserRepositoryInterface {
	mock := &UserRepositoryInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
