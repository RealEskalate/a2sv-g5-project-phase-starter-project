package controller

import (
	"blog/config"
	"blog/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type SignupController struct {
	SignupUsecase domain.SignupUsecase
	Env           *config.Env
}

// Signup creates a new user
func (sc *SignupController) Signup(c *gin.Context) {
	var user domain.AuthSignup
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	returnedUser, _ := sc.SignupUsecase.GetUserByEmail(c, user.Email)
	if returnedUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email already exists"})
		return
	}
	returnedUser, _ = sc.SignupUsecase.GetUserByUsername(c, user.Username)
	if returnedUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	userID, err := sc.SignupUsecase.RegisterUser(c, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	Accesstoken, err := sc.SignupUsecase.CreateAccessToken(&user, sc.Env.AccessTokenSecret, sc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	RefreshToken, err := sc.SignupUsecase.CreateRefreshToken(&user, sc.Env.RefreshTokenSecret, sc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	resp := domain.SignUpResponse{
		ID:           *userID,
		AcessToken:   Accesstoken,
		RefreshToken: RefreshToken,
	}
	c.JSON(http.StatusOK, gin.H{"data": resp})

}
