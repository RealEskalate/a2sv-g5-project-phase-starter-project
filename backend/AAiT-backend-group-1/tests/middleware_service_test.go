package tests

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/domain"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/infrastructure"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/mocks"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type MiddlewareServiceTestSuite struct {
	suite.Suite
	jwtService   *mocks.JwtService
	cacheService *mocks.CacheService
	middleware   domain.MiddlewareService
}

func (suite *MiddlewareServiceTestSuite) SetupTest() {
	suite.jwtService = new(mocks.JwtService)
	suite.cacheService = new(mocks.CacheService)
	suite.middleware = infrastructure.NewMiddlewareService(suite.jwtService, suite.cacheService)

	gin.SetMode(gin.TestMode)
}

func (suite *MiddlewareServiceTestSuite) TestAuthenticate() {
	claims := jwt.MapClaims{
		"user_id":  "user",
		"username": "user",
		"role":     "user",
		"exp":      1000,
		"iss":      "user",
	}

	parsedToken := &jwt.Token{Claims: claims}
	suite.jwtService.On("ValidateAccessToken", "token").Return(parsedToken, nil)
	suite.cacheService.On("Get", "token").Return("", nil)

	handler := suite.middleware.Authenticate()
	c := &gin.Context{}
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Add("Authorization", "Bearer token")

	c.Request = req
	handler(c)

	suite.jwtService.AssertExpectations(suite.T())
	suite.cacheService.AssertExpectations(suite.T())
}

func (suite *MiddlewareServiceTestSuite) TestAuthenticate_Failure() {
	suite.jwtService.On("ValidateAccessToken", "invalid_token").Return(nil, domain.CustomError{Message: "invalid token", Code: http.StatusUnauthorized})
	suite.cacheService.On("Get", "invalid_token").Return("", nil)
	handler := suite.middleware.Authenticate()
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Add("Authorization", "Bearer invalid_token")

	c.Request = req
	handler(c)

	assert.Equal(suite.T(), http.StatusUnauthorized, recorder.Code)

	suite.jwtService.AssertExpectations(suite.T())
	suite.cacheService.AssertExpectations(suite.T())
}

func (suite *MiddlewareServiceTestSuite) TestAuthenticate_Failure_Cache() {
	suite.cacheService.On("Get", "invalid_token").Return("", fmt.Errorf("error"))
	handler := suite.middleware.Authenticate()
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Add("Authorization", "Bearer invalid_token")

	c.Request = req
	handler(c)

	assert.Equal(suite.T(), http.StatusInternalServerError, recorder.Code)

	suite.jwtService.AssertNotCalled(suite.T(), "ValidateAccessToken", "invalid_token")
	suite.cacheService.AssertExpectations(suite.T())
}

func (suite *MiddlewareServiceTestSuite) TestAuthenticate_Failure_BlackListed() {
	suite.cacheService.On("Get", "invalid_token").Return("exists", nil)
	handler := suite.middleware.Authenticate()
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	req := httptest.NewRequest("GET", "/test", nil)
	req.Header.Add("Authorization", "Bearer invalid_token")

	c.Request = req
	handler(c)

	suite.jwtService.AssertNotCalled(suite.T(), "ValidateAccessToken", "invalid_token")
	suite.cacheService.AssertExpectations(suite.T())
}

func (suite *MiddlewareServiceTestSuite) TestAuthorize_Success() {
	claims := jwt.MapClaims{
		"user_id":  "user",
		"username": "user",
		"role":     "admin",
		"exp":      1000,
	}

	parsedToken := &jwt.Token{Claims: claims, Valid: true}
	suite.jwtService.On("ValidateAccessToken", "valid_token").Return(parsedToken, nil)

	handler := suite.middleware.Authorize("admin", "user")
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	req := httptest.NewRequest("GET", "/admin", nil)
	req.Header.Add("Authorization", "Bearer valid_token")

	c.Request = req

	handler(c)

	assert.Equal(suite.T(), http.StatusOK, recorder.Code)

	suite.jwtService.AssertExpectations(suite.T())
}

func (suite *MiddlewareServiceTestSuite) TestAuthorize_InvalidToken() {
	suite.jwtService.On("ValidateAccessToken", "invalid_token").Return(nil, domain.CustomError{Message: "invalid token", Code: http.StatusUnauthorized})

	handler := suite.middleware.Authorize("admin")
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	req := httptest.NewRequest("GET", "/admin", nil)
	req.Header.Add("Authorization", "Bearer invalid_token")

	c.Request = req

	handler(c)

	assert.Equal(suite.T(), http.StatusUnauthorized, recorder.Code)
	assert.JSONEq(suite.T(), `{"error": "Invalid token"}`, recorder.Body.String())

	suite.jwtService.AssertExpectations(suite.T())
}

func (suite *MiddlewareServiceTestSuite) TestAuthorize_ForbiddenRole() {
	claims := jwt.MapClaims{
		"user_id":  "user",
		"username": "user",
		"role":     "user",
		"exp":      1000,
	}

	parsedToken := &jwt.Token{Claims: claims, Valid: true}
	suite.jwtService.On("ValidateAccessToken", "valid_token").Return(parsedToken, nil)

	handler := suite.middleware.Authorize("admin")
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	req := httptest.NewRequest("GET", "/admin", nil)
	req.Header.Add("Authorization", "Bearer valid_token")

	c.Request = req

	handler(c)

	assert.Equal(suite.T(), http.StatusForbidden, recorder.Code)
	assert.JSONEq(suite.T(), `{"error": "You are not authorized for this action"}`, recorder.Body.String())

	suite.jwtService.AssertExpectations(suite.T())
}

func (suite *MiddlewareServiceTestSuite) TestAuthorize_MissingRoleClaim() {
	claims := jwt.MapClaims{
		"user_id":  "user",
		"username": "user",
		"exp":      1000,
	}

	parsedToken := &jwt.Token{Claims: claims, Valid: true}
	suite.jwtService.On("ValidateAccessToken", "valid_token").Return(parsedToken, nil)

	handler := suite.middleware.Authorize("admin")
	recorder := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(recorder)
	req := httptest.NewRequest("GET", "/admin", nil)
	req.Header.Add("Authorization", "Bearer valid_token")

	c.Request = req

	handler(c)

	assert.Equal(suite.T(), http.StatusForbidden, recorder.Code)
	assert.JSONEq(suite.T(), `{"error": "You are not authorized for this action"}`, recorder.Body.String())

	suite.jwtService.AssertExpectations(suite.T())
}

func TestMiddlewareServiceTestSuite(t *testing.T) {
	suite.Run(t, new(MiddlewareServiceTestSuite))
}
