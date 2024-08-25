package usecase

// import (
// 	"Blog_Starter/domain"
// 	"Blog_Starter/domain/mocks"
// 	"Blog_Starter/utils"
// 	"context"
// 	"errors"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/suite"
// 	"golang.org/x/crypto/bcrypt"
// )

// type LoginUseCaseTestSuite struct {
// 	suite.Suite
// 	loginUseCase    domain.LoginUsecase
// 	userRepoMock    *mocks.UserRepository
// 	tokenManagerMock *mocks.TokenManager
// 	envMock         *utils.Env
// }

// func (suite *LoginUseCaseTestSuite) SetupTest() {
// 	suite.userRepoMock = new(mocks.UserRepository)
// 	suite.tokenManagerMock = new(mocks.TokenManager)
// 	suite.envMock = &utils.Env{
// 		AccessTokenSecret:     "test_secret",
// 		RefreshTokenSecret:    "test_secret",
// 		AccessTokenExpiryHour: 1,
// 		RefreshTokenExpiryHour: 24,
// 	}
// 	suite.loginUseCase = NewLoginUseCase(suite.userRepoMock, suite.tokenManagerMock, time.Second*2, suite.envMock)
// }

// func (suite *LoginUseCaseTestSuite) TestLogin_Success() {
// 	ctx := context.TODO()
// 	email := "test@example.com"
// 	password := "password123"
// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	user := &domain.User{Email: email, Password: string(hashedPassword), IsActivated: true}

// 	suite.userRepoMock.On("GetUserByEmail", ctx, email).Return(user, nil)
// 	suite.tokenManagerMock.On("CreateAccessToken", user, suite.envMock.AccessTokenSecret, suite.envMock.AccessTokenExpiryHour).Return("access_token", nil)
// 	suite.tokenManagerMock.On("CreateRefreshToken", user, suite.envMock.RefreshTokenSecret, suite.envMock.RefreshTokenExpiryHour).Return("refresh_token", nil)
// 	suite.userRepoMock.On("UpdateToken", ctx, "access_token", "refresh_token", user.UserID.Hex()).Return(nil, nil)

// 	req := &domain.UserLogin{Email: email, Password: password}
// 	res, err := suite.loginUseCase.Login(ctx, req)

// 	assert.Nil(suite.T(), err)
// 	assert.Equal(suite.T(), "access_token", res.AccessToken)
// 	assert.Equal(suite.T(), "refresh_token", res.RefreshToken)
// }

// func (suite *LoginUseCaseTestSuite) TestLogin_UserNotActivated() {
// 	ctx := context.TODO()
// 	email := "test@example.com"
// 	password := "password123"
// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
// 	user := &domain.User{Email: email, Password: string(hashedPassword), IsActivated: false}

// 	suite.userRepoMock.On("GetUserByEmail", ctx, email).Return(user, nil)

// 	req := &domain.UserLogin{Email: email, Password: password}
// 	_, err := suite.loginUseCase.Login(ctx, req)

// 	assert.NotNil(suite.T(), err)
// 	assert.EqualError(suite.T(), err, "user is not activated, Verify your email")
// }

// func (suite *LoginUseCaseTestSuite) TestLogin_InvalidPassword() {
// 	ctx := context.TODO()
// 	email := "test@example.com"
// 	password := "password123"
// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("wrongpassword"), bcrypt.DefaultCost)
// 	user := &domain.User{Email: email, Password: string(hashedPassword), IsActivated: true}

// 	suite.userRepoMock.On("GetUserByEmail", ctx, email).Return(user, nil)

// 	req := &domain.UserLogin{Email: email, Password: password}
// 	_, err := suite.loginUseCase.Login(ctx, req)

// 	assert.NotNil(suite.T(), err)
// 	assert.EqualError(suite.T(), err, "password incorrect")
// }

// func (suite *LoginUseCaseTestSuite) TestUpdatePassword_Success() {
// 	ctx := context.TODO()
// 	userID := "user_id"
// 	newPassword := "new_password"
// 	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)

// 	suite.userRepoMock.On("UpdatePassword", ctx, string(hashedPassword), userID).Return(nil, nil)

// 	req := domain.ChangePasswordRequest{Password: newPassword}
// 	err := suite.loginUseCase.UpdatePassword(ctx, req, userID)

// 	assert.Nil(suite.T(), err)
// }

// func (suite *LoginUseCaseTestSuite) TestLogOut_Success() {
// 	ctx := context.TODO()
// 	userID := "user_id"

// 	suite.userRepoMock.On("UpdateToken", ctx, "", "", userID).Return(nil, nil)

// 	err := suite.loginUseCase.LogOut(ctx, userID)

// 	assert.Nil(suite.T(), err)
// }

// func TestLoginUseCaseTestSuite(t *testing.T) {
// 	suite.Run(t, new(LoginUseCaseTestSuite))
// }
