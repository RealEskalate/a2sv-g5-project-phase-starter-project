package controllers

import (
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/usecases"
	"fmt"
	"github.com/gin-gonic/gin"
)


type UserController struct {
	user_usecase usecases.UserUsecaseInterface
}
type UserControllerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
	VerifyEmail(c *gin.Context)
}

func NewUserController(u usecases.UserUsecaseInterface) UserControllerInterface {
	return &UserController{
		user_usecase: u,
	}
}
func (uc *UserController) Register(c *gin.Context) {
    var user *models.User
    err := c.ShouldBindJSON(&user)
    if err != nil {
        fmt.Println("Error binding JSON:", err)
        c.JSON(400, gin.H{"message": "invalid json format"})
        return
    }

    fmt.Println("User data received:", user)

    registeredUser, err := uc.user_usecase.SignUp(user)
    if err != nil {
        fmt.Println("Error during SignUp:", err)
        c.JSON(500, gin.H{"error": err.Error()})
        return
    }

    fmt.Println("User registered successfully with ID:", registeredUser.ID.Hex())
    c.JSON(200, gin.H{"user": registeredUser.ID.Hex()})
}

func (uc *UserController) Login(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "json format error"})
		return
	}
	accessTkn, refreshTkn, err := uc.user_usecase.Login(user)
	if err != nil {
		c.JSON(500, gin.H{"error":err.Error()})
		return
	}
	c.JSON(200, gin.H{"accessToken": accessTkn, "refreshToken": refreshTkn})
}

func (uc *UserController) RefreshToken(c *gin.Context) {
	var request struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, gin.H{"error": "Refresh token is required"})
		return
	}

	newAccessToken, err := uc.user_usecase.RefreshToken(request.RefreshToken)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"access_token": newAccessToken})
}

func (uc *UserController) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(400, gin.H{"error": "Missing token"})
		return
	}

	accTkn, refTkn, err := uc.user_usecase.VerifyEmailToken(token)
	if err != nil {
		c.JSON(401, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Email successfully verified", "access_token": accTkn, "refresh_token": refTkn})
}

