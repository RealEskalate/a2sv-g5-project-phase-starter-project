package controllers

import (
	"group3-blogApi/domain"
	"group3-blogApi/infrastracture"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserUsecase domain.UserUsecase
}

func NewUserController(userUsecase domain.UserUsecase) *UserController {
	return &UserController{UserUsecase: userUsecase}
}

func (uc *UserController) Login(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	ipAddress := c.ClientIP()
    userAgent := c.Request.UserAgent()
    deviceFingerprint := infrastracture.GenerateDeviceFingerprint(ipAddress, userAgent)

	LogInResponse, err := uc.UserUsecase.Login(&user, deviceFingerprint)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	
	c.JSON(200, gin.H{"tokens": LogInResponse})


}

func (uc *UserController) RefreshToken(c *gin.Context) {
	var refreshRequest domain.RefreshTokenRequest

	if err := c.ShouldBindJSON(&refreshRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
	}
	ipAddress := c.ClientIP()
    userAgent := c.Request.UserAgent()
    deviceFingerprint := infrastracture.GenerateDeviceFingerprint(ipAddress, userAgent)

	refreshResponse, err := uc.UserUsecase.RefreshToken(refreshRequest.UserID, deviceFingerprint, refreshRequest.Token)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(200, gin.H{"tokens": refreshResponse})
}

func (uc *UserController) Register(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}


	err := uc.UserUsecase.Register(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Your Account created successfully please check your email to activate your account"})

}


func (uc *UserController) ActivateAccount(c *gin.Context) {
	token := c.Query("token")
	Email := c.Query("Email")
	err := uc.UserUsecase.AccountActivation(token, Email)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Account activated successfully"})
}