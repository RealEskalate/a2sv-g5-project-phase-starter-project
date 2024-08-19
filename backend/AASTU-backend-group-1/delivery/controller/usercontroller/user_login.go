package usercontroller

import (
	"blogs/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserController) LoginUser(ctx *gin.Context) {
	var userData struct {
		Email    string `json:"email"`
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := ctx.ShouldBindJSON(&userData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if userData.Email == "" && userData.Username == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "email or username is required"})
		return
	}

	if userData.Password == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "password is required"})
		return
	}

	var accessToken, refreshToken string
	if userData.Email != "" {
		accessToken, refreshToken, err = u.UserUsecase.LoginUser(userData.Email, userData.Password)
	} else {
		accessToken, refreshToken, err = u.UserUsecase.LoginUser(userData.Username, userData.Password)
	}

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

	ctx.JSON(http.StatusOK, gin.H{
		"message":       "Login successfull",
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}
