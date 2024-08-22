package controllers_test

// import (
// 	"blogapp/Domain"
// 	"blogapp/Dtos"
// 	jwtservice "blogapp/Infrastructure/jwt_service"
// 	"blogapp/mocks"
// 	"bytes"
// 	"encoding/json"
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	controllers "blogapp/Delivery/controllers"

// 	"bou.ke/monkey"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/suite"
// 	"go.mongodb.org/mongo-driver/bson/primitive"
// )

// type AuthControllerTestSuite struct {
// 	suite.Suite
// 	mockAuthUseCase *mocks.AuthUseCase
// 	userController  *controllers.AuthController
// 	patch           *monkey.PatchGuard
// 	userID          primitive.ObjectID
// 	Role            string
// }

// func (suite *AuthControllerTestSuite) SetupTest() {
// 	suite.mockAuthUseCase = new(mocks.AuthUseCase)
// 	var err error
// 	suite.userController = controllers.NewAuthController(suite.mockAuthUseCase)
// 	assert.NoError(suite.T(), err)
// 	suite.userID = user_id
// 	suite.Role = "admin"
// 	suite.patch = monkey.Patch(controllers.Getclaim, mockExtractUser)

// 	suite.patch = monkey.Patch(jwtservice.VerifyToken, func(token string) (string, error) {
// 		fmt.Println("Mocked VerifyToken")
// 		return "valid_token", nil
// 	})
// }

// func (suite *AuthControllerTestSuite) TestLogin() {
// 	// Arrange
// 	user := Domain.User{Email: "test@example.com", Password: "password"}
// 	suite.mockAuthUseCase.On("Login", mock.Anything, &user).Return(Domain.Tokens{}, nil, 200).Once()

// 	// Marshal user to JSON
// 	userJSON, _ := json.Marshal(user)
// 	req, _ := http.NewRequest(http.MethodPost, "/auth/login", bytes.NewBuffer(userJSON))
// 	req.Header.Set("Content-Type", "application/json")

// 	w := httptest.NewRecorder()
// 	c, _ := gin.CreateTestContext(w)
// 	c.Request = req

// 	// Act
// 	suite.userController.Login(c)

// 	// Assert
// 	assert.Equal(suite.T(), http.StatusOK, w.Code)
// 	suite.mockAuthUseCase.AssertExpectations(suite.T())
// }

// func TestRegister(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	t.Run("Success", func(t *testing.T) {
// 		// Arrange
// 		mockAuthUseCase := new(mocks.AuthUseCase)
// 		authController := controllers.NewAuthController(mockAuthUseCase)

// 		newUser := Dtos.RegisterUserDto{
// 			Email:    "test@example.com",
// 			Password: "StrongPassword123!",
// 			UserName: "testuser",
// 		}

// 		createdUser := Domain.OmitedUser{
// 			Email: newUser.Email,
// 		}

// 		mockAuthUseCase.On("Register", mock.Anything, &newUser).Return(&createdUser, nil, http.StatusCreated)

// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)

// 		jsonValue, _ := json.Marshal(newUser)
// 		c.Request, _ = http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonValue))
// 		c.Request.Header.Set("Content-Type", "application/json")

// 		// Act
// 		authController.Register(c)

// 		// Assert
// 		assert.Equal(t, http.StatusCreated, w.Code)
// 		mockAuthUseCase.AssertExpectations(t)
// 	})

// 	t.Run("Invalid JSON", func(t *testing.T) {
// 		// Arrange
// 		mockAuthUseCase := new(mocks.AuthUseCase)
// 		authController := controllers.NewAuthController(mockAuthUseCase)

// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)

// 		c.Request, _ = http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer([]byte("{invalid json}")))
// 		c.Request.Header.Set("Content-Type", "application/json")

// 		// Act
// 		authController.Register(c)

// 		// Assert
// 		assert.Equal(t, http.StatusBadRequest, w.Code)
// 	})

// 	t.Run("Validation Error", func(t *testing.T) {
// 		// Arrange
// 		mockAuthUseCase := new(mocks.AuthUseCase)
// 		authController := controllers.NewAuthController(mockAuthUseCase)

// 		newUser := Dtos.RegisterUserDto{
// 			Email:    "invalid-email",
// 			Password: "weakpassword",
// 			UserName: "testuser",
// 		}

// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)

// 		jsonValue, _ := json.Marshal(newUser)
// 		c.Request, _ = http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonValue))
// 		c.Request.Header.Set("Content-Type", "application/json")

