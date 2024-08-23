package controller

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	controllers "github.com/aait.backend.g5.main/backend/Delivery/Controllers"
	mocks "github.com/aait.backend.g5.main/backend/Mocks/usecase_mock"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

// Test Suite
type RefreshControllerTestSuite struct {
	suite.Suite
	mockUsecase      *mocks.RefreshUsecase
	mockJwtService   *mocks.JwtService
	mockOauthService *mocks.MockOAuthService
	controller       *controllers.RefreshController
	router           *gin.Engine
}

func (suite *RefreshControllerTestSuite) SetupSuite() {
	suite.mockUsecase = new(mocks.RefreshUsecase)
	suite.mockJwtService = new(mocks.JwtService)
	suite.mockOauthService = new(mocks.MockOAuthService)
	suite.controller = controllers.NewRefreshController(suite.mockUsecase, suite.mockJwtService, suite.mockOauthService)
	suite.router = gin.Default()

	// Define the route
	suite.router.POST("/refresh", suite.controller.Refresh)
}

func (suite *RefreshControllerTestSuite) TearDownSuite() {
	suite.mockUsecase.AssertExpectations(suite.T())
	suite.mockJwtService.AssertExpectations(suite.T())
}

func (suite *RefreshControllerTestSuite) TestRefresh_Success() {
	userId := "123"
	authHeader := "Bearer some-valid-refresh-token"
	parsedAuthParts := []string{"Bearer", "some-valid-refresh-token"}
	expectedAccessToken := "new-access-token"

	// Mocking JwtService.ValidateAuthHeader
	suite.mockJwtService.On("ValidateAuthHeader", authHeader).Return(parsedAuthParts, nil).Once()

	request, _ := http.NewRequest(http.MethodPost, "/refresh", nil)
	request.Header.Set("Authorization", authHeader)

	// Use Gin's context and manually set the user ID
	responseWriter := httptest.NewRecorder()
	suite.mockUsecase.On("RefreshToken", mock.Anything, "", parsedAuthParts[1]).Return(expectedAccessToken, nil).Once()

	// Set user ID in context
	ctx, _ := gin.CreateTestContext(responseWriter)
	ctx.Request = request
	ctx.Set("id", userId)

	// suite.controller.Refresh(ctx)

	suite.router.ServeHTTP(ctx.Writer, request)

	suite.Equal(http.StatusOK, ctx.Writer.Status())
	suite.Contains(responseWriter.Body.String(), expectedAccessToken)
}

func (suite *RefreshControllerTestSuite) TestRefresh_Unauthorized_InvalidAuthHeader() {
	authHeader := "invalid-auth-header"
	expectedError := errors.New("invalid authorization header")

	// Mocking JwtService.ValidateAuthHeader
	suite.mockJwtService.On("ValidateAuthHeader", authHeader).Return(nil, expectedError).Once()

	request, _ := http.NewRequest(http.MethodPost, "/refresh", nil)
	request.Header.Set("Authorization", authHeader)
	// request.Header.Set("Content-Type", "application/json")

	responseWriter := httptest.NewRecorder()

	// Set user ID in context
	ctx, _ := gin.CreateTestContext(responseWriter)
	ctx.Request = request

	suite.router.ServeHTTP(ctx.Writer, request)

	suite.Equal(http.StatusUnauthorized, ctx.Writer.Status())
	suite.Contains(responseWriter.Body.String(), expectedError.Error())
}

func TestRefreshControllerTestSuite(t *testing.T) {
	suite.Run(t, new(RefreshControllerTestSuite))
}
