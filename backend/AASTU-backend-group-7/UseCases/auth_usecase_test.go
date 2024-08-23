package usecases_test

import (
	"blogapp/Domain"
	"blogapp/Dtos"
	usecases "blogapp/UseCases"
	"blogapp/mocks"
	"context"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthUsecaseSuite struct {
	suite.Suite
	context     context.Context
	authUsecase *usecases.AuthUseCase
	repo        *mocks.AuthRepository
}

func (suite *AuthUsecaseSuite) SetupTest() {
	suite.repo = new(mocks.AuthRepository)
	suite.authUsecase = usecases.NewAuthUseCase(suite.repo)
	suite.context = context.Background()
}

func (suite *AuthUsecaseSuite) TestLogin() {
	c, _ := gin.CreateTestContext(nil)
	user := Dtos.LoginUserDto{
		Email:    "asdfs@gmail.com",
		Password: "password",
	}
	suite.repo.On("Login", mock.Anything, mock.Anything).Return(Domain.Tokens{}, nil, 200)
	_, err, status := suite.authUsecase.Login(c, &user)
	suite.Nil(err)
	suite.Equal(200, status)

}
func (suite *AuthUsecaseSuite) TestRegister() {
	c, _ := gin.CreateTestContext(nil)
	// user := Domain.User{
	// 	Email:    "",
	// 	Password: "password",
	// }
	suite.repo.On("Register", mock.Anything, mock.Anything).Return(&Domain.OmitedUser{}, nil, 200)
	_, err, status := suite.authUsecase.Register(c, &Dtos.RegisterUserDto{})
	suite.Nil(err)
	suite.Equal(200, status)
}

func (suite *AuthUsecaseSuite) TestLogout() {
	c, _ := gin.CreateTestContext(nil)
	suite.repo.On("Logout", mock.Anything, mock.Anything).Return(nil, 200)
	err, status := suite.authUsecase.Logout(c, primitive.NewObjectID())
	suite.Nil(err)
	suite.Equal(200, status)
}

func (suite *AuthUsecaseSuite) TestForgetPassword() {
	c, _ := gin.CreateTestContext(nil)
	suite.repo.On("ForgetPassword", mock.Anything, mock.Anything).Return(nil, 200)
	err, status := suite.authUsecase.ForgetPassword(c, "")
	suite.Nil(err)
	suite.Equal(200, status)
}

func (suite *AuthUsecaseSuite) TestResetPassword() {
	c, _ := gin.CreateTestContext(nil)
	suite.repo.On("ResetPassword", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, 200)
	err, status := suite.authUsecase.ResetPassword(c, "", "", "")
	suite.Nil(err)
	suite.Equal(200, status)
}

func (suite *AuthUsecaseSuite) TestGoogleLogin() {
	c, _ := gin.CreateTestContext(nil)
	suite.repo.On("GoogleLogin", mock.Anything).Return("")
	suite.authUsecase.GoogleLogin(c)
}

func (suite *AuthUsecaseSuite) TestCallbackHandler() {
	c, _ := gin.CreateTestContext(nil)
	suite.repo.On("CallbackHandler", mock.Anything, mock.Anything).Return(Domain.Tokens{}, nil, 200)
	_, err, status := suite.authUsecase.CallbackHandler(c, "")
	suite.Nil(err)
	suite.Equal(200, status)
}

func (suite *AuthUsecaseSuite) TestActivateAccount() {
	c, _ := gin.CreateTestContext(nil)
	suite.repo.On("ActivateAccount", mock.Anything, mock.Anything).Return(nil, 200)
	err, status := suite.authUsecase.ActivateAccount(c, "")
	suite.Nil(err)
	suite.Equal(200, status)
}

func TestAuthUsecaseSuite(t *testing.T) {
	suite.Run(t, new(AuthUsecaseSuite))
}
