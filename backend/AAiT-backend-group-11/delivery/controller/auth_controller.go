package controller

import (
	"backend-starter-project/domain/dto"
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthController struct {
	authService interfaces.AuthenticationService
}
func NewAuthController(authService interfaces.AuthenticationService) *AuthController{
	return &AuthController{authService: authService}
}

func (controller *AuthController) RegisterUser(c *gin.Context){
	
	var userRequest entities.User
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entities.User{
		ID: primitive.NewObjectID(),
		Username: userRequest.Username,
		Email: userRequest.Email,
		Password: userRequest.Password,
		IsVerified: false,
		Role: "user",
		Profile: entities.Profile{},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createdUser, err := controller.authService.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"data": createdUser})
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
}


func (controller *AuthController) Logout(c *gin.Context){
	userId:=c.GetString("userId")
	
	controller.authService.Logout(userId)
	c.JSON(200,gin.H{"message":"succesfully logged out"})
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

func (controller *AuthController) VerifyEmail(c *gin.Context)  {

	var emailVerification entities.EmailVerificationRequest

	if err := c.ShouldBindJSON(&emailVerification); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.authService.VerifyEmail(emailVerification.Email, emailVerification.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email verified successfully"})
}