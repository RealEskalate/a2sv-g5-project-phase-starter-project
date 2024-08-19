package controllers

import (
	"AAiT-backend-group-8/Domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserUsecase Domain.IUserUseCase
}

func NewUserHandler(userUseCase Domain.IUserUseCase) *UserHandler {
	return &UserHandler{UserUsecase: userUseCase}
}

func (h *UserHandler) RegisterUser(c *gin.Context) {
	var user Domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err := h.UserUsecase.RegisterUser(&user)
	if err != nil {
		if err.Error() == "email already exists" {
			c.JSON(400, gin.H{"error": err.Error()})
		} else {
			c.JSON(500, gin.H{"error": err.Error()})
		}
		return
	}

	c.JSON(200, gin.H{"message": "Registration successful. Check your email for verification link."})
}

func (h *UserHandler) VerifyEmail(c *gin.Context) {
	token := c.Query("token")
	if token == "" {
		c.JSON(400, gin.H{"error": "Invalid token"})
		return
	}

	err := h.UserUsecase.VerifyEmail(token)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"message": "Email verified successfully"})
}

func (h *UserHandler) Login(c *gin.Context) {
	type email_pass struct {
		email    string
		password string
	}

	var ep email_pass

	bind_err := c.BindJSON(&ep)
	if bind_err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid request payload"})
		return
	}

	token, refresher, err := h.UserUsecase.Login(ep.email, ep.password)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": token, "refresher": refresher})

}
