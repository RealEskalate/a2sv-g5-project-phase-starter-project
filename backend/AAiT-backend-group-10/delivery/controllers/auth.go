package controllers

import (
	"net/http"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase usecases.UserUsecase
}

func NewUserController(uc usecases.UserUsecase) *UserController {
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
