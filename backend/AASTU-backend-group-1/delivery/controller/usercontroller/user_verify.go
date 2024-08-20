package usercontroller

import (
	"blogs/config"
	"log"
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
		code := config.GetStatusCode(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(code, gin.H{"error": "Internal server error"})
			return
		}

		ctx.JSON(code, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "user verified"})
}
