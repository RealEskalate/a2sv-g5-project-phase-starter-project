package controller

import (
	"backend-starter-project/domain/dto"
	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"log"
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
	var response dto.Response
	if err := c.ShouldBindJSON(&userRequest); err != nil {
		response.Success = false
		response.Error = "Invalid request payload"
		c.JSON(http.StatusBadRequest, response)
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
		response.Success = false
		response.Error = err.Error()
		c.JSON(http.StatusInternalServerError,response)
		return
	}

	userResponse := dto.UserResponse{
		Username: createdUser.Username,
		Email:    createdUser.Email,
		Role:     createdUser.Role,
	}

	response.Data = userResponse
	c.JSON(http.StatusCreated, response)
}

func (controller *AuthController) Login(c *gin.Context) {
	var user dto.LoginDto
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	var response dto.Response
	refreshToken, accessToken, err := controller.authService.Login(user.Email, user.Password)
	if err != nil {
		response.Success = false
		response.Error = err.Error()
		c.JSON(400, response)
		return
	}

	response.Data = gin.H{ "access_token": accessToken}
	response.Message = "Logged in successfully" 
	c.Set("userId", refreshToken.UserID)
	c.SetCookie("refresh_token", refreshToken.Token, int(refreshToken.ExpiresAt.Unix()), "/", "localhost", false, true)
	c.JSON(http.StatusOK, response)
}

func (controller *AuthController) Logout(c *gin.Context) {
	userId := c.GetString("userId")
	var response dto.Response
	err := controller.authService.Logout(userId)
	log.Println("logout" + userId)
	if err != nil {
		response.Success = false
		response.Error = err.Error()
		c.JSON(400, response)
		return
	}
	response.Success = true
	response.Message = "succesfully logged out"
	c.JSON(200, response)
}

func (controller *AuthController) RefreshAccessToken(c *gin.Context) {
	var token entities.RefreshToken
	var response dto.Response
	err := c.ShouldBindJSON(&token)
	if err != nil {
		response.Success = true
		response.Error = err.Error()
		c.JSON(400, response)
		return
	}

	newToken,err :=  controller.authService.RefreshAccessToken(&token)
	if err != nil {
		response.Success = false
		response.Error = err.Error()
		c.JSON(400, response)
		return
	}
	response.Success = true
	response.Data = gin.H{"access_token": newToken}
	c.JSON(200, response)
}

func (controller *AuthController) VerifyEmail(c *gin.Context) {
	var emailVerification entities.EmailVerificationRequest

	if err := c.ShouldBindJSON(&emailVerification); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	err := controller.authService.VerifyEmail(emailVerification.Email, emailVerification.Code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Email verification failed",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Email verified successfully",
	})
}

func (controller *AuthController) RequestPasswordReset(c *gin.Context) {
	var forgetPasswordRequest entities.ForgetPasswordRequest

	if err := c.ShouldBindJSON(&forgetPasswordRequest); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	err := controller.passwordResetService.RequestPasswordReset(forgetPasswordRequest.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to send password reset link",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Password reset link sent to your email",
	})
}

func (controller *AuthController) ResetPassword(c *gin.Context) {
	var passwordReset entities.PasswordReset

	if err := c.ShouldBindJSON(&passwordReset); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	err := controller.passwordResetService.ResetPassword(passwordReset.Token, passwordReset.NewPassword)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Password reset failed",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Password reset successfully",
	})
}

func (controller *AuthController) ResendOtp(c *gin.Context) {
	var request entities.ResendOTPRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	err = controller.authService.ResendOtp(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to resend OTP",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "OTP sent successfully",
	})
}
