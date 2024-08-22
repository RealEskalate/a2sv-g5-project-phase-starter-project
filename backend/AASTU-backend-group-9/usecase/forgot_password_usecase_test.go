package usecase

import (
    "context"
    "errors"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"

    "blog/domain"
    "blog/domain/mocks"
)

type ForgotPasswordUsecaseSuite struct {
    suite.Suite
    userRepoMock   *mocks.UserRepository
    otpRepoMock    *mocks.OTPRepository
    forgotPassword *forgotPasswordUsecase
}

func (suite *ForgotPasswordUsecaseSuite) SetupTest() {
    suite.userRepoMock = new(mocks.UserRepository)
    suite.otpRepoMock = new(mocks.OTPRepository)
    suite.forgotPassword = &forgotPasswordUsecase{
        userRepository: suite.userRepoMock,
        otpRepository:  suite.otpRepoMock,
        contextTimeout: time.Second * 2,
    }
}

func (suite *ForgotPasswordUsecaseSuite) TestSendResetOTP_Success() {
    ctx, cancel := context.WithTimeout(context.Background(), suite.forgotPassword.contextTimeout)
    defer cancel()

    email := "test@example.com"
    smtpUsername := "smtpUser"
    smtpPassword := "smtpPass"

    suite.userRepoMock.On("GetUserByEmail", ctx, email).Return(&domain.User{Email: email}, nil)

    err := suite.forgotPassword.SendResetOTP(ctx, email, smtpUsername, smtpPassword)

    assert.NoError(suite.T(), err)
    suite.userRepoMock.AssertExpectations(suite.T())
}

func (suite *ForgotPasswordUsecaseSuite) TestSendResetOTP_Failure() {
    ctx, cancel := context.WithTimeout(context.Background(), suite.forgotPassword.contextTimeout)
    defer cancel()

    email := "test@example.com"
    smtpUsername := "smtpUser"
    smtpPassword := "smtpPass"

    suite.userRepoMock.On("GetUserByEmail", ctx, email).Return(nil, errors.New("user not found"))

    err := suite.forgotPassword.SendResetOTP(ctx, email, smtpUsername, smtpPassword)

    assert.Error(suite.T(), err)
    suite.userRepoMock.AssertExpectations(suite.T())
}

func (suite *ForgotPasswordUsecaseSuite) TestResetPassword_Success() {
    ctx, cancel := context.WithTimeout(context.Background(), suite.forgotPassword.contextTimeout)
    defer cancel()

    email := "test@example.com"
    otpValue := "123456"
    newPassword := "newPassword"

    suite.otpRepoMock.On("GetOTPByEmail", ctx, email).Return(&domain.OTP{Value: otpValue}, nil)
    suite.otpRepoMock.On("ValidateOTP", ctx, email, otpValue).Return(true, nil)
    suite.userRepoMock.On("UpdatePassword", ctx, email, newPassword).Return(nil)

    err := suite.forgotPassword.ResetPassword(ctx, email, otpValue, newPassword)

    assert.NoError(suite.T(), err)
    suite.otpRepoMock.AssertExpectations(suite.T())
    suite.userRepoMock.AssertExpectations(suite.T())
}

func (suite *ForgotPasswordUsecaseSuite) TestResetPassword_Failure() {
    ctx, cancel := context.WithTimeout(context.Background(), suite.forgotPassword.contextTimeout)
    defer cancel()

    email := "test@example.com"
    otpValue := "123456"
    newPassword := "newPassword"

    suite.otpRepoMock.On("GetOTPByEmail", ctx, email).Return(nil, errors.New("OTP not found"))
    suite.otpRepoMock.On("ValidateOTP", ctx, email, otpValue).Return(false, errors.New("invalid OTP"))

    err := suite.forgotPassword.ResetPassword(ctx, email, otpValue, newPassword)

    assert.Error(suite.T(), err)
    suite.otpRepoMock.AssertExpectations(suite.T())
}

func TestForgotPasswordUsecaseSuite(t *testing.T) {
    suite.Run(t, new(ForgotPasswordUsecaseSuite))
}