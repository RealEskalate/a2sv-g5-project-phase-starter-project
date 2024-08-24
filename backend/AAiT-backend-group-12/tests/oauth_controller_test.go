package tests

import (
	"blog_api/delivery/controllers"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/markbates/goth"
	"github.com/stretchr/testify/suite"
)

type OAuthControllerTestSuite struct {
	suite.Suite
	controller *controllers.OAuthController
}

func (suite *OAuthControllerTestSuite) SetupTest() {
	// Mock CompleteUserAuth function
	mockCompleteUserAuth := func(res http.ResponseWriter, req *http.Request) (goth.User, error) {
		return goth.User{Email: "test@example.com", Name: "Test User"}, nil
	}

	// Mock BeginAuthHandler function
	mockBeginAuthHandler := func(res http.ResponseWriter, req *http.Request) {
		http.Redirect(res, req, "/redirect-url", http.StatusTemporaryRedirect)
	}

	// Initialize the controller with mocks
	suite.controller = controllers.NewOAuthController(
		mockCompleteUserAuth,
		mockBeginAuthHandler,
	)
}

func (suite *OAuthControllerTestSuite) TestGoogleAuthInit() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/auth/google/start", nil)

	// Call the handler
	suite.controller.GoogleAuthInit(c)

	// Check if the response code is 307 (Temporary Redirect)
	suite.Equal(http.StatusTemporaryRedirect, w.Code)

	// Check if the Location header contains "redirect-url"
	suite.Contains(w.Header().Get("Location"), "/redirect-url")
}

func (suite *OAuthControllerTestSuite) TestOAuthCallback() {
	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)
	c.Request, _ = http.NewRequest("GET", "/auth/google/callback", nil)

	// Mock the CompleteUserAuth function to return a test user
	suite.controller.CompleteUserAuth = func(res http.ResponseWriter, req *http.Request) (goth.User, error) {
		return goth.User{
			Email: "test@example.com",
			Name:  "Test User",
		}, nil
	}

	// Call the handler
	suite.controller.OAuthCallback(c)

	// Check if the response code is 202 (Accepted)
	suite.Equal(http.StatusAccepted, w.Code)
}

func TestOAuthControllerTestSuite(t *testing.T) {
	suite.Run(t, new(OAuthControllerTestSuite))
}
