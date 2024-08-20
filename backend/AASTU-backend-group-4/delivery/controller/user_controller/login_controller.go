package user_controller

import (
	"blog-api/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) LoginController(c *gin.Context) {
	var loginRequest domain.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload. Please provide email or username and password."})
		return
	}

	if loginRequest.Email == "" && loginRequest.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Either email or username must be provided."})
		return
	}

	response, err := uc.usecase.LoginUser(c, loginRequest, uc.Env)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
