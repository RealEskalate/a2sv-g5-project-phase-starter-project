package usercontroller

import (
	"blogs/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserController) GoogleLogin(ctx *gin.Context) {
	url, err := u.UserUsecase.GoogleLogin()

	if err != nil {
		log.Println(err)

		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal server error",
		})
		return
	}

	log.Println(url)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (u *UserController) GoogleCallback(ctx *gin.Context) {
	state := ctx.Query("state")
	code := ctx.Query("code")

	if state == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "State is required",
		})
		return
	}

	if code == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Code is required",
		})
		return
	}

	accessToken, refreshToken, err := u.UserUsecase.GoogleCallback(state, code)

	if err != nil {
		statusCode := config.GetStatusCode(err)

		if statusCode == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(statusCode, gin.H{
				"error": "Internal server error",
			})
			return
		}

		ctx.JSON(statusCode, gin.H{
			"error": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
