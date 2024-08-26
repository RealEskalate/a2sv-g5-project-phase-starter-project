package usecase

import (
    "context"
    "errors"
    "testing"
    "time"

    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/suite"
    "go.mongodb.org/mongo-driver/bson/primitive"

    "blog/domain"
    "blog/domain/mocks"
)

type LogoutUsecaseSuite struct {
    suite.Suite
    tokenRepoMock  *mocks.TokenRepository
    logoutUsecase  *LogoutUsecase
}

func (suite *LogoutUsecaseSuite) SetupTest() {
    suite.tokenRepoMock = new(mocks.TokenRepository)
    suite.logoutUsecase = &LogoutUsecase{
        tokenRepository: suite.tokenRepoMock,
        contextTimeout:  time.Second * 2,
    }
}

func (suite *LogoutUsecaseSuite) TestLogout_Success() {
    ctx, cancel := context.WithTimeout(context.Background(), suite.logoutUsecase.contextTimeout)
    defer cancel()
    refreshToken := "valid_refresh_token"
    tokenID, _ := primitive.ObjectIDFromHex("507f1f77bcf86cd799439011")
    token := &domain.Token{ID: tokenID, RefreshToken: refreshToken}

    suite.tokenRepoMock.On("FindTokenByRefreshToken", ctx, refreshToken).Return(token, nil)
    suite.tokenRepoMock.On("DeleteToken", ctx, token.ID).Return(nil)

    err := suite.logoutUsecase.Logout(ctx, refreshToken,"DeviceId")

    assert.NoError(suite.T(), err)
    suite.tokenRepoMock.AssertExpectations(suite.T())
}

func (suite *LogoutUsecaseSuite) TestLogout_Failure() {
    ctx, cancel := context.WithTimeout(context.Background(), suite.logoutUsecase.contextTimeout)
    defer cancel()
    refreshToken := "invalid_refresh_token"

    suite.tokenRepoMock.On("FindTokenByRefreshToken", ctx, refreshToken).Return(nil, errors.New("token not found"))

    err := suite.logoutUsecase.Logout(ctx, refreshToken,"deliceId")

    assert.Error(suite.T(), err)
    assert.Equal(suite.T(), "token not found", err.Error())
    suite.tokenRepoMock.AssertExpectations(suite.T())
}

func TestLogoutUsecaseSuite(t *testing.T) {
    suite.Run(t, new(LogoutUsecaseSuite))
}