package controllers

import (
	"net/http"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase usecases.IAuthUsecase
}
type ForgotPasswordRequestDTO struct {
	Email string `json:"email" binding:"required,email"`
}

type ResetPasswordRequestDTO struct {
	Token       string `json:"token" binding:"required"`
	NewPassword string `json:"new_password" binding:"required,min=8"`
}

func NewUserController(uc usecases.IAuthUsecase) *UserController {
	return &UserController{
		userUsecase: uc,
	}
}

func (ctrl *UserController) Register(c *gin.Context) {
	var userDTO domain.RegisterUserDTO
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

func (uc *UserController) Login(c *gin.Context) {
	var loginDTO domain.LoginUserDTO
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

func (uc *UserController) RefreshToken(c *gin.Context) {
	var refreshTokenDTO struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	if err := c.ShouldBindJSON(&refreshTokenDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := uc.userUsecase.RefreshTokens(refreshTokenDTO.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tokens)
}

func (uc *UserController) ForgotPassword(c *gin.Context) {
	var forgotPasswordDTO usecases.ForgotPasswordRequestDTO
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

func (uc *UserController) ResetPassword(c *gin.Context) {
	var resetPasswordDTO usecases.ResetPasswordRequestDTO
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
