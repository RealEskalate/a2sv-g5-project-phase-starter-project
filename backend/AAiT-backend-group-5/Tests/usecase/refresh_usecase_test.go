package usecase

import (
	"context"
	"testing"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/aait.backend.g5.main/backend/Mocks"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type RefreshUsecaseTestSuite struct {
	suite.Suite
	ctrl            *gomock.Controller
	mockJwtService  *mocks.MockJwtService
	mockSessionRepo *mocks.MockSessionRepository
	mockUserRepo    *mocks.MockUserRepository
	refreshUsecase  interfaces.RefreshUsecase
	ctx             context.Context
}

func (suite *RefreshUsecaseTestSuite) SetupTest() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.mockJwtService = mocks.NewMockJwtService(suite.ctrl)
	suite.mockSessionRepo = mocks.NewMockSessionRepository(suite.ctrl)
	suite.mockUserRepo = mocks.NewMockUserRepository(suite.ctrl)
	suite.ctx = context.Background()
	suite.refreshUsecase = usecases.NewRefreshUsecase(suite.mockJwtService, suite.mockSessionRepo, suite.mockUserRepo, suite.ctx)
}

func (suite *RefreshUsecaseTestSuite) TearDownTest() {
	suite.ctrl.Finish()
}

func (suite *RefreshUsecaseTestSuite) TestRefreshTokenSuccess() {
	userID := "user123"
	refreshToken := "validRefreshToken"
	accessToken := "newAccessToken"
	newRefreshToken := "newRefreshToken"
	user := &models.User{ID: userID}
	session := &models.Session{UserID: userID, RefreshToken: refreshToken}

	suite.mockSessionRepo.EXPECT().
		GetToken(suite.ctx, userID).
		Return(session, nil)

	suite.mockUserRepo.EXPECT().
		GetUserByID(suite.ctx, userID).
		Return(user, nil)

	suite.mockJwtService.EXPECT().
		CreateAccessToken(*user, 60).
		Return(accessToken, nil)

	suite.mockJwtService.EXPECT().
		CreateRefreshToken(*user, 60).
		Return(newRefreshToken, nil)

	suite.mockSessionRepo.EXPECT().
		UpdateToken(suite.ctx, &models.Session{
			UserID:       userID,
			AccessToken:  accessToken,
			RefreshToken: newRefreshToken,
		}).
		Return(nil)

	token, err := suite.refreshUsecase.RefreshToken(userID, refreshToken)

	suite.Nil(err)
	suite.Equal(accessToken, token)
}

func (suite *RefreshUsecaseTestSuite) TestRefreshTokenInvalidToken() {
	userID := "user123"
	refreshToken := "invalidRefreshToken"
	session := &models.Session{UserID: userID, RefreshToken: "validRefreshToken"}

	suite.mockSessionRepo.EXPECT().
		GetToken(suite.ctx, userID).
		Return(session, nil)

	token, err := suite.refreshUsecase.RefreshToken(userID, refreshToken)

	suite.NotNil(err)
	suite.Equal("Invalid refresh token", err.Message)
	suite.Equal("", token)
}

func (suite *RefreshUsecaseTestSuite) TestRefreshTokenSessionNotFound() {
	userID := "user123"
	refreshToken := "someRefreshToken"

	suite.mockSessionRepo.EXPECT().
		GetToken(suite.ctx, userID).
		Return(nil, models.Unauthorized("Session not found"))

	token, err := suite.refreshUsecase.RefreshToken(userID, refreshToken)

	suite.NotNil(err)
	suite.Equal("Session not found", err.Message)
	suite.Equal("", token)
}

func (suite *RefreshUsecaseTestSuite) TestRefreshTokenUserNotFound() {
	userID := "user123"
	refreshToken := "validRefreshToken"
	session := &models.Session{UserID: userID, RefreshToken: refreshToken}

	suite.mockSessionRepo.EXPECT().
		GetToken(suite.ctx, userID).
		Return(session, nil)

	suite.mockUserRepo.EXPECT().
		GetUserByID(suite.ctx, userID).
		Return(nil, models.BadRequest("User not found"))

	token, err := suite.refreshUsecase.RefreshToken(userID, refreshToken)

	suite.NotNil(err)
	suite.Equal("User not found", err.Message)
	suite.Equal("", token)
}

func TestRefreshUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(RefreshUsecaseTestSuite))
}
