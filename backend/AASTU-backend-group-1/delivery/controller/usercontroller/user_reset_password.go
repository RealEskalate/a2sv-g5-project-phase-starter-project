package usercontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserController) ResetPassword(ctx *gin.Context) {
	tokenString := ctx.Query("token")
	if tokenString == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Token is required",
		})
		return
	}

	err := u.UserUsecase.ResetPassword(tokenString)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Password reset successfully",
	})
}
