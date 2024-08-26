package controllers

import (
	"group3-blogApi/domain"
	"group3-blogApi/infrastracture"
	"net/http"

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

	LogInResponse, uerr := uc.UserUsecase.Login(&user, deviceFingerprint)
	if uerr.Message != "" {
		c.JSON(uerr.StatusCode, gin.H{"error": uerr.Message})
		return
	}

	c.JSON(200, gin.H{"tokens": LogInResponse})

}

// TODO: get userID from context rather than query
func (uc *UserController) Logout(c *gin.Context) {
	var logoutRequest domain.LogoutRequest
	if err := c.ShouldBindJSON(&logoutRequest); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ipAddress := c.ClientIP()
	userAgent := c.Request.UserAgent()
	deviceFingerprint := infrastracture.GenerateDeviceFingerprint(ipAddress, userAgent)

	err := uc.UserUsecase.Logout(logoutRequest.UserID, deviceFingerprint, logoutRequest.Token)
	if err.Message != "" {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(200, gin.H{"message": "Logged out successfully"})
}

// TODO: get userID from context rather than query
func (uc *UserController) LogoutAll(c *gin.Context) {
	userID := c.GetString("user_id")
	err := uc.UserUsecase.LogoutAllDevices(userID)
	if err.Message != "" {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(200, gin.H{"message": "Logged out from all devices"})
}

// TODO: get userID from context rather than query
func (uc *UserController) LogoutDevice(c *gin.Context) {
	userID := c.GetString("user_id")
	deviceID := c.Query("deviceID")
	err := uc.UserUsecase.LogoutDevice(userID, deviceID)
	if err.Message != "" {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(200, gin.H{"message": "Logged out successfully"})
}

// TODO: get userID from context rather than query
func (uc *UserController) GetDevices(c *gin.Context) {
	userID := c.GetString("user_id")
	devices, err := uc.UserUsecase.GetDevices(userID)
	if err.Message != "" {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(200, gin.H{"devices": devices})
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
	if err.Message != "" {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
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
	if err.Message != "" {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(200, gin.H{"message": "Your Account created successfully please check your email to activate your account"})

}

func (uc *UserController) ActivateAccountMe(c *gin.Context) {
	userID := c.GetString("user_id")

	err := uc.UserUsecase.ActivateAccountMe(userID)
	if err.Message != "" {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(200, gin.H{"message": "Check your email."})
}

func (uc *UserController) ActivateAccount(c *gin.Context) {
	var activateReq domain.ActivateRequest
	if err := c.ShouldBindJSON(&activateReq); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	err := uc.UserUsecase.AccountActivation(activateReq.ActivationToken, activateReq.Email)
	if err.Message != "" {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(200, gin.H{"message": "Account activated successfully"})
}

// reset password

func (uc *UserController) SendPasswordResetLink(c *gin.Context) {
	var req domain.ResetPasswordRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := uc.UserUsecase.SendPasswordResetLink(req.Email)
	if err.Message != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset link sent"})
}

func (uc *UserController) ResetPassword(c *gin.Context) {
	var req struct {
		Password string `json:"password"`
	}
	token := c.Param("token")
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := uc.UserUsecase.ResetPassword(token, req.Password)
	if err.Message != "" {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password has been reset"})
}
