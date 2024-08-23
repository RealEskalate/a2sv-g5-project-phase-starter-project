package usecases

import (
	infrastructure "blogs/Infrastructure"
	"blogs/mocks"
	"context"
	"errors"
	"testing"
	"time"

	domain "blogs/Domain"
	// Replace "Path/To" with the actual path to the passwordService package
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type SignupUsecaseTestSuite struct {
	suite.Suite
	SignupUsecaseTestSuite domain.SignupUseCase
	mockSignupRepo         *mocks.SignupRepository
	mockUnverified         *mocks.UnverifiedUserRepository
	contextTimeout         time.Duration
}

func (suite *SignupUsecaseTestSuite) SetupTest() {
	suite.mockSignupRepo = new(mocks.SignupRepository)
	suite.mockUnverified = new(mocks.UnverifiedUserRepository)
	suite.contextTimeout = time.Second * 5
	suite.SignupUsecaseTestSuite = NewSignupUseCase(suite.mockSignupRepo, suite.mockUnverified, suite.contextTimeout, infrastructure.NewPasswordService())
}
func (suite *SignupUsecaseTestSuite) TestCreate() {
	suite.Run("TestSuccess", func() {
		user := domain.User{
			Email:    "user@dfc.com",
			Username: "User",
			Password: "@123asdDdfsd",
		}
		suite.mockSignupRepo.On("FindUserByEmail", mock.Anything, user.Email).Return(domain.User{}, errors.New("no user")).Once()
		suite.mockSignupRepo.On("Create", mock.Anything, mock.AnythingOfType("domain.User")).Return(domain.User{}, nil).Once()
		suite.mockSignupRepo.On("SetOTP", mock.Anything, user.Email, mock.AnythingOfType("string")).Return(nil).Once()
		result := suite.SignupUsecaseTestSuite.Create(context.TODO(), user)
		suite.Equal(result, &domain.SuccessResponse{Message: "Registerd Sucessfully Verify your account", Data: "", Status: 201})
		suite.mockSignupRepo.AssertExpectations(suite.T())
		suite.mockSignupRepo.AssertCalled(suite.T(), "FindUserByEmail", mock.Anything, user.Email)
		suite.mockSignupRepo.AssertCalled(suite.T(), "Create", mock.Anything, mock.Anything)
	})
	suite.Run("required", func() {
		user := domain.User{
			Email:    "",
			Username: "user",
			Password: "password",
		}
		suite.mockSignupRepo.On("FindUserByEmail", user.Email).Return(domain.User{}, nil).Once()
		suite.mockSignupRepo.On("Create", user).Return(domain.User{}, nil).Once()
		result := suite.SignupUsecaseTestSuite.Create(context.TODO(), user)
		suite.Equal(result, &domain.ErrorResponse{Message: "All fields are required", Status: 400})
		// suite.mockSignupRepo.AssertExpectations(suite.T())
		suite.mockSignupRepo.AssertNotCalled(suite.T(), "FindUserByEmail", mock.Anything, user.Email)
		suite.mockSignupRepo.AssertNotCalled(suite.T(), "CreateUser", mock.Anything, mock.Anything)
	})
	suite.Run("invalid email", func() {
		user := domain.User{
			Email:    "1234567",
			Username: "user",
			Password: "password",
		}
		suite.mockSignupRepo.On("FindUserByEmail", user.Email).Return(domain.User{}, nil).Once()
		suite.mockSignupRepo.On("Create", user).Return(domain.User{}, nil).Once()
		result := suite.SignupUsecaseTestSuite.Create(context.TODO(), user)
		suite.Equal(result, &domain.ErrorResponse{Message: "Invalid email format", Status: 400})
		// suite.mockSignupRepo.AssertExpectations(suite.T())
		suite.mockSignupRepo.AssertNotCalled(suite.T(), "FindUserByEmail", mock.Anything, user.Email)
		suite.mockSignupRepo.AssertNotCalled(suite.T(), "CreateUser", mock.Anything, mock.Anything)
	})

}

func (suite *SignupUsecaseTestSuite) TestVerifyOTP() {
	suite.Run("TestSuccess", func() {
		otp := domain.OtpToken{
			Email:     "test",
			ExpiresAt: time.Now().Add(time.Minute * 5),
		}
		suite.mockSignupRepo.On("FindUserByEmail", mock.Anything, otp.Email).Return(domain.User{Email: "test", ExpiresAt: time.Now().Add(time.Minute * 5)}, nil).Once()
		suite.mockSignupRepo.On("VerifyUser", mock.Anything, mock.Anything).Return(domain.User{}, nil).Once()
		result := suite.SignupUsecaseTestSuite.VerifyOTP(context.TODO(), otp)
		suite.Equal(result, &domain.SuccessResponse{Message: "Account verified successfully", Data: domain.User{}, Status: 200})
		suite.mockSignupRepo.AssertExpectations(suite.T())
		suite.mockSignupRepo.AssertCalled(suite.T(), "FindUserByEmail", mock.Anything, otp.Email)
		suite.mockSignupRepo.AssertCalled(suite.T(), "VerifyUser", mock.Anything, mock.Anything)
	})
}
func (suite *SignupUsecaseTestSuite) TestForgotPassword() {
	suite.Run("TestSuccess", func() {
		email := domain.ForgotPasswordRequest{
			Email: "adanemoges6@gmail.com",
		}
		suite.mockSignupRepo.On("FindUserByEmail", mock.Anything, email.Email).Return(domain.User{}, nil).Once()
		suite.mockSignupRepo.On("SetResetToken", mock.Anything, email, mock.AnythingOfType("string"), mock.Anything).Return(domain.User{}, nil).Once()
		// suite.mockSignupRepo.On("SendResetEmail", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

		result := suite.SignupUsecaseTestSuite.ForgotPassword(context.TODO(), email)

		suite.Equal(result, &domain.SuccessResponse{Message: "Reset email sent", Data: "", Status: 200})
		suite.mockSignupRepo.AssertExpectations(suite.T())
		suite.mockSignupRepo.AssertCalled(suite.T(), "FindUserByEmail", mock.Anything, email.Email)
		suite.mockSignupRepo.AssertCalled(suite.T(), "SetResetToken", mock.Anything, email, mock.AnythingOfType("string"), mock.Anything)
		// suite.mockSignupRepo.AssertCalled(suite.T(), "SendResetEmail", mock.Anything, mock.AnythingOfType("string"))
	})
	suite.Run("TestError", func() {
		email := domain.ForgotPasswordRequest{
			Email: "test",
		}
		suite.mockSignupRepo.On("FindUserByEmail", mock.Anything, email.Email).Return(domain.User{}, errors.New("no user")).Once()
		result := suite.SignupUsecaseTestSuite.ForgotPassword(context.TODO(), email)
		suite.Equal(result, &domain.ErrorResponse{Message: "User not found", Status: 404})
		suite.mockSignupRepo.AssertExpectations(suite.T())
		suite.mockSignupRepo.AssertCalled(suite.T(), "FindUserByEmail", mock.Anything, email.Email)
	})
}

func (suite *SignupUsecaseTestSuite) TestResetPassword() {
	suite.Run("TestSuccess", func() {
		password := domain.ResetPasswordRequest{
			Password: "newPasswor@d123",
		}
		token := ""
		user := domain.User{
			Email:                "test@example.com",
			Password:             "oldPassword123@",
			ResetPasswordToken:   token,
			ResetPasswordExpires: time.Now().Add(time.Minute * 5),
		}
		suite.mockSignupRepo.On("FindUserByResetToken", mock.Anything, token).Return(user, nil).Once()
		suite.mockSignupRepo.On("UpdateUser", mock.Anything, mock.Anything).Return(domain.User{}, nil).Once()
		result := suite.SignupUsecaseTestSuite.ResetPassword(context.TODO(), password, token)
		suite.Equal(result, &domain.SuccessResponse{Message: "Password Reset Sucessfully", Status: 200})
		suite.mockSignupRepo.AssertExpectations(suite.T())
		suite.mockSignupRepo.AssertCalled(suite.T(), "FindUserByResetToken", mock.Anything, token)
		suite.mockSignupRepo.AssertCalled(suite.T(), "UpdateUser", mock.Anything, mock.Anything)
	})
	suite.Run("TestInvalidToken", func() {
		password := domain.ResetPasswordRequest{
			Password: "newPassword12@3",
		}
		token := "invalidToken"
		suite.mockSignupRepo.On("FindUserByResetToken", mock.Anything, token).Return(domain.User{}, errors.New("invalid token")).Once()
		result := suite.SignupUsecaseTestSuite.ResetPassword(context.TODO(), password, token)
		suite.Equal(result, &domain.ErrorResponse{Message: "Invalid reset token", Status: 400})
		suite.mockSignupRepo.AssertExpectations(suite.T())
		suite.mockSignupRepo.AssertCalled(suite.T(), "FindUserByResetToken", mock.Anything, token)
	})
	suite.Run("TestExpiredToken", func() {
		password := domain.ResetPasswordRequest{
			Password: "newPassword123@",
		}
		token := "expiredToken"
		user := domain.User{
			Email:                "test@example.com",
			Password:             "oldPasswor@d123",
			ResetPasswordToken:   token,
			ResetPasswordExpires: time.Now().Add(-time.Minute * 5),
		}
		suite.mockSignupRepo.On("FindUserByResetToken", mock.Anything, token).Return(user, nil).Once()
		result := suite.SignupUsecaseTestSuite.ResetPassword(context.TODO(), password, token)
		suite.Equal(result, &domain.ErrorResponse{Message: "Reset token expired", Status: 400})
		suite.mockSignupRepo.AssertExpectations(suite.T())
		suite.mockSignupRepo.AssertCalled(suite.T(), "FindUserByResetToken", mock.Anything, token)
	})
	suite.Run("TestInvalidPassword", func() {
		password := domain.ResetPasswordRequest{
			Password: "weak",
		}
		token := "resetToken"
		user := domain.User{
			Email:                "test@example.com",
			Password:             "oldPassword123",
			ResetPasswordToken:   token,
			ResetPasswordExpires: time.Now().Add(time.Minute * 5),
		}
		suite.mockSignupRepo.On("FindUserByResetToken", mock.Anything, mock.Anything).Return(user, nil).Once()
		result := suite.SignupUsecaseTestSuite.ResetPassword(context.TODO(), password, token)
		suite.Equal(&domain.ErrorResponse{Message: "password must be between 8 and 30 characters long", Status: 400}, result)
		// suite.mockSignupRepo.AssertExpectations(suite.T())
		suite.mockSignupRepo.AssertCalled(suite.T(), "FindUserByResetToken", mock.Anything, mock.Anything)
	})
}

func (suite *SignupUsecaseTestSuite) TestHandleUnverifiedUser() {
	suite.Run("TestSuccess", func() {
		user := domain.User{
			Email:                "test@example.com",
			Username:             "testuser",
			Password:             "password123",
			Verified:             false,
			ExpiresAt:            time.Now().Add(-time.Minute * 5),
			ResetPasswordToken:   "",
			ResetPasswordExpires: time.Time{},
		}
		suite.mockSignupRepo.On("UpdateUser", mock.Anything, mock.Anything).Return(domain.User{}, nil).Once()
		suite.mockSignupRepo.On("SetOTP", mock.Anything, user.Email, mock.AnythingOfType("string")).Return(nil).Once()

		result := suite.SignupUsecaseTestSuite.HandleUnverifiedUser(context.TODO(), user)
		suite.Equal(result, &domain.SuccessResponse{Message: "OTP send to your Email Verify Your Account", Data: "", Status: 201})
		suite.mockSignupRepo.AssertExpectations(suite.T())
		suite.mockSignupRepo.AssertCalled(suite.T(), "UpdateUser", mock.Anything, mock.Anything)
		suite.mockSignupRepo.AssertCalled(suite.T(), "SetOTP", mock.Anything, user.Email, mock.AnythingOfType("string"))

	})
	suite.Run("TestError", func() {
		user := domain.User{
			Email:                "test@example.com",
			Username:             "testuser",
			Password:             "password123",
			Verified:             false,
			ExpiresAt:            time.Now(),
			ResetPasswordToken:   "",
			ResetPasswordExpires: time.Time{},
		}
		suite.mockSignupRepo.On("UpdateUser", mock.Anything, mock.Anything).Return(domain.User{}, errors.New("error updating user")).Once()
		result := suite.SignupUsecaseTestSuite.HandleUnverifiedUser(context.TODO(), user)
		suite.Equal(result, &domain.ErrorResponse{Message: "Error in setting Expiration time", Status: 500})
		suite.mockSignupRepo.AssertExpectations(suite.T())
		suite.mockSignupRepo.AssertCalled(suite.T(), "UpdateUser", mock.Anything, mock.Anything)
	})
}
func TestSignUpUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(SignupUsecaseTestSuite))
}
