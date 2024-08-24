package tests

import (
	"blog_api/domain"
	"blog_api/infrastructure/middleware"
	"blog_api/mocks"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/suite"
)

const (
	VALID_ROLE   = "valid_role"
	INVALID_ROLE = "invalid_role"
)

type AuthMiddlewareSuite struct {
	suite.Suite
	jwtService      *mocks.JWTServiceInterface
	cacheRepository *mocks.CacheRepositoryInterface
	testingServer   *httptest.Server
}

func (suite *AuthMiddlewareSuite) SetupSuite() {
	suite.jwtService = &mocks.JWTServiceInterface{}
	suite.cacheRepository = &mocks.CacheRepositoryInterface{}

	router := gin.Default()
	router.GET("/", middleware.AuthMiddlewareWithRoles(suite.jwtService, suite.cacheRepository, VALID_ROLE), func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, "")
	})

	suite.testingServer = httptest.NewServer(router)
}

func (suite *AuthMiddlewareSuite) TestAuthMiddleware_Positive() {
	rawToken := "lksadfjlsadfj.elrqlkrjqwaklsdfm.lsiasdflkxzjcvlk"
	token := &jwt.Token{}

	suite.cacheRepository.On("IsCached", rawToken).Return(false).Once()
	suite.jwtService.On("ValidateAndParseToken", rawToken).Return(token, nil).Once()
	suite.jwtService.On("GetTokenType", token).Return("accessToken", nil).Once()
	suite.jwtService.On("GetExpiryDate", token).Return(time.Now().Add(time.Minute), nil).Once()
	suite.jwtService.On("GetRole", token).Return(VALID_ROLE, nil).Once()
	suite.jwtService.On("GetUsername", token).Return("username", nil).Once()

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, suite.testingServer.URL+"/", nil)
	suite.NoError(err, "no error during request creation")
	request.Header.Add("Authorization", "bearer "+rawToken)
	response, err := client.Do(request)
	suite.NoError(err, "no error during request")
	if response != nil {
		defer response.Body.Close()
	}

	suite.cacheRepository.AssertExpectations(suite.T())
	suite.jwtService.AssertExpectations(suite.T())
	suite.Equal(http.StatusOK, response.StatusCode)
}

func (suite *AuthMiddlewareSuite) TestAuthMiddleware_Negative_NoAuthHeader() {
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, suite.testingServer.URL+"/", nil)
	suite.NoError(err, "no error during request creation")
	response, err := client.Do(request)
	suite.NoError(err, "no error during request")
	if response != nil {
		defer response.Body.Close()
	}

	suite.cacheRepository.AssertExpectations(suite.T())
	suite.jwtService.AssertExpectations(suite.T())
	suite.Equal(http.StatusUnauthorized, response.StatusCode)
}

func (suite *AuthMiddlewareSuite) TestAuthMiddleware_Negative_InvalidHeader() {
	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, suite.testingServer.URL+"/", nil)
	suite.NoError(err, "no error during request creation")
	request.Header.Add("Authorization", "bearesdafr "+"asdf")
	response, err := client.Do(request)
	suite.NoError(err, "no error during request")
	if response != nil {
		defer response.Body.Close()
	}

	suite.cacheRepository.AssertExpectations(suite.T())
	suite.jwtService.AssertExpectations(suite.T())
	suite.Equal(http.StatusUnauthorized, response.StatusCode)
}

func (suite *AuthMiddlewareSuite) TestAuthMiddleware_Negative_TokenBlacklisted() {
	rawToken := "lksadfjlsadfj.elrqlkrjqwaklsdfm.lsiasdflkxzjcvlk"

	suite.cacheRepository.On("IsCached", rawToken).Return(true).Once()

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, suite.testingServer.URL+"/", nil)
	suite.NoError(err, "no error during request creation")
	request.Header.Add("Authorization", "bearer "+rawToken)
	response, err := client.Do(request)
	suite.NoError(err, "no error during request")
	if response != nil {
		defer response.Body.Close()
	}

	suite.cacheRepository.AssertExpectations(suite.T())
	suite.jwtService.AssertExpectations(suite.T())
	suite.Equal(http.StatusUnauthorized, response.StatusCode)
}

