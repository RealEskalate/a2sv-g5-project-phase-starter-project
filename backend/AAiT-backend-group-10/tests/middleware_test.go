package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"aait.backend.g10/domain"
	"aait.backend.g10/infrastructures"
	"aait.backend.g10/tests/mocks"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MiddlewareTestSuite struct {
	suite.Suite
	mockJwtService *mocks.IJwtService
	router         *gin.Engine
}

func (s *MiddlewareTestSuite) SetupTest() {
	s.mockJwtService = new(mocks.IJwtService)
	s.router = gin.Default()
	s.router.Use(gin.Logger())

	// Setup routes for testing
	authGroup := s.router.Group("/")
	authGroup.Use(infrastructures.AuthMiddleware(s.mockJwtService))
	authGroup.GET("/protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Success"})
	})

	adminGroup := s.router.Group("/admin")
	adminGroup.Use(infrastructures.AuthMiddleware(s.mockJwtService), infrastructures.AdminMiddleWare())
	adminGroup.GET("/admin-protected", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Admin Success"})
	})
}

func (s *MiddlewareTestSuite) TestAuthMiddleware_Success() {
	token := &jwt.Token{Valid: true}
	claims := jwt.MapClaims{
		"id":       uuid.New().String(),
		"is_admin": false,
	}

	s.mockJwtService.On("CheckToken", mock.Anything).Return(token, nil)
	s.mockJwtService.On("FindClaim", token).Return(claims, true)

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer valid.token")
	w := httptest.NewRecorder()

	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusOK, w.Code)
	assert.JSONEq(s.T(), `{"message": "Success"}`, w.Body.String())
}

func (s *MiddlewareTestSuite) TestAuthMiddleware_Failure_NoToken() {
	req, _ := http.NewRequest("GET", "/protected", nil)
	w := httptest.NewRecorder()

	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusUnauthorized, w.Code)
	assert.JSONEq(s.T(), `{"Error": "Authorization header is required"}`, w.Body.String())
}

func (s *MiddlewareTestSuite) TestAuthMiddleware_Failure_InvalidTokenFormat() {
	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "InvalidTokenFormat")
	w := httptest.NewRecorder()

	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusUnauthorized, w.Code)
	assert.JSONEq(s.T(), `{"message": "Invalid Authorization header"}`, w.Body.String())
}

func (s *MiddlewareTestSuite) TestAuthMiddleware_Failure_ExpiredOrInvalidToken() {
	s.mockJwtService.On("CheckToken", mock.Anything).Return(nil, &domain.CustomError{Message: "Invalid or expired token"})

	req, _ := http.NewRequest("GET", "/protected", nil)
	req.Header.Set("Authorization", "Bearer invalid.token")
	w := httptest.NewRecorder()

	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusUnauthorized, w.Code)
	assert.JSONEq(s.T(), `{"error": "Invalid or expired token"}`, w.Body.String())
}

func (s *MiddlewareTestSuite) TestAdminMiddleware_Success() {
	token := &jwt.Token{Valid: true}
	claims := jwt.MapClaims{
		"id":       uuid.New().String(),
		"is_admin": true,
	}

	s.mockJwtService.On("CheckToken", mock.Anything).Return(token, nil)
	s.mockJwtService.On("FindClaim", token).Return(claims, true)

	req, _ := http.NewRequest("GET", "/admin/admin-protected", nil)
	req.Header.Set("Authorization", "Bearer valid.token")
	w := httptest.NewRecorder()

	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusOK, w.Code)
	assert.JSONEq(s.T(), `{"message": "Admin Success"}`, w.Body.String())
}

func (s *MiddlewareTestSuite) TestAdminMiddleware_Failure_NotAdmin() {
	token := &jwt.Token{Valid: true}
	claims := jwt.MapClaims{
		"id":       uuid.New().String(),
		"is_admin": false,
	}

	s.mockJwtService.On("CheckToken", mock.Anything).Return(token, nil)
	s.mockJwtService.On("FindClaim", token).Return(claims, true)

	req, _ := http.NewRequest("GET", "/admin/admin-protected", nil)
	req.Header.Set("Authorization", "Bearer valid.token")
	w := httptest.NewRecorder()

	s.router.ServeHTTP(w, req)

	assert.Equal(s.T(), http.StatusForbidden, w.Code)
	assert.JSONEq(s.T(), `{"message": "Sorry, you must be an admin"}`, w.Body.String())
}

func TestMiddlewareTestSuite(t *testing.T) {
	suite.Run(t, new(MiddlewareTestSuite))
}
