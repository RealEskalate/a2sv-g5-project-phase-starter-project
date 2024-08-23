package usercontroller

import (
	"blogs/config"
	"blogs/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserController) LogoutUser(ctx *gin.Context) {
	claims, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Error:   "cannot get claims",
		})
		return
	}

	err := u.UserUsecase.LogoutUser(claims.Username)
	if err != nil {
		code := config.GetStatusCode(err)
		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Failed to logout",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Successfully logout",
	})
}