// 		// Act
// 		authController.Register(c)

// 		// Assert
// 		assert.Equal(t, http.StatusBadRequest, w.Code)
// 	})

// 	t.Run("Registration Error", func(t *testing.T) {
// 		// Arrange
// 		mockAuthUseCase := new(mocks.AuthUseCase)
// 		authController := controllers.NewAuthController(mockAuthUseCase)

// 		newUser := Dtos.RegisterUserDto{
// 			Email:    "test@example.com",
// 			Password: "StrongPassword123!",
// 			UserName: "testuser",
// 		}

// 		mockAuthUseCase.On("Register", mock.Anything, &newUser).Return(nil, errors.New("registration error"), http.StatusInternalServerError)

// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)

// 		jsonValue, _ := json.Marshal(newUser)
// 		c.Request, _ = http.NewRequest(http.MethodPost, "/register", bytes.NewBuffer(jsonValue))
// 		c.Request.Header.Set("Content-Type", "application/json")

// 		// Act
// 		authController.Register(c)

// 		// Assert
// 		assert.Equal(t, http.StatusInternalServerError, w.Code)
// 		mockAuthUseCase.AssertExpectations(t)
// 	})
// }

// func TestForgetPassword(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	t.Run("Success", func(t *testing.T) {
// 		// Arrange
// 		mockAuthUseCase := new(mocks.AuthUseCase)
// 		authController := controllers.NewAuthController(mockAuthUseCase)

// 		email := "test@example.com"
// 		mockAuthUseCase.On("ForgetPassword", mock.Anything, email).Return(nil, http.StatusOK)

// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)

// 		c.Request, _ = http.NewRequest(http.MethodPost, "/forget-password", strings.NewReader("email="+email))
// 		c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 		// Act
// 		authController.ForgetPassword(c)

// 		// Assert
// 		assert.Equal(t, http.StatusOK, w.Code)
// 		mockAuthUseCase.AssertExpectations(t)
// 	})

// 	t.Run("Error", func(t *testing.T) {
// 		// Arrange
// 		mockAuthUseCase := new(mocks.AuthUseCase)
// 		authController := controllers.NewAuthController(mockAuthUseCase)

// 		email := "test@example.com"
// 		mockAuthUseCase.On("ForgetPassword", mock.Anything, email).Return(errors.New("some error"), http.StatusInternalServerError)

// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)

// 		c.Request, _ = http.NewRequest(http.MethodPost, "/forget-password", strings.NewReader("email="+email))
// 		c.Request.Header.Set("Content-Type", "application/x-www-form-urlencoded")

// 		// Act
// 		authController.ForgetPassword(c)

// 		// Assert
// 		assert.Equal(t, http.StatusInternalServerError, w.Code)
// 		mockAuthUseCase.AssertExpectations(t)
// 	})
// }

// func TestResetPassword(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	t.Run("Invalid Token", func(t *testing.T) {
// 		mockAuthUseCase := new(mocks.AuthUseCase)
// 		authController := controllers.NewAuthController(mockAuthUseCase)

// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)

// 		c.Params = gin.Params{{Key: "reset_token", Value: "invalid_token"}}
// 		c.Request, _ = http.NewRequest(http.MethodPost, "/reset-password", nil)

// 		authController.ResetPassword(c)

// 		assert.Equal(t, http.StatusBadRequest, w.Code)
// 	})

// 	t.Run("Password Required", func(t *testing.T) {
// 		mockAuthUseCase := new(mocks.AuthUseCase)
// 		authController := controllers.NewAuthController(mockAuthUseCase)

// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)

// 		c.Params = gin.Params{{Key: "reset_token", Value: "valid_token"}}
// 		c.Request, _ = http.NewRequest(http.MethodPost, "/reset-password", nil)

// 		authController.ResetPassword(c)

// 		assert.Equal(t, http.StatusBadRequest, w.Code)
// 	})

// 	t.Run("Success", func(t *testing.T) {
// 		mockAuthUseCase := new(mocks.AuthUseCase)
// 		authController := controllers.NewAuthController(mockAuthUseCase)

// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)

// 		c.Params = gin.Params{{Key: "reset_token", Value: "valid_token"}}
// 		c.Request, _ = http.NewRequest(http.MethodPost, "/reset-password", bytes.NewBufferString("password=newpassword"))

