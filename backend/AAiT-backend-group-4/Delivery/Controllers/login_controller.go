package controllers

import (
	bootstrap "aait-backend-group4/Bootstrap"
	domain "aait-backend-group4/Domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

func (lc *LoginController) Login(c *gin.Context) {
	var request domain.LoginRequest

	err := c.ShouldBindJSON(&request)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	accessToken, refreshToken, err := lc.LoginUsecase.LoginWithIdentifier(c, request.Identifier)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	msg := map[string]string{
		"accessToken":  accessToken,
		"refreshToken": refreshToken,
	}

	c.JSON(http.StatusOK, msg)
}
