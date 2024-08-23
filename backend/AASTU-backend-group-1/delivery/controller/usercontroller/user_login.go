package usercontroller

import (
	"blogs/config"
	"blogs/domain"
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
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	if userData.Email == "" && userData.Username == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "email or username is required",
		})
		return
	}

	if userData.Password == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "password is required",
		})
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
		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Failed to login",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data: gin.H{
			"access_token":  accessToken,
			"refresh_token": refreshToken,
		},
	})
}
