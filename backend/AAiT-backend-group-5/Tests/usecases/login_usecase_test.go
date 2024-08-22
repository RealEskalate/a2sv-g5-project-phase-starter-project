package usecases_test

import (
	"context"
	"testing"

	config "github.com/aait.backend.g5.main/backend/Config"
	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	mocks "github.com/aait.backend.g5.main/backend/Mocks"
	usecases "github.com/aait.backend.g5.main/backend/UseCases"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

type LoginUsecaseTestSuite struct {
	suite.Suite
	jwtServiceMock      *mocks.MockJwtService
	passwordServiceMock *mocks.MockPasswordService
	repositoryMock      *mocks.MockUserRepository
	sessionMock         *mocks.MockSessionRepository
	env                 config.Env
	loginUsecase        interfaces.LoginUsecase
	ctr                 *gomock.Controller
}

func (suite *LoginUsecaseTestSuite) SetupSuite() {
	suite.ctr = gomock.NewController(suite.T())
	suite.jwtServiceMock = mocks.NewMockJwtService(suite.ctr)
	suite.passwordServiceMock = mocks.NewMockPasswordService(suite.ctr)
	suite.repositoryMock = mocks.NewMockUserRepository(suite.ctr)
	suite.sessionMock = mocks.NewMockSessionRepository(suite.ctr)

	// Provide environment values
	suite.env = config.Env{
		ACCESS_TOKEN_EXPIRY_HOUR:  1,
		REFRESH_TOKEN_EXPIRY_HOUR: 24,
	}

	suite.loginUsecase = usecases.NewLoginUsecase(
		suite.jwtServiceMock,
		suite.passwordServiceMock,
		suite.repositoryMock,
		suite.sessionMock,
		suite.env,
	)
}

func (suite *LoginUsecaseTestSuite) TearDownSuite() {
	suite.ctr.Finish()
}

func (suite *LoginUsecaseTestSuite) TestLoginUser_Success() {
	ctx := context.Background()
	userRequest := dtos.LoginRequest{
		UsernameOrEmail: "user@example.com",
		Password:        "password123",
	}

	user := &models.User{
		ID:       "user1",
		Email:    "user@example.com",
		Password: "hashed_password",
	}

	accessToken := "access_token"
	refreshToken := "refresh_token"

	suite.repositoryMock.
		EXPECT().
		GetUserByEmailOrUsername(ctx, userRequest.UsernameOrEmail, userRequest.UsernameOrEmail).
		Return(user, nil)

	suite.passwordServiceMock.
		EXPECT().
		ValidatePassword(userRequest.Password, user.Password).
		Return(true)

	suite.jwtServiceMock.
		EXPECT().
		CreateAccessToken(*user, suite.env.ACCESS_TOKEN_EXPIRY_HOUR).
		Return(accessToken, nil)

	suite.jwtServiceMock.
		EXPECT().
		CreateRefreshToken(*user, suite.env.REFRESH_TOKEN_EXPIRY_HOUR).
		Return(refreshToken, nil)

	suite.sessionMock.
		EXPECT().
		GetToken(ctx, user.ID).
		Return(nil, nil)
	suite.sessionMock.
		EXPECT().
		SaveToken(ctx, &models.Session{
			UserID:       user.ID,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}).
		Return(nil)

	result, err := suite.loginUsecase.LoginUser(ctx, userRequest)
	if err != nil {
		suite.T().Errorf("expected no error, got %v", err)
	}
	suite.Equal(&dtos.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, result)
}

func (suite *LoginUsecaseTestSuite) TestLoginUser_UserNotFound() {
	ctx := context.Background()
	userReqest := dtos.LoginRequest{
		UsernameOrEmail: "user@example.com",
		Password:        "password123",
	}

	suite.repositoryMock.
		EXPECT().
		GetUserByEmailOrUsername(ctx, userReqest.UsernameOrEmail, userReqest.UsernameOrEmail).
		Return(nil, models.NotFound("User not found"))

	result, err := suite.loginUsecase.LoginUser(ctx, userReqest)
	if err == nil {
		suite.T().Error("expected error, got none")
	}
	suite.Equal(models.NotFound("User not found"), err)
	suite.Nil(result)
}

func (suite *LoginUsecaseTestSuite) TestLoginUser_InvalidPassword() {
	ctx := context.Background()
	userReqest := dtos.LoginRequest{
		UsernameOrEmail: "user@example.com",
		Password:        "password123",
	}

	user := &models.User{
		ID:       "user1",
		Email:    "user@example.com",
		Password: "hashed_password",
	}

	suite.repositoryMock.
		EXPECT().
		GetUserByEmailOrUsername(ctx, userReqest.UsernameOrEmail, userReqest.UsernameOrEmail).
		Return(user, nil)

	suite.passwordServiceMock.
		EXPECT().
		ValidatePassword(userReqest.Password, user.Password).
		Return(false)

	result, err := suite.loginUsecase.LoginUser(ctx, userReqest)
	if err == nil {
		suite.T().Error("expected error, got none")
	}
	suite.Equal(models.Unauthorized("Invalid creaditional"), err)
	suite.Nil(result)
}

func (suite *LoginUsecaseTestSuite) TestLoginUser_TokenGenerationError() {
	ctx := context.Background()
	userReqest := dtos.LoginRequest{
		UsernameOrEmail: "user@example.com",
		Password:        "password123",
	}

	user := &models.User{
		ID:       "user1",
		Email:    "user@example.com",
		Password: "hashed_password",
	}

	suite.repositoryMock.
		EXPECT().
		GetUserByEmailOrUsername(ctx, userReqest.UsernameOrEmail, userReqest.UsernameOrEmail).
		Return(user, nil)

	suite.passwordServiceMock.
		EXPECT().
		ValidatePassword(userReqest.Password, user.Password).
		Return(true)

	suite.jwtServiceMock.
		EXPECT().
		CreateAccessToken(*user, suite.env.ACCESS_TOKEN_EXPIRY_HOUR).
		Return("", models.InternalServerError("Error occurred while generating the access token"))

	suite.jwtServiceMock.
		EXPECT().
		CreateRefreshToken(*user, suite.env.REFRESH_TOKEN_EXPIRY_HOUR).
		Return("", models.InternalServerError("Error occurred while creating the refresh token"))

	result, err := suite.loginUsecase.LoginUser(ctx, userReqest)
	if err == nil {
		suite.T().Error("expected error, got none")
	}
	suite.Equal(models.InternalServerError("Something went wrong"), err)
	suite.Nil(result)
}

func (suite *LoginUsecaseTestSuite) TestLoginUser_UpdateTokenError() {
	ctx := context.Background()
	userReqest := dtos.LoginRequest{
		UsernameOrEmail: "user@example.com",
		Password:        "password123",
	}

	user := &models.User{
		ID:       "user1",
		Email:    "user@example.com",
		Password: "hashed_password",
	}

	accessToken := "access_token"
	refreshToken := "refresh_token"

	suite.repositoryMock.
		EXPECT().
		GetUserByEmailOrUsername(ctx, userReqest.UsernameOrEmail, userReqest.UsernameOrEmail).
		Return(user, nil)

	suite.passwordServiceMock.
		EXPECT().
		ValidatePassword(userReqest.Password, user.Password).
		Return(true)

	suite.jwtServiceMock.
		EXPECT().
		CreateAccessToken(*user, suite.env.ACCESS_TOKEN_EXPIRY_HOUR).
		Return(accessToken, nil)

	suite.jwtServiceMock.
		EXPECT().
		CreateRefreshToken(*user, suite.env.REFRESH_TOKEN_EXPIRY_HOUR).
		Return(refreshToken, nil)

	suite.sessionMock.
		EXPECT().
		GetToken(ctx, user.ID).
		Return(&models.Session{
			UserID:       user.ID,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}, nil)

	suite.sessionMock.
		EXPECT().
		UpdateToken(ctx, &models.Session{
			UserID:       user.ID,
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}).
		Return(models.InternalServerError("Error updating token"))

	result, err := suite.loginUsecase.LoginUser(ctx, userReqest)
	if err == nil {
		suite.T().Error("expected error, got none")
	}
	suite.Equal(models.InternalServerError("Error updating token"), err)
	suite.Nil(result)
}

func TestLoginUsecaseTestSuite(t *testing.T) {
	suite.Run(t, new(LoginUsecaseTestSuite))
}
