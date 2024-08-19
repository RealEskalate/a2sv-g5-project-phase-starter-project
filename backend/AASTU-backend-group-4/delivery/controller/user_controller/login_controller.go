package user_controller

import (
	"blog-api/domain"
	"blog-api/infrastructure/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UserController) LoginController(c *gin.Context) {
	var loginRequest domain.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload. Please provide email or username and password."})
		return
	}

	if loginRequest.Email == "" && loginRequest.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Either email or username must be provided."})
		return
	}

	var u domain.User
	var err error
	if loginRequest.Email != "" {
		u, err = uc.usecase.GetByEmail(c, loginRequest.Email)
	} else if loginRequest.Username != "" {
		u, err = uc.usecase.GetByUsername(c, loginRequest.Username)
	}

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials. User not found."})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(loginRequest.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials. Incorrect password."})
		return
	}

	accessToken, err := auth.CreateAccessToken(&u, uc.Env.AccessTokenSecret, uc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate access token"})
		return
	}

	refreshToken, err := auth.CreateRefreshToken(&u, uc.Env.RefreshTokenSecret, uc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate refresh token"})
		return
	}

	response := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, response)
}
