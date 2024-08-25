package usecase

import (
	"Blog_Starter/domain"
	"Blog_Starter/domain/mocks"
	"context"
	"errors"
	"testing"
	"time"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type SignupUseCaseTestSuite struct {
	suite.Suite
	signupUseCase domain.SignupUsecase
	userRepo      *mocks.UserRepository
}

func (suite *SignupUseCaseTestSuite) SetupTest() {
	suite.userRepo = new(mocks.UserRepository)
	suite.signupUseCase = NewSignUpUsecase(suite.userRepo, 2*time.Second)
}

func (suite *SignupUseCaseTestSuite) TestCreateUser_Success() {
	userSignUp := &domain.UserSignUp{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(userSignUp.Password), bcrypt.DefaultCost)

	userCreate := &domain.User{
		UserID:      primitive.NewObjectID(),
		Username:    userSignUp.Username,
		Password:    string(hashedPassword),
		Email:       userSignUp.Email,
		CreatedAt:   time.Now(),
		IsActivated: false,
	}

	suite.userRepo.On("GetUserByEmail", mock.Anything, userSignUp.Email).Return(nil, errors.New("not found"))
	suite.userRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("*domain.User")).Return(userCreate, nil)

	user, err := suite.signupUseCase.CreateUser(context.Background(), userSignUp)
	suite.NoError(err)
	suite.NotNil(user)
	suite.Equal(userCreate.Email, user.Email)
	suite.Equal(userCreate.Username, user.Username)
	suite.userRepo.AssertExpectations(suite.T())
}

func (suite *SignupUseCaseTestSuite) TestCreateUser_AlreadyExists() {
	userSignUp := &domain.UserSignUp{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	existingUser := &domain.User{
		UserID:      primitive.NewObjectID(),
		Username:    userSignUp.Username,
		Password:    "hashedpassword",
		Email:       userSignUp.Email,
		IsActivated: true,
	}

	suite.userRepo.On("GetUserByEmail", mock.Anything, userSignUp.Email).Return(existingUser, nil)

	user, err := suite.signupUseCase.CreateUser(context.Background(), userSignUp)
	suite.Error(err)
	suite.Nil(user)
	suite.EqualError(err, "user already exists")
	suite.userRepo.AssertExpectations(suite.T())
}

func (suite *SignupUseCaseTestSuite) TestVerifyEmail_Success() {
	verifyRequest := &domain.VerifyEmailRequest{
		Email: "test@example.com",
		OTP:   "123456",
	}

	user := &domain.User{
		UserID:      primitive.NewObjectID(),
		Username:    "testuser",
		Email:       verifyRequest.Email,
		IsActivated: false,
	}

	suite.userRepo.On("GetUserByEmail", mock.Anything, verifyRequest.Email).Return(user, nil)
	suite.userRepo.On("UpdateSignup", mock.Anything, user).Return(nil)

	resp, err := suite.signupUseCase.VerifyEmail(context.Background(), verifyRequest)
	suite.NoError(err)
	suite.NotNil(resp)
	suite.True(resp.IsActivated)
	suite.userRepo.AssertExpectations(suite.T())
}

func (suite *SignupUseCaseTestSuite) TestVerifyEmail_AlreadyActivated() {
	verifyRequest := &domain.VerifyEmailRequest{
		Email: "test@example.com",
		OTP:   "123456",
	}

	user := &domain.User{
		UserID:      primitive.NewObjectID(),
		Username:    "testuser",
		Email:       verifyRequest.Email,
		IsActivated: true,
	}

	suite.userRepo.On("GetUserByEmail", mock.Anything, verifyRequest.Email).Return(user, nil)

	resp, err := suite.signupUseCase.VerifyEmail(context.Background(), verifyRequest)
	suite.Error(err)
	suite.Nil(resp)
	suite.EqualError(err, "user already activated")
	suite.userRepo.AssertExpectations(suite.T())
}

func (suite *SignupUseCaseTestSuite) TestResendOTP_Success() {
	resendRequest := &domain.ResendOTPRequest{
		Email: "test@example.com",
	}

	user := &domain.User{
		UserID:      primitive.NewObjectID(),
		Username:    "testuser",
		Email:       resendRequest.Email,
		IsActivated: false,
	}

	suite.userRepo.On("GetUserByEmail", mock.Anything, resendRequest.Email).Return(user, nil)

	err := suite.signupUseCase.ResendOTP(context.Background(), resendRequest)
	suite.NoError(err)
	suite.userRepo.AssertExpectations(suite.T())
}

func (suite *SignupUseCaseTestSuite) TestResendOTP_AlreadyActivated() {
	resendRequest := &domain.ResendOTPRequest{
		Email: "test@example.com",
	}

	user := &domain.User{
		UserID:      primitive.NewObjectID(),
		Username:    "testuser",
		Email:       resendRequest.Email,
		IsActivated: true,
	}

	suite.userRepo.On("GetUserByEmail", mock.Anything, resendRequest.Email).Return(user, nil)

	err := suite.signupUseCase.ResendOTP(context.Background(), resendRequest)
	suite.Error(err)
	suite.EqualError(err, "failed to resend otp. User account already activated")
	suite.userRepo.AssertExpectations(suite.T())
}

func TestSignupUseCaseTestSuite(t *testing.T) {
	suite.Run(t, new(SignupUseCaseTestSuite))
}
