package controllers

import (
	"AAIT-backend-group-3/internal/domain/dtos"
	"AAIT-backend-group-3/internal/domain/models"
	"AAIT-backend-group-3/internal/usecases"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type UserController struct {
	user_usecase usecases.UserUsecaseInterface
}
type UserControllerInterface interface {
	Register(c *gin.Context)
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
	VerifyEmail(c *gin.Context)
	Logout(c *gin.Context)
	UpdateProfile(c *gin.Context)
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
		c.JSON(400, gin.H{"message": "invalid json format"})
		return
	}
	_, err = uc.user_usecase.SignUp(user)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "verification email sent to your email"})
}

func (uc *UserController) Login(c *gin.Context) {
	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"message": "json format error"})
		return
	}
	accessTkn, refreshTkn, err := uc.user_usecase.Login(user)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
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

func (uc *UserController) Logout(c *gin.Context) {
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

func (uc *UserController) UpdateProfile(c *gin.Context){
	claims, _ := c.Get("claims")
    jwtClaims, ok := claims.(jwt.MapClaims)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse claims"})
        return
    }
    userID, ok := jwtClaims["userId"].(string)
    if !ok {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to parse user ID"})
        return
    }

	var updateData dtos.Profile
    if username := c.PostForm("username"); username != "" {
        updateData.UserName = username
    }
    if phoneNum := c.PostForm("phone_num"); phoneNum != "" {
        updateData.PhoneNum = phoneNum
    }
    if bio := c.PostForm("bio"); bio != "" {
        updateData.Bio = bio
    }

    image, _, err := c.Request.FormFile("profile_pic")
    if err != nil && err != http.ErrMissingFile {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to process profile picture"})
        return
    }
    if err := uc.user_usecase.UpdateProfile(userID, &updateData, image); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update profile"})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Profile updated successfully"})
}