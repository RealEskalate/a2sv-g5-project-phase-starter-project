package usecases

import (
	domain "aait-backend-group4/Domain"
	"aait-backend-group4/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type SignupUsecaseTest struct {
	suite.Suite
	mockUserRepo        *mocks.UserRepository
	mockOTPUsecase      *mocks.OTPUsecase
	mockPasswordService *mocks.PasswordInfrastructure
	mockTokenService    *mocks.TokenInfrastructure
	signupUsecase       domain.SignupUsecase
	contextTimeout      time.Duration
}

func (suite *SignupUsecaseTest) SetupSuite() {
	suite.mockUserRepo = new(mocks.UserRepository)
	suite.mockOTPUsecase = new(mocks.OTPUsecase)
	suite.mockPasswordService = new(mocks.PasswordInfrastructure)
	suite.mockTokenService = new(mocks.TokenInfrastructure)
	suite.contextTimeout = time.Second * 2
	suite.signupUsecase = NewSingupUsecase(suite.mockUserRepo, suite.mockOTPUsecase,
		suite.contextTimeout, suite.mockPasswordService, suite.mockTokenService)
}

func (suite *SignupUsecaseTest) TearDownSuite() {
	suite.mockUserRepo = nil
	suite.mockOTPUsecase = nil
	suite.mockPasswordService = nil
	suite.mockTokenService = nil
	suite.signupUsecase = nil
}

func (suite *SignupUsecaseTest) TestSignup_EmailAlreadyExists() {
	ctx, cancel := context.WithTimeout(context.Background(), suite.contextTimeout)
	defer cancel()
	signupRequest := &domain.SignupRequest{
		Email:     "test@example.com",
		User_Name: "testuser",
		Password:  "password123",
	}

	// Update the mock expectation to use mock.Anything for the context argument
	suite.mockUserRepo.On("GetByEmail", mock.Anything, signupRequest.Email).Return(domain.User{}, nil)

	_, err := suite.signupUsecase.Signup(ctx, signupRequest)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "email already exists", err.Error())
}

func (suite *SignupUsecaseTest) TestSignup_UsernameAlreadyTaken() {
	ctx, cancel := context.WithTimeout(context.Background(), suite.contextTimeout)
	defer cancel()
	signupRequest := &domain.SignupRequest{
		Email:     "newuser@example.com",
		User_Name: "existinguser",
		Password:  "password123",
	}

	// Return empty User and error for GetByEmail to simulate email not found
	suite.mockUserRepo.On("GetByEmail", mock.Anything, signupRequest.Email).Return(domain.User{}, errors.New("not found"))

	// Return a non-empty User and nil error for GetByUsername to simulate username already taken
	existingUser := domain.User{Username: "existinguser"}
	suite.mockUserRepo.On("GetByUsername", mock.Anything, signupRequest.User_Name).Return(existingUser, nil)

	_, err := suite.signupUsecase.Signup(ctx, signupRequest)

	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "username alread taken", err.Error())
}

func (suite *SignupUsecaseTest) TestSignup_Success() {
	ctx, cancel := context.WithTimeout(context.Background(), suite.contextTimeout)
	defer cancel()
	signupRequest := &domain.SignupRequest{
		First_Name: "John",
		Last_Name:  "Doe",
		Email:      "johndoe@example.com",
		User_Name:  "johndoe",
		Password:   "password123",
		Image_Path: "path/to/image.jpg",
	}

	// Simulate email not found
	suite.mockUserRepo.On("GetByEmail", mock.Anything, signupRequest.Email).Return(domain.User{}, errors.New("not found"))

	// Simulate username not found
	suite.mockUserRepo.On("GetByUsername", mock.Anything, signupRequest.User_Name).Return(domain.User{}, errors.New("not found"))

	// Simulate password hashing
	hashedPassword := "hashed_password_123"
	suite.mockPasswordService.On("HashPassword", signupRequest.Password).Return(hashedPassword, nil)

	// Simulate user creation
	suite.mockUserRepo.On("CreateUser", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil)

	// Simulate OTP generation
	expectedOTPResponse := domain.OTPVerificationResponse{
		Status:  "success",
		Message: "OTP sent successfully",
	}
	suite.mockOTPUsecase.On("GenerateOTP", mock.Anything, mock.AnythingOfType("*domain.UserOTPRequest")).Return(expectedOTPResponse, nil)

	// Call the Signup method
	resp, err := suite.signupUsecase.Signup(ctx, signupRequest)

	// Assertions
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), resp)
	assert.Equal(suite.T(), expectedOTPResponse.Status, resp.Status)
	assert.Equal(suite.T(), expectedOTPResponse.Message, resp.Message)

	// Verify that all expected mock calls were made
	suite.mockUserRepo.AssertExpectations(suite.T())
	suite.mockPasswordService.AssertExpectations(suite.T())
	suite.mockOTPUsecase.AssertExpectations(suite.T())

	// Check the created user
	createUserCall := suite.mockUserRepo.Calls[2]
	createdUser, ok := createUserCall.Arguments.Get(1).(*domain.User)

	if ok {
		assert.Equal(suite.T(), signupRequest.First_Name, createdUser.First_Name)
		assert.Equal(suite.T(), signupRequest.Last_Name, createdUser.Last_Name)
		assert.Equal(suite.T(), signupRequest.Email, createdUser.Email)
		assert.Equal(suite.T(), signupRequest.User_Name, createdUser.Username)
		assert.Equal(suite.T(), hashedPassword, createdUser.Password)
		assert.Equal(suite.T(), &signupRequest.Image_Path, createdUser.ProfileImage)
		assert.Equal(suite.T(), "USER", createdUser.User_Role)
		assert.False(suite.T(), createdUser.Verified)
	}
}
func TestSignupUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(SignupUsecaseTest))
}
