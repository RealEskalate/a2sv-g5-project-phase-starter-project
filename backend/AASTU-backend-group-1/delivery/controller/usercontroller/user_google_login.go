package usercontroller

import (
	"blogs/config"
	"blogs/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserController) GoogleLogin(ctx *gin.Context) {
	url, err := u.UserUsecase.GoogleLogin()
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Error:   "Failed to get google login url",
		})
		return
	}

	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

func (u *UserController) GoogleCallback(ctx *gin.Context) {
	state := ctx.Query("state")
	code := ctx.Query("code")

	if state == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "State is required",
		})
		return
	}

	if code == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "Code is required",
		})
		return
	}

	accessToken, refreshToken, err := u.UserUsecase.GoogleCallback(state, code)

	if err != nil {
		statusCode := config.GetStatusCode(err)
		ctx.JSON(statusCode, domain.APIResponse{
			Status:  statusCode,
			Message: "Failed to login with google",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Successfully login with google",
		Data: gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	})
}
