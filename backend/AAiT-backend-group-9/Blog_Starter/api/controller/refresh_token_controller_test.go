package controller

import (
    "Blog_Starter/domain"
    "Blog_Starter/domain/mocks"
    "bytes"
    "encoding/json"
    "errors"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    "github.com/stretchr/testify/assert"
    "github.com/stretchr/testify/mock"
    "github.com/stretchr/testify/suite"
)

type RefreshTokenControllerTestSuite struct {
    suite.Suite
    refreshTokenController *RefreshTokenController
    mockRefreshTokenUsecase *mocks.RefreshTokenUsecase
}

func (suite *RefreshTokenControllerTestSuite) SetupTest() {
    suite.mockRefreshTokenUsecase = new(mocks.RefreshTokenUsecase)
    suite.refreshTokenController = NewRefreshTokenController(suite.mockRefreshTokenUsecase)
}

func (suite *RefreshTokenControllerTestSuite) TestRefreshToken_Success() {
    userID := "test_user_id"
    refreshTokenRequest := domain.RefreshTokenRequest{
        RefreshToken: "valid_refresh_token",
    }

    refreshTokenResponse := &domain.RefreshTokenResponse{
        AccessToken:  "new_access_token",
        RefreshToken: "new_refresh_token",
    }

    suite.mockRefreshTokenUsecase.On("CheckRefreshToken", mock.Anything, userID, refreshTokenRequest.RefreshToken).Return(nil)
    suite.mockRefreshTokenUsecase.On("UpdateTokens", mock.Anything, userID).Return(refreshTokenResponse, nil)

    body, _ := json.Marshal(refreshTokenRequest)
    req, err := http.NewRequest(http.MethodPost, "/refresh-token", bytes.NewBuffer(body))
    assert.NoError(suite.T(), err)

    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    router := gin.Default()
    router.POST("/refresh-token", func(c *gin.Context) {
        c.Set("userID", userID)
        suite.refreshTokenController.RefreshToken(c)
    })
    router.ServeHTTP(rr, req)

    assert.Equal(suite.T(), http.StatusOK, rr.Code)
    assert.JSONEq(suite.T(), `{"accessToken":"new_access_token","refreshToken":"new_refresh_token"}`, rr.Body.String())
}

func (suite *RefreshTokenControllerTestSuite) TestRefreshToken_InvalidRequest() {
    req, err := http.NewRequest(http.MethodPost, "/refresh-token", bytes.NewBuffer([]byte(`{}`)))
    assert.NoError(suite.T(), err)

    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    router := gin.Default()
    router.POST("/refresh-token", suite.refreshTokenController.RefreshToken)
    router.ServeHTTP(rr, req)

    assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
    assert.Contains(suite.T(), rr.Body.String(), "error")
}

func (suite *RefreshTokenControllerTestSuite) TestRefreshToken_CheckRefreshTokenError() {
    userID := "test_user_id"
    refreshTokenRequest := domain.RefreshTokenRequest{
        RefreshToken: "invalid_refresh_token",
    }

    suite.mockRefreshTokenUsecase.On("CheckRefreshToken", mock.Anything, userID, refreshTokenRequest.RefreshToken).Return(errors.New("invalid refresh token"))

    body, _ := json.Marshal(refreshTokenRequest)
    req, err := http.NewRequest(http.MethodPost, "/refresh-token", bytes.NewBuffer(body))
    assert.NoError(suite.T(), err)

    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    router := gin.Default()
    router.POST("/refresh-token", func(c *gin.Context) {
        c.Set("userID", userID)
        suite.refreshTokenController.RefreshToken(c)
    })
    router.ServeHTTP(rr, req)

    assert.Equal(suite.T(), http.StatusBadRequest, rr.Code)
    assert.Contains(suite.T(), rr.Body.String(), "invalid refresh token")
}

func (suite *RefreshTokenControllerTestSuite) TestRefreshToken_UpdateTokensError() {
    userID := "test_user_id"
    refreshTokenRequest := domain.RefreshTokenRequest{
        RefreshToken: "valid_refresh_token",
    }

    suite.mockRefreshTokenUsecase.On("CheckRefreshToken", mock.Anything, userID, refreshTokenRequest.RefreshToken).Return(nil)
    suite.mockRefreshTokenUsecase.On("UpdateTokens", mock.Anything, userID).Return(nil, errors.New("internal server error"))

    body, _ := json.Marshal(refreshTokenRequest)
    req, err := http.NewRequest(http.MethodPost, "/refresh-token", bytes.NewBuffer(body))
    assert.NoError(suite.T(), err)

    req.Header.Set("Content-Type", "application/json")
    rr := httptest.NewRecorder()
    router := gin.Default()
    router.POST("/refresh-token", func(c *gin.Context) {
        c.Set("userID", userID)
        suite.refreshTokenController.RefreshToken(c)
    })
    router.ServeHTTP(rr, req)

    assert.Equal(suite.T(), http.StatusInternalServerError, rr.Code)
    assert.Contains(suite.T(), rr.Body.String(), "internal server error")
}

func TestRefreshTokenControllerTestSuite(t *testing.T) {
    suite.Run(t, new(RefreshTokenControllerTestSuite))
}