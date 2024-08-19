package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserController) VerifyUser(ctx *gin.Context) {
	token := ctx.Query("token")
	if token == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "token is required"})
		return
	}

	err := u.UserUsecase.VerifyUser(token)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user verified"})
}
