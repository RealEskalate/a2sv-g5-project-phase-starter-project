package controllers

import (
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/usecases"
	"fmt"

	"github.com/gin-gonic/gin"
)


type UserController struct {
	user_usecase *usecases.UserUsecase
}

func NewUserController(u *usecases.UserUsecase) *UserController{
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
		c.JSON(500, gin.H{"error": err.Error()})
		return 
	}
	c.JSON(200, gin.H{"message":"verification email sent successfully"})
}

func (usecases *UserController) Login(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message":"json format error"})
		return
	}
	accessTkn, refreshTkn, err := usecases.user_usecase.Login(user)
	if err != nil {
		c.JSON(500, gin.H{"error":err.Error()})
		return
	}
	c.JSON(200, gin.H{"accessToken": accessTkn, "refreshToken":refreshTkn})
}


func (usecases *UserController) RefreshToken(c *gin.Context) {
    var request struct {
        RefreshToken string `json:"refresh_token" binding:"required"`
    }
    if err := c.ShouldBindJSON(&request); err != nil {
        c.JSON(400, gin.H{"error": "Refresh token is required"})
        return
    }

    newAccessToken, err := usecases.user_usecase.RefreshToken(request.RefreshToken)
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
	c.JSON(200, gin.H{"message": "Email successfully verified", "access_token": accTkn, "refresh_token":refTkn})
}

func (uc *UserController) Logout(c *gin.Context){
	token, ok := c.Get("token")
	if !ok {
		c.JSON(400, gin.H{"error": "Missing token"})
		return
	}
	tokenStr, ok := token.(string)
	fmt.Println(tokenStr)
	
	if !ok {
		c.JSON(500, gin.H{"error": "Invalid token format"})
		return
	}
	err := uc.user_usecase.Logout(tokenStr)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to logout"})
		return
	}
	c.JSON(200, gin.H{"message": "Successfully logged out"})
}