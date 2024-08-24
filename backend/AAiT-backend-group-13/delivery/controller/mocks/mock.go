// mocks/handlers_mocks.go
package mocks

import (
	"time"

	"github.com/stretchr/testify/mock"

	"github.com/group13/blog/usecase/user/command"
	"github.com/group13/blog/usecase/user/query"
	"github.com/group13/blog/usecase/user/result"
	"github.com/group13/blog/usecase/password_reset"
)

type PromoteHandlerMock struct {
	mock.Mock
}

func (m *PromoteHandlerMock) Handle(cmd *usercmd.PromoteCommand) (bool, error) {
	args := m.Called(cmd)
	return args.Bool(0), args.Error(1)
}

type LoginHandlerMock struct {
	mock.Mock
}

func (m *LoginHandlerMock) Handle(query *userqry.LoginQuery) (*result.LoginInResult, error) {
	args := m.Called(query)
	return args.Get(0).(*result.LoginInResult), args.Error(1)
}

type SignupHandlerMock struct {
	mock.Mock
}

func (m *SignupHandlerMock) Handle(cmd *usercmd.SignUpCommand) (*result.SignUpResult, error) {
	args := m.Called(cmd)
	return args.Get(0).(*result.SignUpResult), args.Error(1)
}

type ResetPasswordHandlerMock struct {
	mock.Mock
}

func (m *ResetPasswordHandlerMock) Handle(cmd *passwordreset.ResetCommand) (bool, error) {
	args := m.Called(cmd)
	return args.Bool(0), args.Error(1)
}

type SendcodeHandlerMock struct {
	mock.Mock
}

func (m *SendcodeHandlerMock) Handle(email string) (time.Time, error) {
	args := m.Called(email)
	return args.Get(0).(time.Time), args.Error(1)
}

type ValidateCodeHandlerMock struct {
	mock.Mock
}

func (m *ValidateCodeHandlerMock) Handle(cmd *passwordreset.ValidateCodeCommand) (string, error) {
	args := m.Called(cmd)
	return args.String(0), args.Error(1)
}

type ValidateEmailHandlerMock struct {
	mock.Mock
}

func (m *ValidateEmailHandlerMock) Handle(email string) (*result.ValidateEmailResult, error) {
	args := m.Called(email)
	return args.Get(0).(*result.ValidateEmailResult), args.Error(1)
}

type UpdateProfileHandlerMock struct {
	mock.Mock
}

func (m *UpdateProfileHandlerMock) Handle(cmd *usercmd.UpdateProfileCommand) (*result.UpdateProfileResult, error) {
	args := m.Called(cmd)
	return args.Get(0).(*result.UpdateProfileResult), args.Error(1)
}
