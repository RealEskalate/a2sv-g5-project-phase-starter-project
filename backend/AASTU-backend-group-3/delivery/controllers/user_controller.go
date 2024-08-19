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

	LogInResponse, err := uc.UserUsecase.Login(&user, deviceFingerprint)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
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
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Logged out successfully"})
}

// TODO: get userID from context rather than query
func (uc *UserController) LogoutAll(c *gin.Context) {
	userID := c.Query("userID")
	err := uc.UserUsecase.LogoutAllDevices(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Logged out from all devices"})
}

// TODO: get userID from context rather than query
func (uc *UserController) LogoutDevice(c *gin.Context) {
	userID := c.Query("userID")
	deviceID := c.Query("deviceID")
	err := uc.UserUsecase.LogoutDevice(userID, deviceID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Logged out successfully"})
}

// TODO: get userID from context rather than query
func (uc *UserController) GetDevices(c *gin.Context) {
	userID := c.Query("userID")
	devices, err := uc.UserUsecase.GetDevices(userID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
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


// reset password

func (uc *UserController) SendPasswordResetLink(c *gin.Context) {
	var req domain.ResetPasswordRequest
	
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := uc.UserUsecase.SendPasswordResetLink(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password has been reset"})
}