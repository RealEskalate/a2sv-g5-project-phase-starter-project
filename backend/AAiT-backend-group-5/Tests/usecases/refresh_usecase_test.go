package usecases_test

import (
	"context"
	"testing"

	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type RefreshUsecaseTestSuite struct {
	suite.Suite
	jwtServiceMock        *mocks.MockJwtService
	sessionRepositoryMock *mocks.MockSessionRepository
	userRepositoryMock    *mocks.MockUserRepository
	refreshUsecase        interfaces.RefreshUsecase
	ctrl                  *gomock.Controller
}

func (suite *RefreshUsecaseTestSuite) SetupSuite() {
	suite.ctrl = gomock.NewController(suite.T())
	suite.jwtServiceMock = mocks.NewMockJwtService(suite.ctrl)
	suite.sessionRepositoryMock = mocks.NewMockSessionRepository(suite.ctrl)
	suite.userRepositoryMock = mocks.NewMockUserRepository(suite.ctrl)
	suite.refreshUsecase = usecases.NewRefreshUsecase(
		suite.jwtServiceMock,
		suite.sessionRepositoryMock,
		suite.userRepositoryMock,
	)
}

func (suite *RefreshUsecaseTestSuite) TearDownSuite() {
	suite.ctrl.Finish()
}

func (suite *RefreshUsecaseTestSuite) TestRefreshToken_Success() {
	ctx := context.Background()
	userID := "user1"
	refreshToken := "valid_refresh_token"
	newAccessToken := "new_access_token"

	user := &models.User{
		ID:    userID,
		Email: "user@example.com",
	}

	session := &models.Session{
		UserID:       userID,
		RefreshToken: refreshToken,
		AccessToken:  newAccessToken,
	}

	suite.sessionRepositoryMock.
		EXPECT().
		GetToken(ctx, userID).
		Return(session, nil)

	suite.userRepositoryMock.
		EXPECT().
		GetUserByID(ctx, userID).
		Return(user, nil)

	suite.jwtServiceMock.
		EXPECT().
		CreateAccessToken(*user, 60).
		Return(newAccessToken, nil)
	suite.sessionRepositoryMock.
		EXPECT().
		UpdateToken(ctx, session).
		Return(nil)

	accessToken, err := suite.refreshUsecase.RefreshToken(ctx, userID, refreshToken)
	suite.Nil(err)
	suite.Equal(newAccessToken, accessToken)
}

func (suite *RefreshUsecaseTestSuite) TestRefreshToken_InvalidRefreshToken() {
	ctx := context.Background()
	userID := "user1"
	refreshToken := "invalid_refresh_token"

	session := &models.Session{
		UserID:       userID,
		RefreshToken: "different_refresh_token",
	}

	suite.sessionRepositoryMock.
		EXPECT().
		GetToken(ctx, userID).
		Return(session, nil)

	accessToken, err := suite.refreshUsecase.RefreshToken(ctx, userID, refreshToken)
	suite.Equal(models.Unauthorized("Invalid refresh token"), err)
	suite.Empty(accessToken)
}

func (suite *RefreshUsecaseTestSuite) TestRefreshToken_GetSessionError() {
	ctx := context.Background()
	userID := "user1"
	refreshToken := "any_refresh_token"

	suite.sessionRepositoryMock.
		EXPECT().
		GetToken(ctx, userID).
		Return(nil, models.InternalServerError("Error fetching session"))

	accessToken, err := suite.refreshUsecase.RefreshToken(ctx, userID, refreshToken)
	suite.Equal(models.InternalServerError("Error fetching session"), err)
	suite.Empty(accessToken)
}

func (suite *RefreshUsecaseTestSuite) TestRefreshToken_GetUserError() {
	ctx := context.Background()
	userID := "user1"
	refreshToken := "valid_refresh_token"

	session := &models.Session{
		UserID:       userID,
		RefreshToken: refreshToken,
	}

	suite.sessionRepositoryMock.
		EXPECT().
		GetToken(ctx, userID).
		Return(session, nil)

	suite.userRepositoryMock.
		EXPECT().
		GetUserByID(ctx, userID).
		Return(nil, models.InternalServerError("Error fetching user"))

	accessToken, err := suite.refreshUsecase.RefreshToken(ctx, userID, refreshToken)
	suite.Equal(models.InternalServerError("Error fetching user"), err)
	suite.Empty(accessToken)
}

func (suite *RefreshUsecaseTestSuite) TestRefreshToken_CreateAccessTokenError() {
	ctx := context.Background()
	userID := "user1"
	refreshToken := "valid_refresh_token"

	user := &models.User{
		ID:    userID,
		Email: "user@example.com",
	}

	session := &models.Session{
		UserID:       userID,
		RefreshToken: refreshToken,
	}

	suite.sessionRepositoryMock.
		EXPECT().
		GetToken(ctx, userID).
		Return(session, nil)

	suite.userRepositoryMock.
		EXPECT().
		GetUserByID(ctx, userID).
		Return(user, nil)

	suite.jwtServiceMock.
		EXPECT().
		CreateAccessToken(*user, 60).
		Return("", models.InternalServerError("Error creating access token"))

	

	accessToken, err := suite.refreshUsecase.RefreshToken(ctx, userID, refreshToken)
	suite.Equal(models.InternalServerError("An unexpected error occurred"), err)
	suite.Empty(accessToken)
}

func (suite *RefreshUsecaseTestSuite) TestRefreshToken_UpdateTokenError() {
	ctx := context.Background()
	userID := "user1"
	refreshToken := "valid_refresh_token"

	user := &models.User{
		ID:    userID,
		Email: "user@example.com",
	}

	session := &models.Session{
		UserID:       userID,
		RefreshToken: refreshToken,
	}

	newAccessToken := "new_access_token"
	suite.sessionRepositoryMock.
		EXPECT().
		GetToken(ctx, userID).
		Return(session, nil)

	suite.userRepositoryMock.
		EXPECT().
		GetUserByID(ctx, userID).
		Return(user, nil)

	suite.jwtServiceMock.
		EXPECT().
		CreateAccessToken(*user, 60).
		Return(newAccessToken, nil)

	suite.sessionRepositoryMock.
		EXPECT().
		UpdateToken(ctx, &models.Session{
			UserID:       userID,
			RefreshToken: refreshToken,
			AccessToken:  newAccessToken,
		}).
		Return(models.InternalServerError("Error updating token"))

	accessToken, err := suite.refreshUsecase.RefreshToken(ctx, userID, refreshToken)
	suite.Equal(models.InternalServerError("Error updating token"), err)
	suite.Empty(accessToken)
}

func TestRefreshUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(RefreshUsecaseTestSuite))
}
