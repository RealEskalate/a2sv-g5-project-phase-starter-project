package usecases_test

import (
	bootstrap "aait-backend-group4/Bootstrap"
	domain "aait-backend-group4/Domain"
	usecases "aait-backend-group4/Usecases"
	"aait-backend-group4/mocks"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OTPUsecaseTestSuite struct {
	suite.Suite
	mockOTPRepo          *mocks.OTPRepository
	mockOTPServices      *mocks.OtpInfrastructure
	mockPasswordServices *mocks.PasswordInfrastructure
	mockUserRepo         *mocks.UserRepository
	otpUsecase           domain.OTPUsecase
	env                  bootstrap.Env
}

func (suite *OTPUsecaseTestSuite) SetupTest() {
	suite.mockOTPRepo = new(mocks.OTPRepository)
	suite.mockOTPServices = new(mocks.OtpInfrastructure)
	suite.mockPasswordServices = new(mocks.PasswordInfrastructure)
	suite.mockUserRepo = new(mocks.UserRepository)
	suite.env = bootstrap.Env{EmailApiKey: "test-api-key"}
	suite.otpUsecase = usecases.NewOtpUsecase(
		suite.mockOTPRepo,
		2*time.Second,
		suite.mockOTPServices,
		suite.mockPasswordServices,
		suite.env,
		suite.mockUserRepo,
	)
}

func (suite *OTPUsecaseTestSuite) TestGenerateOTP_Success() {
	ctx := context.Background()
	user := &domain.UserOTPRequest{
		UserID: primitive.NewObjectID().Hex(),
		Email:  "test@example.com",
	}
	otpCode := "123456"
	hashedOTP := "hashed123456"

	suite.mockOTPServices.On("CreateOTP", user).Return(otpCode, nil)
	suite.mockPasswordServices.On("HashPassword", otpCode).Return(hashedOTP, nil)
	suite.mockOTPRepo.On("CreateOTP", mock.Anything, mock.AnythingOfType("*domain.UserOTPVerification")).Return(nil)
	suite.mockOTPServices.On("SendEmail", user.Email, "Email Verification", suite.env.EmailApiKey, otpCode).Return(nil)

	response, err := suite.otpUsecase.GenerateOTP(ctx, user)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Not Verified", response.Status)
	assert.Equal(suite.T(), "OTP Sent to your email please verify your email", response.Message)

	suite.mockOTPServices.AssertExpectations(suite.T())
	suite.mockPasswordServices.AssertExpectations(suite.T())
	suite.mockOTPRepo.AssertExpectations(suite.T())
}

func (suite *OTPUsecaseTestSuite) TestVerifyOTP_Success() {
	ctx := context.Background()
	email := "test@example.com"
	otpCode := "123456"
	hashedOTP := "hashed123456"
	userID := primitive.NewObjectID()

	user := domain.User{
		ID:    userID,
		Email: email,
	}

	otpVerification := domain.UserOTPVerification{
		Email:      email,
		OTP:        hashedOTP,
		Expires_At: time.Now().Add(5 * time.Minute),
	}

	suite.mockUserRepo.On("GetByEmail", ctx, email).Return(user, nil)
	suite.mockOTPRepo.On("GetOTPByEmail", ctx, email).Return(otpVerification, nil)
	suite.mockPasswordServices.On("ComparePasswords", otpCode, hashedOTP).Return(nil)
	suite.mockOTPRepo.On("DeleteOTPByEmail", ctx, email).Return(nil)
	suite.mockUserRepo.On("UpdateUser", ctx, userID.Hex(), mock.AnythingOfType("domain.UserUpdate")).Return(user, nil)

	response, err := suite.otpUsecase.VerifyOTP(ctx, email, otpCode)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Verified", response.Status)
	assert.Equal(suite.T(), "Congrats your account is now verified", response.Message)

	suite.mockUserRepo.AssertExpectations(suite.T())
	suite.mockOTPRepo.AssertExpectations(suite.T())
	suite.mockPasswordServices.AssertExpectations(suite.T())
}

func (suite *OTPUsecaseTestSuite) TestResendOTP_Success() {
	ctx := context.Background()
	email := "test@example.com"
	userID := primitive.NewObjectID()

	user := domain.User{
		ID:    userID,
		Email: email,
	}

	suite.mockUserRepo.On("GetByEmail", ctx, email).Return(user, nil)

	// Set up expectations for GenerateOTP (which is called by ResendOTP)
	suite.mockOTPServices.On("CreateOTP", mock.AnythingOfType("*domain.UserOTPRequest")).Return("123456", nil)
	suite.mockPasswordServices.On("HashPassword", mock.Anything).Return("hashed123456", nil)
	suite.mockOTPRepo.On("CreateOTP", mock.Anything, mock.AnythingOfType("*domain.UserOTPVerification")).Return(nil)
	suite.mockOTPServices.On("SendEmail", email, "Email Verification", suite.env.EmailApiKey, mock.Anything).Return(nil)

	response, err := suite.otpUsecase.ResendOTP(ctx, email)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Not Verified", response.Status)
	assert.Equal(suite.T(), "OTP Sent to your email please verify your email", response.Message)

	suite.mockUserRepo.AssertExpectations(suite.T())
	suite.mockOTPServices.AssertExpectations(suite.T())
	suite.mockPasswordServices.AssertExpectations(suite.T())
	suite.mockOTPRepo.AssertExpectations(suite.T())
}

func TestOTPUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(OTPUsecaseTestSuite))
}
