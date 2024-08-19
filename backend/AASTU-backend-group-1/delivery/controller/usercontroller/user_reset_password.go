package usercontroller

import (
	"blogs/config"
	"log"
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
		code := config.GetStatusCode(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(code, gin.H{"error": "Internal server error"})
			return
		}

		ctx.JSON(code, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{
		"message": "Password reset successfully",
	})
}
