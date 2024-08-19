package usercontroller

import (
	"blogs/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserController) RefreshToken(ctx *gin.Context) {
	claims, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		log.Println("Error getting claims")
		ctx.JSON(http.StatusInternalServerError, "Internal server error")
		return
	}

	accessToken, err := u.UserUsecase.RefreshToken(claims)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token": accessToken,
	})
}
