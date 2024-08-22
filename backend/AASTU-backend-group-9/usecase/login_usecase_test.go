package usecase

import (
    "context"
    "errors"
    "testing"
    "time"

   
	 "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"

    "blog/domain"
    "blog/domain/mocks"
)

type LoginUsecaseSuite struct {
    suite.Suite
    userRepoMock   *mocks.UserRepository
    tokenRepoMock  *mocks.TokenRepository
    loginUsecase   *loginUsecase
}

func (suite *LoginUsecaseSuite) SetupTest() {
    suite.userRepoMock = new(mocks.UserRepository)
    suite.tokenRepoMock = new(mocks.TokenRepository)
    suite.loginUsecase = &loginUsecase{
        userRepository:  suite.userRepoMock,
        tokenRepository: suite.tokenRepoMock,
        contextTimeout:  time.Second * 2,
    }
}

func (suite *LoginUsecaseSuite) TestAuthenticateUser_Success() {
    ctx, cancel := context.WithTimeout(context.Background(), suite.loginUsecase.contextTimeout)
    defer cancel()
    login := &domain.AuthLogin{Username: "testuser", Password: "password"}
    
    // Convert string to primitive.ObjectID
    userID, err := primitive.ObjectIDFromHex("507f1f77bcf86cd799439011") 
    if err != nil {
        suite.T().Fatal(err)
    }
    
    user := &domain.User{ID: userID, Username: "testuser"}

    suite.userRepoMock.On("GetUserByEmail", mock.Anything, login.Username).Return(user, nil)
    suite.userRepoMock.On("CheckPassword", mock.Anything, user, login.Password).Return(true, nil)

    result, err := suite.loginUsecase.AuthenticateUser(ctx, login)

    assert.NoError(suite.T(), err)
    assert.Equal(suite.T(), user, result)
    suite.userRepoMock.AssertExpectations(suite.T())
}

func (suite *LoginUsecaseSuite) TestAuthenticateUser_Failure() {
    ctx, cancel := context.WithTimeout(context.Background(), suite.loginUsecase.contextTimeout)
    defer cancel()
    login := &domain.AuthLogin{Username: "testuser", Password: "wrongpassword"}

    suite.userRepoMock.On("GetUserByEmail", mock.Anything, login.Username).Return(nil, errors.New("user not found"))

    result, err := suite.loginUsecase.AuthenticateUser(ctx, login)

    assert.Error(suite.T(), err)
    assert.Nil(suite.T(), result)
    suite.userRepoMock.AssertExpectations(suite.T())
}

func TestLoginUsecaseSuite(t *testing.T) {
    suite.Run(t, new(LoginUsecaseSuite))
}