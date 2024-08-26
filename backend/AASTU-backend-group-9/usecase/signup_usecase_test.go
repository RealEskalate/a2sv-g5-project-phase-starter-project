package usecase

import (
    "blog/domain"
    "blog/domain/mocks"
    "context"
    "errors"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
)

func TestRegisterUser(t *testing.T) {
    mockUserRepo := new(mocks.UserRepository)
    mockTokenRepo := new(mocks.TokenRepository)
    mockOTPRepo := new(mocks.OTPRepository)
    mockUser := domain.AuthSignup{
		Username: "testuser",
        Email:    "test@example.com",
        Password: "password123",
    }

    su := NewSignupUsecase(mockUserRepo, mockTokenRepo, mockOTPRepo, time.Second*2)

    t.Run("success", func(t *testing.T) {
        mockUserRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

        userID, err := su.RegisterUser(context.TODO(), &mockUser)

        assert.NoError(t, err)
        assert.NotNil(t, userID)
        mockUserRepo.AssertExpectations(t)
    })

    t.Run("error", func(t *testing.T) {
        mockUserRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("*domain.User")).Return(errors.New("unexpected error")).Once()

        userID, err := su.RegisterUser(context.TODO(), &mockUser)

        assert.Error(t, err)
        assert.Nil(t, userID)
        mockUserRepo.AssertExpectations(t)
    })
}

func TestSendOTP(t *testing.T) {
    mockUserRepo := new(mocks.UserRepository)
    mockTokenRepo := new(mocks.TokenRepository)
    mockOTPRepo := new(mocks.OTPRepository)
    mockUser := domain.AuthSignup{
        Username: "testuser",
        Email:    "test@example.com",
        Password: "password123",
    }

    su := NewSignupUsecase(mockUserRepo, mockTokenRepo, mockOTPRepo, time.Second*2)

    t.Run("success", func(t *testing.T) {
        mockOTPRepo.On("GetOTPByEmail", mock.Anything, mock.AnythingOfType("string")).Return(nil, nil).Once()
        mockOTPRepo.On("SaveOTP", mock.Anything, mock.AnythingOfType("*domain.OTP")).Return(nil).Once()

        err := su.SendOTP(context.TODO(), &mockUser, "smtpusername", "smtppassword","deviceID")

        assert.NoError(t, err)
        mockOTPRepo.AssertExpectations(t)
    })

    t.Run("error", func(t *testing.T) {
        mockOTPRepo.On("GetOTPByEmail", mock.Anything, mock.AnythingOfType("string")).Return(nil, nil).Once()
        mockOTPRepo.On("SaveOTP", mock.Anything, mock.AnythingOfType("*domain.OTP")).Return(errors.New("unexpected error")).Once()

        err := su.SendOTP(context.TODO(), &mockUser, "smtpusername", "smtppassword","deviceID")

        assert.Error(t, err)
        mockOTPRepo.AssertExpectations(t)
    })
}

func TestVerifyOTP(t *testing.T) {
    mockUserRepo := new(mocks.UserRepository)
    mockTokenRepo := new(mocks.TokenRepository)
    mockOTPRepo := new(mocks.OTPRepository)
    mockOTPRequest := domain.OTPRequest{
        Email: "test@example.com",
        Value: "123456",
    }
    mockOTP := domain.OTP{
        Value:     "123456",
        Email:     "test@example.com",
        ExpiresAt: time.Now().Add(time.Minute * 5),
    }

    su := NewSignupUsecase(mockUserRepo, mockTokenRepo, mockOTPRepo, time.Second*2)

    t.Run("success", func(t *testing.T) {
        mockOTPRepo.On("GetOTPByEmail", mock.Anything, mock.AnythingOfType("string")).Return(&mockOTP, nil).Once()
        mockOTPRepo.On("DeleteOTP", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

        otp, err := su.VerifyOTP(context.TODO(), &mockOTPRequest)

        assert.NoError(t, err)
        assert.NotNil(t, otp)
        mockOTPRepo.AssertExpectations(t)
    })

    t.Run("invalid OTP", func(t *testing.T) {
        mockOTPRepo.On("GetOTPByEmail", mock.Anything, mock.AnythingOfType("string")).Return(&mockOTP, nil).Once()

        mockOTPRequest.Value = "654321"
        otp, err := su.VerifyOTP(context.TODO(), &mockOTPRequest)

        assert.Error(t, err)
        assert.Nil(t, otp)
        mockOTPRepo.AssertExpectations(t)
    })

    t.Run("expired OTP", func(t *testing.T) {
        mockOTP.ExpiresAt = time.Now().Add(-time.Minute * 5)
        mockOTPRepo.On("GetOTPByEmail", mock.Anything, mock.AnythingOfType("string")).Return(&mockOTP, nil).Once()

        otp, err := su.VerifyOTP(context.TODO(), &mockOTPRequest)

        assert.Error(t, err)
        assert.Nil(t, otp)
        mockOTPRepo.AssertExpectations(t)
    })
}