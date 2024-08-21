package controllers

import (
	domain "AAiT-backend-group-2/Domain"

	"AAiT-backend-group-2/Infrastructure/dtos"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct {
	UserUsecase domain.UserUsecase
}


func NewAuthController(userUsecase domain.UserUsecase) *AuthController {
	return &AuthController{
		UserUsecase: userUsecase,
	}
}

func (ctr *AuthController) CreateUser(c *gin.Context) {
	var user domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	err := ctr.UserUsecase.CreateUser(c, user)

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"message": "User Registerd successfully!"})
}

func (ctr *AuthController) Login(c *gin.Context) {
	var loginDto dtos.LoginDTO

	if err := c.BindJSON(&loginDto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

	token, err := ctr.UserUsecase.Login(c, &loginDto)

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"token": token,
	})
}

func (ctr *AuthController) RefreshToken(c *gin.Context) {
	var refreshTokenDto dtos.RefreshTokenDto
	
	if err := c.ShouldBindJSON(&refreshTokenDto); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := ctr.UserUsecase.RefreshToken(c, refreshTokenDto.RefreshToken)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, tokens)
}


