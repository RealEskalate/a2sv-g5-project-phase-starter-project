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
	passwordResetService interfaces.PasswordResetService
}

func NewAuthController(authService interfaces.AuthenticationService, passwordResetService interfaces.PasswordResetService) *AuthController{
	return &AuthController{
		authService: authService,
		passwordResetService: passwordResetService}
	}

func (controller *AuthController) RegisterUser(c *gin.Context) {

	var userRequest dto.UserCreateRequestDTO

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := entities.User{
		ID:         primitive.NewObjectID(),
		Username:   userRequest.Username,
		Email:      userRequest.Email,
		Password:   userRequest.Password,
		IsVerified: false,
		Role:       "user",
		Profile:    entities.Profile{},
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	createdUser, err := controller.authService.RegisterUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	userResponse := dto.UserResponse{
		Username: createdUser.Username,
		Email:    createdUser.Email,
		Role:     createdUser.Role,
	}


	c.JSON(http.StatusCreated, gin.H{"data": userResponse})
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
	c.Set("userId", refreshToken.UserID)

	c.JSON(200, gin.H{"access_token": accessToken})
	c.JSON(200, gin.H{"refresh_token": refreshToken.Token})
	c.JSON(200, gin.H{"message": "login successful"})
}

func (controller *AuthController) Logout(c *gin.Context) {
	userId := c.GetString("userId")

	err:=controller.authService.Logout(userId)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.SetCookie("refresh_token", "", -1, "/", "localhost", false, true)
	c.Header("Authorization", "")
	c.JSON(200, gin.H{"message": "succesfully logged out"})
}

func (controller *AuthController) RefreshAccessToken(c *gin.Context) {
	refresh,err:=c.Cookie("refresh_token")
	
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	controller.authService.RefreshAccessToken(refresh)
}

func (controller *AuthController) VerifyEmail(c *gin.Context) {

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


func (controller *AuthController) RequestPasswordReset(c *gin.Context) {
	
	var forgetPasswordRequest entities.ForgetPasswordRequest

	if err := c.ShouldBindJSON(&forgetPasswordRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.passwordResetService.RequestPasswordReset(forgetPasswordRequest.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset link sent to your email"})
}

func (controller *AuthController) ResetPassword(c *gin.Context) {

	var passwordReset entities.PasswordReset

	if err := c.ShouldBindJSON(&passwordReset); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.passwordResetService.ResetPassword(passwordReset.Token, passwordReset.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}

func (controller *AuthController) ResendOtp(c *gin.Context) {

    var request entities.ResendOTPRequest

    err := c.ShouldBind(&request)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    err = controller.authService.ResendOtp(request)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "OTP sent successfully"})
}