// 		mockAuthUseCase.On("ResetPassword", mock.Anything, mock.Anything, mock.Anything, mock.Anything).Return(nil, http.StatusOK)

// 		authController.ResetPassword(c)

// 		assert.Equal(t, http.StatusOK, w.Code)
// 		mockAuthUseCase.AssertExpectations(t)
// 	})
// }

// func TestCallbackHandler(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	t.Run("Success", func(t *testing.T) {
// 		mockAuthUseCase := new(mocks.AuthUseCase)
// 		authController := controllers.NewAuthController(mockAuthUseCase)

// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)

// 		c.Request, _ = http.NewRequest(http.MethodGet, "/callback?code=valid_code", nil)

// 		token := Domain.Tokens{
// 			AccessToken:  "access_token",
// 			RefreshToken: "refresh_token",
// 		}

// 		mockAuthUseCase.On("CallbackHandler", mock.Anything, "valid_code").Return(&token, nil, http.StatusOK)

// 		authController.CallbackHandler(c)

// 		assert.Equal(t, http.StatusOK, w.Code)
// 		mockAuthUseCase.AssertExpectations(t)
// 	})

// 	t.Run("Error", func(t *testing.T) {
// 		mockAuthUseCase := new(mocks.AuthUseCase)
// 		authController := controllers.NewAuthController(mockAuthUseCase)

// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)

// 		c.Request, _ = http.NewRequest(http.MethodGet, "/callback?code=invalid_code", nil)

// 		mockAuthUseCase.On("CallbackHandler", mock.Anything, "invalid_code").Return(nil, errors.New("callback error"), http.StatusInternalServerError)

// 		authController.CallbackHandler(c)

// 		assert.Equal(t, http.StatusInternalServerError, w.Code)
// 		mockAuthUseCase.AssertExpectations(t)
// 	})
// }

// func TestLoginHandlerGoogle(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	t.Run("Success", func(t *testing.T) {
// 		mockAuthUseCase := new(mocks.AuthUseCase)
// 		authController := controllers.NewAuthController(mockAuthUseCase)

// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)

// 		mockAuthUseCase.On("GoogleLogin", mock.Anything).Return("http://google.com")

// 		authController.LoginHandlerGoogle(c)

// 		assert.Equal(t, http.StatusTemporaryRedirect, w.Code)
// 		mockAuthUseCase.AssertExpectations(t)
// 	})

// 	t.Run("Error", func(t *testing.T) {
// 		mockAuthUseCase := new(mocks.AuthUseCase)
// 		authController := controllers.NewAuthController(mockAuthUseCase)

// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)

// 		mockAuthUseCase.On("GoogleLogin", mock.Anything).Return("")

// 		authController.LoginHandlerGoogle(c)

// 		assert.Equal(t, http.StatusInternalServerError, w.Code)
// 		mockAuthUseCase.AssertExpectations(t)
// 	})
// }

// func TestActivateAccount(t *testing.T) {
// 	gin.SetMode(gin.TestMode)

// 	t.Run("Success", func(t *testing.T) {
// 		mockAuthUseCase := new(mocks.AuthUseCase)
// 		authController := controllers.NewAuthController(mockAuthUseCase)

// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)

// 		c.Params = gin.Params{{Key: "activation_token", Value: "valid_token"}}

// 		mockAuthUseCase.On("ActivateAccount", mock.Anything, "valid_token").Return(nil, http.StatusOK)

// 		authController.ActivateAccount(c)

// 		assert.Equal(t, http.StatusOK, w.Code)
// 		mockAuthUseCase.AssertExpectations(t)
// 	})

// 	t.Run("Error", func(t *testing.T) {
// 		mockAuthUseCase := new(mocks.AuthUseCase)
// 		authController := controllers.NewAuthController(mockAuthUseCase)

// 		w := httptest.NewRecorder()
// 		c, _ := gin.CreateTestContext(w)

// 		c.Params = gin.Params{{Key: "activation_token", Value: "invalid_token"}}

// 		mockAuthUseCase.On("ActivateAccount", mock.Anything, "invalid_token").Return(errors.New("activation error"), http.StatusInternalServerError)

// 		authController.ActivateAccount(c)

// 		assert.Equal(t, http.StatusInternalServerError, w.Code)
// 		mockAuthUseCase.AssertExpectations(t)
// 	})
// }

// func TestAuthController(t *testing.T) {
// 	suite.Run(t, new(AuthControllerTestSuite))
// }
