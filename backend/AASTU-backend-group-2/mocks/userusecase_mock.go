// Code generated by mockery v2.44.1. DO NOT EDIT.

package mocks

import (
	domain "blog_g2/domain"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// UserUsecase is an autogenerated mock type for the UserUsecase type
type UserUsecase struct {
	mock.Mock
}

// ForgotPassword provides a mock function with given fields: c, email
func (_m *UserUsecase) ForgotPassword(c context.Context, email string) *domain.AppError {
	ret := _m.Called(c, email)

	if len(ret) == 0 {
		panic("no return value specified for ForgotPassword")
	}

	var r0 *domain.AppError
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.AppError); ok {
		r0 = rf(c, email)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.AppError)
		}
	}

	return r0
}

// LoginUser provides a mock function with given fields: c, user
func (_m *UserUsecase) LoginUser(c context.Context, user domain.User) (string, *domain.AppError) {
	ret := _m.Called(c, user)

	if len(ret) == 0 {
		panic("no return value specified for LoginUser")
	}

	var r0 string
	var r1 *domain.AppError
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) (string, *domain.AppError)); ok {
		return rf(c, user)
	}
	if rf, ok := ret.Get(0).(func(context.Context, domain.User) string); ok {
		r0 = rf(c, user)
	} else {
		r0 = ret.Get(0).(string)
	}

	if rf, ok := ret.Get(1).(func(context.Context, domain.User) *domain.AppError); ok {
		r1 = rf(c, user)
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(*domain.AppError)
		}
	}

	return r0, r1
}

// LogoutUser provides a mock function with given fields: c, uid
func (_m *UserUsecase) LogoutUser(c context.Context, uid string) *domain.AppError {
	ret := _m.Called(c, uid)

	if len(ret) == 0 {
		panic("no return value specified for LogoutUser")
	}

	var r0 *domain.AppError
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.AppError); ok {
		r0 = rf(c, uid)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.AppError)
		}
	}

	return r0
}

// PromoteDemoteUser provides a mock function with given fields: c, userid, isAdmin
func (_m *UserUsecase) PromoteDemoteUser(c context.Context, userid string, isAdmin bool) *domain.AppError {
	ret := _m.Called(c, userid, isAdmin)

	if len(ret) == 0 {
		panic("no return value specified for PromoteDemoteUser")
	}

	var r0 *domain.AppError
	if rf, ok := ret.Get(0).(func(context.Context, string, bool) *domain.AppError); ok {
		r0 = rf(c, userid, isAdmin)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.AppError)
		}
	}

	return r0
}

// RegisterUser provides a mock function with given fields: c, user
func (_m *UserUsecase) RegisterUser(c context.Context, user *domain.User) *domain.AppError {
	ret := _m.Called(c, user)

	if len(ret) == 0 {
		panic("no return value specified for RegisterUser")
	}

	var r0 *domain.AppError
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) *domain.AppError); ok {
		r0 = rf(c, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.AppError)
		}
	}

	return r0
}

// ResetPassword provides a mock function with given fields: c, token, newPassword
func (_m *UserUsecase) ResetPassword(c context.Context, token string, newPassword string) *domain.AppError {
	ret := _m.Called(c, token, newPassword)

	if len(ret) == 0 {
		panic("no return value specified for ResetPassword")
	}

	var r0 *domain.AppError
	if rf, ok := ret.Get(0).(func(context.Context, string, string) *domain.AppError); ok {
		r0 = rf(c, token, newPassword)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.AppError)
		}
	}

	return r0
}

// UpdateUserDetails provides a mock function with given fields: c, user
func (_m *UserUsecase) UpdateUserDetails(c context.Context, user *domain.User) *domain.AppError {
	ret := _m.Called(c, user)

	if len(ret) == 0 {
		panic("no return value specified for UpdateUserDetails")
	}

	var r0 *domain.AppError
	if rf, ok := ret.Get(0).(func(context.Context, *domain.User) *domain.AppError); ok {
		r0 = rf(c, user)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.AppError)
		}
	}

	return r0
}

// VerifyUserEmail provides a mock function with given fields: c, token
func (_m *UserUsecase) VerifyUserEmail(c context.Context, token string) *domain.AppError {
	ret := _m.Called(c, token)

	if len(ret) == 0 {
		panic("no return value specified for VerifyUserEmail")
	}

	var r0 *domain.AppError
	if rf, ok := ret.Get(0).(func(context.Context, string) *domain.AppError); ok {
		r0 = rf(c, token)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.AppError)
		}
	}

	return r0
}

// NewUserUsecase creates a new instance of UserUsecase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewUserUsecase(t interface {
	mock.TestingT
	Cleanup(func())
}) *UserUsecase {
	mock := &UserUsecase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
