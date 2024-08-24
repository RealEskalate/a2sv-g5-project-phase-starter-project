package usecases_test

import (
	domain "aait-backend-group4/Domain"
	usecases "aait-backend-group4/Usecases"
	"aait-backend-group4/mocks"
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type LogoutUsecaseTestSuite struct {
	suite.Suite
	mockTokenService *mocks.TokenInfrastructure
	logoutUsecase    domain.LogoutUsecase
}

func (suite *LogoutUsecaseTestSuite) SetupTest() {
	suite.mockTokenService = new(mocks.TokenInfrastructure)
	suite.logoutUsecase = usecases.NewLogoutUsecase(suite.mockTokenService)
}

func (suite *LogoutUsecaseTestSuite) TestLogout_Success() {
	ctx := context.Background()
	token := "valid_token"
	userID := "user123"

	suite.mockTokenService.On("ExtractUserIDFromToken", token).Return(userID, nil)
	suite.mockTokenService.On("RemoveTokens", userID).Return(nil)

	response, err := suite.logoutUsecase.Logout(ctx, token)

	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), "Logout successful", response.Message)

	suite.mockTokenService.AssertExpectations(suite.T())
}

func (suite *LogoutUsecaseTestSuite) TestLogout_InvalidToken() {
	ctx := context.Background()
	token := "invalid_token"

	suite.mockTokenService.On("ExtractUserIDFromToken", token).Return("", errors.New("invalid token"))

	response, err := suite.logoutUsecase.Logout(ctx, token)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.LogoutResponse{}, response)

	suite.mockTokenService.AssertExpectations(suite.T())
}

func (suite *LogoutUsecaseTestSuite) TestLogout_RemoveTokensError() {
	ctx := context.Background()
	token := "valid_token"
	userID := "user123"

	suite.mockTokenService.On("ExtractUserIDFromToken", token).Return(userID, nil)
	suite.mockTokenService.On("RemoveTokens", userID).Return(errors.New("failed to remove tokens"))

	response, err := suite.logoutUsecase.Logout(ctx, token)

	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), domain.LogoutResponse{}, response)

	suite.mockTokenService.AssertExpectations(suite.T())
}

func TestLogoutUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(LogoutUsecaseTestSuite))
}
