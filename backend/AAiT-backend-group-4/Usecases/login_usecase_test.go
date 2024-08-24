package usecases

import (
	bootstrap "aait-backend-group4/Bootstrap"
	domain "aait-backend-group4/Domain"
	"aait-backend-group4/mocks"
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MockEnv struct {
	AccessTokenSecret       string
	RefreshTokenSecret      string
	AccessTokenExpiryMinute int
	RefreshTokenExpiryHour  int
}

type LoginUsecaseTest struct {
	suite.Suite
	mockUserRepo        *mocks.UserRepository
	mockPasswordService *mocks.PasswordInfrastructure
	mockTokenService    *mocks.TokenInfrastructure
	loginUsecase        domain.LoginUsecase
	contextTimeout      time.Duration
}

func (suite *LoginUsecaseTest) SetupTest() {
	suite.mockUserRepo = new(mocks.UserRepository)
	suite.mockPasswordService = new(mocks.PasswordInfrastructure)
	suite.mockTokenService = new(mocks.TokenInfrastructure)
	suite.contextTimeout = time.Second * 2
	mockEnv := &bootstrap.Env{}

	suite.loginUsecase = NewLoginUsecase(suite.mockUserRepo, suite.mockTokenService, suite.contextTimeout, mockEnv)
}

func (suite *LoginUsecaseTest) TearDownTest() {
	suite.mockUserRepo = nil
	suite.mockPasswordService = nil
	suite.mockTokenService = nil
	suite.loginUsecase = nil
}

func (suite *LoginUsecaseTest) TestLoginWithIdentifier_ValidUsername() {
	ctx := context.Background()
	email := "test@example.com"
	username := "testuser"
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Email:    email,
		Verified: true,
	}
	accessToken := "access_token"
	refreshToken := "refresh_token"

	suite.mockUserRepo.On("GetByUsername", ctx, username).Return(user, nil)

	suite.mockTokenService.On("CreateAllTokens", &user, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(accessToken, refreshToken, nil)

	suite.mockUserRepo.On("UpdateUser", ctx, user.ID.Hex(), mock.AnythingOfType("domain.UserUpdate")).Return(user, nil)

	resultAccessToken, resultRefreshToken, err := suite.loginUsecase.LoginWithIdentifier(ctx, username)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), accessToken, resultAccessToken)
	assert.Equal(suite.T(), refreshToken, resultRefreshToken)
}

func (suite *LoginUsecaseTest) TestLoginWithIdentifier_ValidEmail() {
	ctx := context.Background()
	email := "test@example.com"
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Email:    email,
		Verified: true,
	}
	accessToken := "access_token"
	refreshToken := "refresh_token"

	suite.mockUserRepo.On("GetByEmail", ctx, email).Return(user, nil)

	suite.mockTokenService.On("CreateAllTokens", &user, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(accessToken, refreshToken, nil)

	suite.mockUserRepo.On("UpdateUser", ctx, user.ID.Hex(), mock.AnythingOfType("domain.UserUpdate")).Return(user, nil)

	resultAccessToken, resultRefreshToken, err := suite.loginUsecase.LoginWithIdentifier(ctx, email)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), accessToken, resultAccessToken)
	assert.Equal(suite.T(), refreshToken, resultRefreshToken)
}

func (suite *LoginUsecaseTest) TestLoginWithIdentifier_UserNotFound() {
	ctx := context.Background()
	username := "nonexistentuser"

	suite.mockUserRepo.On("GetByUsername", ctx, username).Return(domain.User{}, errors.New("user not found"))

	_, _, err := suite.loginUsecase.LoginWithIdentifier(ctx, username)

	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "user with this username not found")
}

func (suite *LoginUsecaseTest) TestLoginWithIdentifier_UnverifiedUser() {
	ctx := context.Background()
	email := "unverified@example.com"
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Email:    email,
		Verified: false,
	}

	suite.mockUserRepo.On("GetByEmail", ctx, email).Return(user, nil)

	_, _, err := suite.loginUsecase.LoginWithIdentifier(ctx, email)

	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "account not verified")
}

func (suite *LoginUsecaseTest) TestLoginWithIdentifier_TokenCreationError() {
	ctx := context.Background()
	username := "testuser"
	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username: username,
		Verified: true,
	}

	suite.mockUserRepo.On("GetByUsername", ctx, username).Return(user, nil)
	suite.mockTokenService.On("CreateAllTokens", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return("", "", errors.New("token creation error"))

	_, _, err := suite.loginUsecase.LoginWithIdentifier(ctx, username)

	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "token creation error")
}

func TestLoginUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(LoginUsecaseTest))
}
