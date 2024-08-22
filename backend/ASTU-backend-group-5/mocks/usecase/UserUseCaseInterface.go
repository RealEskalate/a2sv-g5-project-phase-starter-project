// Code generated by mockery v2.44.2. DO NOT EDIT.

package mocks

import (
	domain "blogApp/internal/domain"

	mock "github.com/stretchr/testify/mock"
)

// UserUseCaseInterface is an autogenerated mock type for the UserUseCaseInterface type
type UserUseCaseInterface struct {
	mock.Mock
}

// AdminRemoveUser provides a mock function with given fields: UserId
func (_m *UserUseCaseInterface) AdminRemoveUser(UserId string) error {
	ret := _m.Called(UserId)

	if len(ret) == 0 {
		panic("no return value specified for AdminRemoveUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(UserId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteUser provides a mock function with given fields: id
func (_m *UserUseCaseInterface) DeleteUser(id string) error {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DemoteFromAdmin provides a mock function with given fields: UserId
func (_m *UserUseCaseInterface) DemoteFromAdmin(UserId string) error {
	ret := _m.Called(UserId)

	if len(ret) == 0 {
		panic("no return value specified for DemoteFromAdmin")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(UserId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FilterUsers provides a mock function with given fields: filter
func (_m *UserUseCaseInterface) FilterUsers(filter map[string]interface{}) ([]*domain.User, error) {
	ret := _m.Called(filter)

	if len(ret) == 0 {
		panic("no return value specified for FilterUsers")
	}

	var r0 []*domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(map[string]interface{}) ([]*domain.User, error)); ok {
		return rf(filter)
	}
	if rf, ok := ret.Get(0).(func(map[string]interface{}) []*domain.User); ok {
		r0 = rf(filter)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(map[string]interface{}) error); ok {
		r1 = rf(filter)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByEmail provides a mock function with given fields: email
func (_m *UserUseCaseInterface) FindUserByEmail(email string) (*domain.User, error) {
	ret := _m.Called(email)

	if len(ret) == 0 {
		panic("no return value specified for FindUserByEmail")
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

// FindUserById provides a mock function with given fields: id
func (_m *UserUseCaseInterface) FindUserById(id string) (*domain.User, error) {
	ret := _m.Called(id)

	if len(ret) == 0 {
		panic("no return value specified for FindUserById")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*domain.User, error)); ok {
		return rf(id)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindUserByUserName provides a mock function with given fields: username
func (_m *UserUseCaseInterface) FindUserByUserName(username string) (*domain.User, error) {
	ret := _m.Called(username)

	if len(ret) == 0 {
		panic("no return value specified for FindUserByUserName")
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

// GetAllUsers provides a mock function with given fields:
func (_m *UserUseCaseInterface) GetAllUsers() ([]*domain.User, error) {
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

// GoogleCallback provides a mock function with given fields: code
func (_m *UserUseCaseInterface) GoogleCallback(code string) (*domain.User, *domain.Token, error) {
	ret := _m.Called(code)

	if len(ret) == 0 {
		panic("no return value specified for GoogleCallback")
	}

	var r0 *domain.User
	var r1 *domain.Token
	var r2 error
	if rf, ok := ret.Get(0).(func(string) (*domain.User, *domain.Token, error)); ok {
		return rf(code)
	}
	if rf, ok := ret.Get(0).(func(string) *domain.User); ok {
		r0 = rf(code)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string) *domain.Token); ok {
		r1 = rf(code)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.Token)
		}
	}

	if rf, ok := ret.Get(2).(func(string) error); ok {
		r2 = rf(code)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// Login provides a mock function with given fields: email, password
func (_m *UserUseCaseInterface) Login(email string, password string) (*domain.User, *domain.Token, error) {
	ret := _m.Called(email, password)

	if len(ret) == 0 {
		panic("no return value specified for Login")
	}

	var r0 *domain.User
	var r1 *domain.Token
	var r2 error
	if rf, ok := ret.Get(0).(func(string, string) (*domain.User, *domain.Token, error)); ok {
		return rf(email, password)
	}
	if rf, ok := ret.Get(0).(func(string, string) *domain.User); ok {
		r0 = rf(email, password)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(string, string) *domain.Token); ok {
		r1 = rf(email, password)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.Token)
		}
	}

	if rf, ok := ret.Get(2).(func(string, string) error); ok {
		r2 = rf(email, password)
	} else {
		r2 = ret.Error(2)
	}

	return r0, r1, r2
}

// PromoteToAdmin provides a mock function with given fields: UserId
func (_m *UserUseCaseInterface) PromoteToAdmin(UserId string) error {
	ret := _m.Called(UserId)

	if len(ret) == 0 {
		panic("no return value specified for PromoteToAdmin")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(UserId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RegisterUser provides a mock function with given fields: _a0
func (_m *UserUseCaseInterface) RegisterUser(_a0 *domain.User) (*domain.User, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for RegisterUser")
	}

	var r0 *domain.User
	var r1 error
	if rf, ok := ret.Get(0).(func(*domain.User) (*domain.User, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(*domain.User) *domain.User); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.User)
		}
	}

	if rf, ok := ret.Get(1).(func(*domain.User) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// RequestEmailVerification provides a mock function with given fields: _a0
func (_m *UserUseCaseInterface) RequestEmailVerification(_a0 domain.User) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for RequestEmailVerification")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(domain.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// RequestPasswordResetUsecase provides a mock function with given fields: userEmail
func (_m *UserUseCaseInterface) RequestPasswordResetUsecase(userEmail string) error {
	ret := _m.Called(userEmail)

	if len(ret) == 0 {
		panic("no return value specified for RequestPasswordResetUsecase")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(userEmail)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// ResetPassword provides a mock function with given fields: token, password, email
func (_m *UserUseCaseInterface) ResetPassword(token string, password string, email string) error {
	ret := _m.Called(token, password, email)

	if len(ret) == 0 {
		panic("no return value specified for ResetPassword")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, string) error); ok {
		r0 = rf(token, password, email)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateUser provides a mock function with given fields: _a0
func (_m *UserUseCaseInterface) UpdateUser(_a0 *domain.User) error {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUser")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(*domain.User) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// VerifyEmail provides a mock function with given fields: token
func (_m *UserUseCaseInterface) VerifyEmail(token string) error {
	ret := _m.Called(token)

	if len(ret) == 0 {
		panic("no return value specified for VerifyEmail")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(token)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewUserUseCaseInterface creates a new instance of UserUseCaseInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUseCaseInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUseCaseInterface {
	mock := &UserUseCaseInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