func (suite *AuthMiddlewareSuite) TestAuthMiddleware_Negative_InvalidToken() {
	rawToken := "lksadfjlsadfj.elrqlkrjqwaklsdfm.lsiasdflkxzjcvlk"
	token := &jwt.Token{}

	suite.cacheRepository.On("IsCached", rawToken).Return(false).Once()
	suite.jwtService.On("ValidateAndParseToken", rawToken).Return(token, fmt.Errorf("")).Once()

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, suite.testingServer.URL+"/", nil)
	suite.NoError(err, "no error during request creation")
	request.Header.Add("Authorization", "bearer "+rawToken)
	response, err := client.Do(request)
	suite.NoError(err, "no error during request")
	if response != nil {
		defer response.Body.Close()
	}

	suite.cacheRepository.AssertExpectations(suite.T())
	suite.jwtService.AssertExpectations(suite.T())
	suite.Equal(http.StatusUnauthorized, response.StatusCode)
}

func (suite *AuthMiddlewareSuite) TestAuthMiddleware_Negative_NoTokenType() {
	rawToken := "lksadfjlsadfj.elrqlkrjqwaklsdfm.lsiasdflkxzjcvlk"
	token := &jwt.Token{}

	sampleErr := domain.NewError("sample error", domain.ERR_UNAUTHORIZED)
	suite.cacheRepository.On("IsCached", rawToken).Return(false).Once()
	suite.jwtService.On("ValidateAndParseToken", rawToken).Return(token, nil).Once()
	suite.jwtService.On("GetTokenType", token).Return("accessToken", sampleErr).Once()

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, suite.testingServer.URL+"/", nil)
	suite.NoError(err, "no error during request creation")
	request.Header.Add("Authorization", "bearer "+rawToken)
	response, err := client.Do(request)
	suite.NoError(err, "no error during request")
	if response != nil {
		defer response.Body.Close()
	}

	suite.cacheRepository.AssertExpectations(suite.T())
	suite.jwtService.AssertExpectations(suite.T())
	suite.Equal(http.StatusUnauthorized, response.StatusCode)
}

func (suite *AuthMiddlewareSuite) TestAuthMiddleware_Negative_InvalidTokenType() {
	rawToken := "lksadfjlsadfj.elrqlkrjqwaklsdfm.lsiasdflkxzjcvlk"
	token := &jwt.Token{}

	suite.cacheRepository.On("IsCached", rawToken).Return(false).Once()
	suite.jwtService.On("ValidateAndParseToken", rawToken).Return(token, nil).Once()
	suite.jwtService.On("GetTokenType", token).Return("INVALID", nil).Once()

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, suite.testingServer.URL+"/", nil)
	suite.NoError(err, "no error during request creation")
	request.Header.Add("Authorization", "bearer "+rawToken)
	response, err := client.Do(request)
	suite.NoError(err, "no error during request")
	if response != nil {
		defer response.Body.Close()
	}

	suite.cacheRepository.AssertExpectations(suite.T())
	suite.jwtService.AssertExpectations(suite.T())
	suite.Equal(http.StatusUnauthorized, response.StatusCode)
}

func (suite *AuthMiddlewareSuite) TestAuthMiddleware_Negative_NoExpiryDate() {
	rawToken := "lksadfjlsadfj.elrqlkrjqwaklsdfm.lsiasdflkxzjcvlk"
	token := &jwt.Token{}

	sampleErr := domain.NewError("sample error", domain.ERR_UNAUTHORIZED)
	suite.cacheRepository.On("IsCached", rawToken).Return(false).Once()
	suite.jwtService.On("ValidateAndParseToken", rawToken).Return(token, nil).Once()
	suite.jwtService.On("GetTokenType", token).Return("accessToken", nil).Once()
	suite.jwtService.On("GetExpiryDate", token).Return(time.Now().Add(time.Minute), sampleErr).Once()

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, suite.testingServer.URL+"/", nil)
	suite.NoError(err, "no error during request creation")
	request.Header.Add("Authorization", "bearer "+rawToken)
	response, err := client.Do(request)
	suite.NoError(err, "no error during request")
	if response != nil {
		defer response.Body.Close()
	}

	suite.cacheRepository.AssertExpectations(suite.T())
	suite.jwtService.AssertExpectations(suite.T())
	suite.Equal(http.StatusUnauthorized, response.StatusCode)
}

