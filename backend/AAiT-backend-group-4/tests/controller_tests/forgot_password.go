package controller_tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"aait-backend-group4/Delivery/controllers"
	"aait-backend-group4/Domain"
	"aait-backend-group4/Mocks/usecase_mocks"
	"aait-backend-group4/Mocks/infrastructure_mocks"
	// "aait-backend-group4/Mocks/repositories_mocks"
	

)



type ForgotPasswordControllerTestSuite struct {
	suite.Suite
	Router  *gin.Engine
	Controller *controllers.ForgotPasswordController
}

func (suite *ForgotPasswordControllerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	router := gin.Default()

	mockOtpService := new(mocks.MockOtpService)
	mockUserRepository := new(MockUserRepository)
	mockPasswordService := new(MockPasswordService)

	usecase := mocks.NewPasswordUsecase(mockUserRepository, mockOtpService, time.Second*10, mockPasswordService)

	suite.Controller = &controllers.ForgotPasswordController{
		OtpService:            mockOtpService,
		ForgotPasswordUsecase: usecase,
		Env:                   &Bootstrap.Env{EmailApiKey: "mock-api-key"},
	}

	routes.NewForgotPasswordRoute(&Bootstrap.Env{EmailApiKey: "mock-api-key"}, time.Second*10, nil, router.Group("/"))
	suite.Router = router
}

func (suite *ForgotPasswordControllerTestSuite) TestForgotPassword() {
	tests := []struct {
		name         string
		requestBody  map[string]string
		mockBehavior func(mockOtpService *mocks.MockOtpService)
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Valid email",
			requestBody:  map[string]string{"email": "user@example.com"},
			mockBehavior: func(m *mocks.MockOtpService) { m.On("SendPasswordResetEmail", "user@example.com", "Reset Password", "mock-api-key").Return(nil) },
			expectedCode: http.StatusOK,
			expectedBody: `{"status":"Not verified","message":"Password reset email sent successfully"}`,
		},
		{
			name:         "Invalid email",
			requestBody:  map[string]string{"email": "invalid-email"},
			mockBehavior: func(m *mocks.MockOtpService) { m.On("SendPasswordResetEmail", "invalid-email", "Reset Password", "mock-api-key").Return(assert.AnError) },
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"error sending password reset email"}`, // Adjust based on actual error
		},
		{
			name:         "Invalid request body",
			requestBody:  "invalid",
			mockBehavior: func(m *mocks.MockOtpService) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"json: cannot unmarshal string into Go value of type map[string]string"}`, // Adjust based on actual error
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			tt.mockBehavior(suite.Controller.OtpService.(*mocks.MockOtpService))
			body, _ := json.Marshal(tt.requestBody)
			req, _ := http.NewRequest(http.MethodPost, "/forgot-password", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")

			w := httptest.NewRecorder()
			suite.Router.ServeHTTP(w, req)

			assert.Equal(suite.T(), tt.expectedCode, w.Code)
			assert.JSONEq(suite.T(), tt.expectedBody, w.Body.String())
		})
	}
}

func (suite *ForgotPasswordControllerTestSuite) TestVerifyForgotPassword() {
	tests := []struct {
		name         string
		email        string
		requestBody  domain.ForgotPasswordRequest
		mockBehavior func(mockUsecase *mocks.ForgotPasswordUsecase)
		expectedCode int
		expectedBody string
	}{
		{
			name:         "Valid password change",
			email:        "user@example.com",
			requestBody:  domain.ForgotPasswordRequest{New_Password: "newpassword", Confirmation: "newpassword"},
			mockBehavior: func(m *mocks.ForgotPasswordUsecase) { m.On("VerifyChangePassword", mock.Anything, "user@example.com", domain.ForgotPasswordRequest{New_Password: "newpassword", Confirmation: "newpassword"}).Return(domain.ForgotPasswordResponse{Message: "Password updated successfully"}, nil) },
			expectedCode: http.StatusOK,
			expectedBody: `{"message":"Password updated successfully"}`,
		},
		{
			name:         "Passwords do not match",
			email:        "user@example.com",
			requestBody:  domain.ForgotPasswordRequest{New_Password: "newpassword", Confirmation: "differentpassword"},
			mockBehavior: func(m *mocks.ForgotPasswordUsecase) { m.On("VerifyChangePassword", mock.Anything, "user@example.com", domain.ForgotPasswordRequest{New_Password: "newpassword", Confirmation: "differentpassword"}).Return(domain.ForgotPasswordResponse{}, assert.AnError) },
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"passwords do not match"}`, // Adjust based on actual error
		},
		{
			name:         "Invalid request body",
			email:        "user@example.com",
			requestBody:  "invalid",
			mockBehavior: func(m *mocks.ForgotPasswordUsecase) {},
			expectedCode: http.StatusBadRequest,
			expectedBody: `{"error":"json: cannot unmarshal string into Go value of type domain.ForgotPasswordRequest"}`, // Adjust based on actual error
		},
		{
			name:         "Email not in context",
			email:        "",
			requestBody:  domain.ForgotPasswordRequest{New_Password: "newpassword", Confirmation: "newpassword"},
			mockBehavior: func(m *mocks.ForgotPasswordUsecase) {},
			expectedCode: http.StatusUnauthorized,
			expectedBody: `{"error":"Email not found in context"}`,
		},
	}

	for _, tt := range tests {
		suite.Run(tt.name, func() {
			req, _ := http.NewRequest(http.MethodPost, "/submit-change-password", nil)
			req.Header.Set("Content-Type", "application/json")
			req = req.WithContext(context.WithValue(req.Context(), "email", tt.email))

			body, _ := json.Marshal(tt.requestBody)
			req.Body = ioutil.NopCloser(bytes.NewBuffer(body))

			w := httptest.NewRecorder()
			suite.Router.ServeHTTP(w, req)

			assert.Equal(suite.T(), tt.expectedCode, w.Code)
			assert.JSONEq(suite.T(), tt.expectedBody, w.Body.String())
		})
	}
}

func (suite *ForgotPasswordControllerTestSuite) TestServePage() {
	req, _ := http.NewRequest(http.MethodGet, "/forgot-password", nil)
	w := httptest.NewRecorder()
	suite.Router.ServeHTTP(w, req)

	assert.Equal(suite.T(), http.StatusOK, w.Code)
	assert.Contains(suite.T(), w.Body.String(), "password_reset.html") // Adjust based on actual content
}

func TestForgotPasswordController(t *testing.T) {
	suite.Run(t, new(ForgotPasswordControllerTestSuite))
}
