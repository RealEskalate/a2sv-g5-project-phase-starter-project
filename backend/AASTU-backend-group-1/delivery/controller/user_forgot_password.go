package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserController) ForgotPassword(ctx *gin.Context) {
	var input struct {
		Email       string `json:"email"`
		NewPassword string `json:"new_password"`
	}

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if input.Email == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email field is required"})
		return
	}

	if input.NewPassword == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "new_password field is required"})
		return
	}

	err = u.UserUsecase.ForgotPassword(input.Email, input.NewPassword)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Password reset token sent successfully",
	})
}
