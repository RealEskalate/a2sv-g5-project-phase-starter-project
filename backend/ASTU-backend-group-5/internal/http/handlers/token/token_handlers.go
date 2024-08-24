package token

import (
	"blogApp/internal/domain"
	"blogApp/internal/usecase"
	"blogApp/pkg/jwt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type TokenHandler struct {
	tokenUseCase usecase.TokenUsecase
}

func NewTokenHandler(tokenUseCase usecase.TokenUsecase) *TokenHandler {
	return &TokenHandler{
		tokenUseCase: tokenUseCase,
	}
}

func (h *TokenHandler) RefreshToken(c *gin.Context) {
	refreshToken := strings.Split(c.GetHeader("Authorization"), " ")[1]
	if refreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Refresh token required"})
		return
	}
	claims, err := jwt.ValidateToken(refreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	refresh, err := jwt.GenerateJWT(claims.ID, claims.Email, claims.Role, claims.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	access, err := jwt.GenerateJWT(claims.ID, claims.Email, claims.Role, claims.Username)
	token := domain.Token{
		RefreshToken: refresh,
		AccessToken:  access,
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (h *TokenHandler) LogOut(c *gin.Context) {
	token := domain.Token{}

	accessTokenHeader := c.GetHeader("x_access_token")
	refreshTokenHeader := c.GetHeader("x_refresh_token")

	accessTokenParts := strings.Split(accessTokenHeader, " ")
	if len(accessTokenParts) != 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid access token format"})
		return
	}
	token.AccessToken = accessTokenParts[1]

	refreshTokenParts := strings.Split(refreshTokenHeader, " ")
	if len(refreshTokenParts) != 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid refresh token format"})
		return
	}
	token.RefreshToken = refreshTokenParts[1]

	// Check if the tokens are empty after splitting
	if token.AccessToken == "" || token.RefreshToken == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Authorization token required"})
		return
	}

	// if err := c.ShouldBindJSON(&token); err != nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	// 	return
	// }
	err := h.tokenUseCase.BlacklistToken(c, token.RefreshToken, domain.TokenType("refresh"), time.Now())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	err = h.tokenUseCase.BlacklistToken(c, token.AccessToken, domain.TokenType("access"), time.Now())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "logged out successfully"})
}
