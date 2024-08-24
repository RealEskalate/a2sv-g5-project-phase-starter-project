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

// // Test Suite
// type ForgotPasswordControllerTestSuite struct {
// 	suite.Suite
// 	mockUsecase *mocks.PasswordUsecase
// 	controller  *controllers.ForgotPasswordController
// 	router      *gin.Engine
// }

// func (suite *ForgotPasswordControllerTestSuite) SetupSuite() {
// 	suite.mockUsecase = new(mocks.PasswordUsecase)
// 	suite.controller = controllers.NewForgotPasswordController(suite.mockUsecase)
// 	suite.router = gin.Default()

// 	// Define the routes
// 	suite.router.POST("/forgot-password", suite.controller.ForgotPasswordRequest)
// 	suite.router.POST("/reset-password/:id", suite.controller.SetNewPassword)
// }

// func (suite *ForgotPasswordControllerTestSuite) TearDownSuite() {
// 	suite.mockUsecase.AssertExpectations(suite.T())
// }

// func (suite *ForgotPasswordControllerTestSuite) TestForgotPasswordRequest_Success() {
// 	requestPayload := dtos.PasswordResetRequest{Email: "test@example.com"}
// 	resetURL := "http://example.com/reset-password/12345"

// 	suite.mockUsecase.On("GenerateResetURL", mock.Anything, requestPayload.Email).Return(resetURL, nil).Once()
// 	suite.mockUsecase.On("SendResetEmail", mock.Anything, requestPayload.Email, resetURL).Return(nil).Once()

// 	body, _ := json.Marshal(requestPayload)
// 	request, _ := http.NewRequest(http.MethodPost, "/forgot-password", bytes.NewBuffer(body))
// 	request.Header.Set("Content-Type", "application/json")
// 	responseWriter := httptest.NewRecorder()

// 	suite.router.ServeHTTP(responseWriter, request)

// 	suite.Equal(http.StatusOK, responseWriter.Code)
// 	suite.Contains(responseWriter.Body.String(), "confirmation email sent")
// }

// func (suite *ForgotPasswordControllerTestSuite) TestForgotPasswordRequest_BadRequest() {
// 	request, _ := http.NewRequest(http.MethodPost, "/forgot-password", nil)
// 	request.Header.Set("Content-Type", "application/json")
// 	responseWriter := httptest.NewRecorder()

// 	suite.router.ServeHTTP(responseWriter, request)

// 	suite.Equal(http.StatusBadRequest, responseWriter.Code)
// 	suite.Contains(responseWriter.Body.String(), "invalid payload")
// }

// func (suite *ForgotPasswordControllerTestSuite) TestForgotPasswordRequest_InternalServerError_GenerateURL() {
// 	requestPayload := dtos.PasswordResetRequest{Email: "test@example.com"}
// 	expectedError := &models.ErrorResponse{
// 		Code:    http.StatusInternalServerError,
// 		Message: "failed to generate reset URL",
// 	}

// 	suite.mockUsecase.On("GenerateResetURL", mock.Anything, requestPayload.Email).Return("", expectedError).Once()

// 	body, _ := json.Marshal(requestPayload)
// 	request, _ := http.NewRequest(http.MethodPost, "/forgot-password", bytes.NewBuffer(body))
// 	request.Header.Set("Content-Type", "application/json")
// 	responseWriter := httptest.NewRecorder()

// 	suite.router.ServeHTTP(responseWriter, request)

// 	suite.Equal(expectedError.Code, responseWriter.Code)
// 	suite.Contains(responseWriter.Body.String(), "failed to generate reset URL")
// }

// func (suite *ForgotPasswordControllerTestSuite) TestForgotPasswordRequest_InternalServerError_SendEmail() {
// 	requestPayload := dtos.PasswordResetRequest{Email: "test@example.com"}
// 	resetURL := "http://example.com/reset-password/12345"
// 	expectedError := &models.ErrorResponse{
// 		Code:    http.StatusInternalServerError,
// 		Message: "failed to send email",
// 	}

// 	suite.mockUsecase.On("GenerateResetURL", mock.Anything, requestPayload.Email).Return(resetURL, nil).Once()
// 	suite.mockUsecase.On("SendResetEmail", mock.Anything, requestPayload.Email, resetURL).Return(expectedError).Once()

// 	body, _ := json.Marshal(requestPayload)
// 	request, _ := http.NewRequest(http.MethodPost, "/forgot-password", bytes.NewBuffer(body))
// 	request.Header.Set("Content-Type", "application/json")
// 	responseWriter := httptest.NewRecorder()

// 	suite.router.ServeHTTP(responseWriter, request)

// 	suite.Equal(expectedError.Code, responseWriter.Code)
// 	suite.Contains(responseWriter.Body.String(), "failed to send email")
// }

// func (suite *ForgotPasswordControllerTestSuite) TestSetNewPassword_Success() {
// 	shortURLCode := "12345"
// 	newPassword := "newSecurePassword"
// 	requestPayload := dtos.SetUpPasswordRequest{Password: newPassword}

// 	suite.mockUsecase.On("SetUpdateUserPassword", mock.Anything, shortURLCode, newPassword).Return(nil).Once()

// 	body, _ := json.Marshal(requestPayload)
// 	request, _ := http.NewRequest(http.MethodPost, "/reset-password/"+shortURLCode, bytes.NewBuffer(body))
// 	request.Header.Set("Content-Type", "application/json")
// 	responseWriter := httptest.NewRecorder()

// 	suite.router.ServeHTTP(responseWriter, request)

// 	suite.Equal(http.StatusOK, responseWriter.Code)
// 	suite.Contains(responseWriter.Body.String(), "password reset, login again")
// }

// func (suite *ForgotPasswordControllerTestSuite) TestSetNewPassword_BadRequest() {
// 	request, _ := http.NewRequest(http.MethodPost, "/reset-password/12345", nil)
// 	request.Header.Set("Content-Type", "application/json")
// 	responseWriter := httptest.NewRecorder()

// 	suite.router.ServeHTTP(responseWriter, request)

// 	suite.Equal(http.StatusBadRequest, responseWriter.Code)
// 	suite.Contains(responseWriter.Body.String(), "invalid request")
// }

// func (suite *ForgotPasswordControllerTestSuite) TestSetNewPassword_InternalServerError() {
// 	shortURLCode := "12345"
// 	newPassword := "newSecurePassword"
// 	requestPayload := dtos.SetUpPasswordRequest{Password: newPassword}
// 	expectedError := &models.ErrorResponse{
// 		Code:    http.StatusInternalServerError,
// 		Message: "failed to reset password",
// 	}

// 	suite.mockUsecase.On("SetUpdateUserPassword", mock.Anything, shortURLCode, newPassword).Return(expectedError).Once()

// 	body, _ := json.Marshal(requestPayload)
// 	request, _ := http.NewRequest(http.MethodPost, "/reset-password/"+shortURLCode, bytes.NewBuffer(body))
// 	request.Header.Set("Content-Type", "application/json")
// 	responseWriter := httptest.NewRecorder()

// 	suite.router.ServeHTTP(responseWriter, request)

// 	suite.Equal(expectedError.Code, responseWriter.Code)
// 	suite.Contains(responseWriter.Body.String(), "failed to reset password")
// }

// func TestForgotPasswordControllerTestSuite(t *testing.T) {
// 	suite.Run(t, new(ForgotPasswordControllerTestSuite))
// }
