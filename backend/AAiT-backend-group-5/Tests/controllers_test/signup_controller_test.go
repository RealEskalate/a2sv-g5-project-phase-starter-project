package controller

// import (
// 	"bytes"
// 	"encoding/json"
// 	"net/http"
// 	"net/http/httptest"
// 	"testing"

// 	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
// 	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
// 	models "github.com/aait.backend.g5.main/backend/Domain/Models"
// 	mocks "github.com/aait.backend.g5.main/backend/Mocks/usecase_mock"
// 	"github.com/gin-gonic/gin"
// 	"github.com/stretchr/testify/mock"
// 	"github.com/stretchr/testify/suite"
// )

// type SignupControllerTestSuite struct {
// 	suite.Suite
// 	mockSignupUsecase   *mocks.SignupUsecase
// 	mockPasswordUsecase *mocks.PasswordUsecase
// 	controller          *controllers.SignupController
// 	router              *gin.Engine
// }

// func (suite *SignupControllerTestSuite) SetupSuite() {
// 	suite.mockSignupUsecase = new(mocks.SignupUsecase)
// 	suite.mockPasswordUsecase = new(mocks.PasswordUsecase)
// 	suite.controller = controllers.NewSignupController(suite.mockSignupUsecase, suite.mockPasswordUsecase)
// 	suite.router = gin.Default()

// 	// Define the routes
// 	suite.router.POST("/signup", suite.controller.Signup)
// 	suite.router.POST("/confirm/:id", suite.controller.ConfirmRegistration)
// }

// func (suite *SignupControllerTestSuite) TearDownSuite() {
// 	suite.mockSignupUsecase.AssertExpectations(suite.T())
// 	suite.mockPasswordUsecase.AssertExpectations(suite.T())
// }

// func (suite *SignupControllerTestSuite) TestSignup_Success() {
// 	newUser := &models.User{
// 		Username: "testuser",
// 		Name:     "Test User",
// 		Email:    "test@example.com",
// 	}

// 	suite.mockSignupUsecase.On("CreateUser", mock.Anything, newUser).Return(nil).Once()

// 	body, _ := json.Marshal(newUser)
// 	request, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
// 	request.Header.Set("Content-Type", "application/json")

// 	responseWriter := httptest.NewRecorder()
// 	suite.router.ServeHTTP(responseWriter, request)

// 	suite.Equal(http.StatusOK, responseWriter.Code)
// 	suite.Contains(responseWriter.Body.String(), "check your email")
// }

// func (suite *SignupControllerTestSuite) TestSignup_BadRequest() {
// 	request, _ := http.NewRequest(http.MethodPost, "/signup", nil)
// 	responseWriter := httptest.NewRecorder()
// 	suite.router.ServeHTTP(responseWriter, request)

// 	suite.Equal(http.StatusBadRequest, responseWriter.Code)
// 	suite.Contains(responseWriter.Body.String(), "invalid request payload")
// }

// func (suite *SignupControllerTestSuite) TestSignup_InternalServerError() {
// 	newUser := &models.User{
// 		Username: "testuser",
// 		Name:     "Test User",
// 		Email:    "test@example.com",
// 	}
// 	expectedError := &models.ErrorResponse{
// 		Code:    http.StatusInternalServerError,
// 		Message: "internal server error",
// 	}

// 	suite.mockSignupUsecase.On("CreateUser", mock.Anything, newUser).Return(expectedError).Once()

// 	body, _ := json.Marshal(newUser)
// 	request, _ := http.NewRequest(http.MethodPost, "/signup", bytes.NewBuffer(body))
// 	request.Header.Set("Content-Type", "application/json")

// 	responseWriter := httptest.NewRecorder()
// 	suite.router.ServeHTTP(responseWriter, request)

// 	suite.Equal(expectedError.Code, responseWriter.Code)
// 	suite.Contains(responseWriter.Body.String(), "internal server error")
// }

// func (suite *SignupControllerTestSuite) TestConfirmRegistration_Success() {
// 	requestBody := dtos.SetUpPasswordRequest{
// 		Password: "test_password",
// 	}

// 	suite.mockPasswordUsecase.On("SetNewUserPassword", mock.Anything, "123", requestBody.Password).Return(nil).Once()

// 	body, _ := json.Marshal(requestBody)
// 	request, _ := http.NewRequest(http.MethodPost, "/confirm/123", bytes.NewBuffer(body))
// 	request.Header.Set("Content-Type", "application/json")

// 	responseWriter := httptest.NewRecorder()
// 	suite.router.ServeHTTP(responseWriter, request)

// 	suite.Equal(http.StatusOK, responseWriter.Code)
// 	suite.Contains(responseWriter.Body.String(), "registration successful, proceed to login")
// }

// func (suite *SignupControllerTestSuite) TestConfirmRegistration_BadRequest() {
// 	request, _ := http.NewRequest(http.MethodPost, "/confirm/123", nil)
// 	responseWriter := httptest.NewRecorder()
// 	suite.router.ServeHTTP(responseWriter, request)

// 	suite.Equal(http.StatusBadRequest, responseWriter.Code)
// 	suite.Contains(responseWriter.Body.String(), "invalid request payload")
// }

// func (suite *SignupControllerTestSuite) TestConfirmRegistration_InternalServerError() {
// 	requestBody := dtos.SetUpPasswordRequest{
// 		Password: "test_password",
// 	}
// 	expectedError := &models.ErrorResponse{
// 		Code:    http.StatusInternalServerError,
// 		Message: "internal server error",
// 	}

// 	suite.mockPasswordUsecase.On("SetNewUserPassword", mock.Anything, "123", requestBody.Password).Return(expectedError).Once()

// 	body, _ := json.Marshal(requestBody)
// 	request, _ := http.NewRequest(http.MethodPost, "/confirm/123", bytes.NewBuffer(body))
// 	request.Header.Set("Content-Type", "application/json")

// 	responseWriter := httptest.NewRecorder()
// 	suite.router.ServeHTTP(responseWriter, request)

// 	suite.Equal(expectedError.Code, responseWriter.Code)
// 	suite.Contains(responseWriter.Body.String(), "internal server error")
// }

// func TestSignupControllerTestSuite(t *testing.T) {
// 	suite.Run(t, new(SignupControllerTestSuite))
// }
