package controllers

import (
	"group3-blogApi/domain"

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
	
	token, err := uc.UserUsecase.Login(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"token": token})


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