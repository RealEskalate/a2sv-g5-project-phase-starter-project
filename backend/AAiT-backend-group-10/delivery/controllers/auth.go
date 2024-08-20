package controllers

import (
	"net/http"

	"aait.backend.g10/usecases"
	"aait.backend.g10/usecases/dto"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	userUsecase usecases.IAuthUsecase
}

func NewAuthController(uc usecases.IAuthUsecase) *AuthController {
	return &AuthController{
		userUsecase: uc,
	}
}

func (ctrl *AuthController) Register(c *gin.Context) {
	var userDTO dto.RegisterUserDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	user, err := ctrl.userUsecase.RegisterUser(&userDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (uc *AuthController) Login(c *gin.Context) {
	var loginDTO dto.LoginUserDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := uc.userUsecase.LoginUser(&loginDTO)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func (uc *AuthController) RefreshToken(c *gin.Context) {
	var refreshToken dto.RefreshTokenDTO
	if err := c.ShouldBindJSON(&refreshToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := uc.userUsecase.RefreshTokens(refreshToken.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func (uc *AuthController) ForgotPassword(c *gin.Context) {
	var forgotPasswordDTO dto.ForgotPasswordRequestDTO
	if err := c.ShouldBindJSON(&forgotPasswordDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uc.userUsecase.ForgotPassword(&forgotPasswordDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset link sent to your email"})
}

func (uc *AuthController) ResetPassword(c *gin.Context) {
	var resetPasswordDTO dto.ResetPasswordRequestDTO
	if err := c.ShouldBindJSON(&resetPasswordDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := uc.userUsecase.ResetPassword(&resetPasswordDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Password reset successfully"})
}
