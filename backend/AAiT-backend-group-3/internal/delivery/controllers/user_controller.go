package controllers

import (
	"github.com/gin-gonic/gin"
	"AAIT-backend-group-3/internal/usecases"
	"AAIT-backend-group-3/internal/domain/models"
)


type UserController struct {
	user_usecase *usecases.UserUsecase
}

func NewController(u *usecases.UserUsecase) *UserController{
	return &UserController{
		user_usecase: u,
	}
}


func (usecases *UserController) Register(c *gin.Context) {
	var user *models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(400, gin.H{"message": "invalid json format"})
		return 
	}
	err = usecases.user_usecase.SignUp(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err})
	}
	c.JSON(200, gin.H{"message": "User registered successfully"})
}

func (usecases *UserController) Login(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message":"json format error"})
		return
	}
	accessTkn, refreshTkn, err := usecases.user_usecase.Login(user)
	if err != nil {
		c.JSON(500, gin.H{"error":err})
		return
	}
	c.JSON(200, gin.H{"accessToken": accessTkn, "refreshToken":refreshTkn})
}


func (usecases *UserController) RefreshToken(c *gin.Context) {
    userID, exists := c.Get("userID")
    if !exists {
        c.JSON(401, gin.H{"error": "User ID not found in request"})
        return
    }
    var request struct {
        RefreshToken string `json:"refresh_token" binding:"required"`
    }
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(400, gin.H{"error": "Refresh token is required"})
        return
    }

    newAccessToken, err := usecases.user_usecase.RefreshToken(userID.(string), request.RefreshToken)
    if err != nil {
        c.JSON(401, gin.H{"error": err.Error()})
        return
    }
    c.JSON(200, gin.H{"access_token": newAccessToken})
}
