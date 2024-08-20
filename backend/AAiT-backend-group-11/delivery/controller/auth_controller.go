package controller

import (
	"backend-starter-project/domain/dto"
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	authService interfaces.AuthenticationService
}

func (controller *AuthController) RegisterUser() {

}

func (controller *AuthController) Login(c *gin.Context) {
	var user dto.LoginDto
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	refreshToken, accessToken, err := controller.authService.Login(user.Email, user.Password)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.Header("Authorization", "Bearer "+accessToken)
	c.SetCookie("refresh_token", refreshToken.Token, int(refreshToken.ExpiresAt.Unix()), "/", "localhost", false, true)
	return
}


func (controller *AuthController) Logout(c *gin.Context){
	userId:=c.GetString("userId")
	
	controller.authService.Logout(userId)
	c.JSON(200,gin.H{"message":"succesfully loged out"})
	return

}

func (controller *AuthController) RefreshAccessToken(c *gin.Context){
	var token entities.RefreshToken
	err := c.ShouldBindJSON(&token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	controller.authService.RefreshAccessToken(&token)
}