func (suite *AuthMiddlewareSuite) TestAuthMiddleware_Negative_ExpiredToken() {
	rawToken := "lksadfjlsadfj.elrqlkrjqwaklsdfm.lsiasdflkxzjcvlk"
	token := &jwt.Token{}

	suite.cacheRepository.On("IsCached", rawToken).Return(false).Once()
	suite.jwtService.On("ValidateAndParseToken", rawToken).Return(token, nil).Once()
	suite.jwtService.On("GetTokenType", token).Return("accessToken", nil).Once()
	suite.jwtService.On("GetExpiryDate", token).Return(time.Now().Add(time.Minute*-1), nil).Once()

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, suite.testingServer.URL+"/", nil)
	suite.NoError(err, "no error during request creation")
	request.Header.Add("Authorization", "bearer "+rawToken)
	response, err := client.Do(request)
	suite.NoError(err, "no error during request")
	if response != nil {
		defer response.Body.Close()
	}

	suite.cacheRepository.AssertExpectations(suite.T())
	suite.jwtService.AssertExpectations(suite.T())
	suite.Equal(http.StatusUnauthorized, response.StatusCode)
}

func (suite *AuthMiddlewareSuite) TestAuthMiddleware_Negative_NoRole() {
	rawToken := "lksadfjlsadfj.elrqlkrjqwaklsdfm.lsiasdflkxzjcvlk"
	token := &jwt.Token{}

	sampleErr := domain.NewError("sample error", domain.ERR_UNAUTHORIZED)
	suite.cacheRepository.On("IsCached", rawToken).Return(false).Once()
	suite.jwtService.On("ValidateAndParseToken", rawToken).Return(token, nil).Once()
	suite.jwtService.On("GetTokenType", token).Return("accessToken", nil).Once()
	suite.jwtService.On("GetExpiryDate", token).Return(time.Now().Add(time.Minute), nil).Once()
	suite.jwtService.On("GetRole", token).Return(VALID_ROLE, sampleErr).Once()

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, suite.testingServer.URL+"/", nil)
	suite.NoError(err, "no error during request creation")
	request.Header.Add("Authorization", "bearer "+rawToken)
	response, err := client.Do(request)
	suite.NoError(err, "no error during request")
	if response != nil {
		defer response.Body.Close()
	}

	suite.cacheRepository.AssertExpectations(suite.T())
	suite.jwtService.AssertExpectations(suite.T())
	suite.Equal(http.StatusUnauthorized, response.StatusCode)
}

func (suite *AuthMiddlewareSuite) TestAuthMiddleware_Negative_InvalidRole() {
	rawToken := "lksadfjlsadfj.elrqlkrjqwaklsdfm.lsiasdflkxzjcvlk"
	token := &jwt.Token{}

	suite.cacheRepository.On("IsCached", rawToken).Return(false).Once()
	suite.jwtService.On("ValidateAndParseToken", rawToken).Return(token, nil).Once()
	suite.jwtService.On("GetTokenType", token).Return("accessToken", nil).Once()
	suite.jwtService.On("GetExpiryDate", token).Return(time.Now().Add(time.Minute), nil).Once()
	suite.jwtService.On("GetRole", token).Return(INVALID_ROLE, nil).Once()
	suite.jwtService.On("GetUsername", token).Return("username", nil).Once()

	client := &http.Client{}
	request, err := http.NewRequest(http.MethodGet, suite.testingServer.URL+"/", nil)
	suite.NoError(err, "no error during request creation")
	request.Header.Add("Authorization", "bearer "+rawToken)
	response, err := client.Do(request)
	suite.NoError(err, "no error during request")
	if response != nil {
		defer response.Body.Close()
	}

	suite.cacheRepository.AssertExpectations(suite.T())
	suite.jwtService.AssertExpectations(suite.T())
	suite.Equal(http.StatusForbidden, response.StatusCode)
}

func (suite *AuthMiddlewareSuite) TearDownSuite() {
	suite.testingServer.Close()
}

func TestMiddleware(t *testing.T) {
	suite.Run(t, new(AuthMiddlewareSuite))
}
