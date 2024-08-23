package usercontroller

import (
	"blogs/config"
	"blogs/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserController) ChangePassword(ctx *gin.Context) {
	claims, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Error:   "cannot get claims",
		})
		return
	}

	var Request struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}

	err := ctx.ShouldBind(&Request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	if Request.OldPassword == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Old password is required",
		})
		return
	}

	if Request.NewPassword == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "New password is required",
		})
		return
	}

	err = u.UserUsecase.ChangePassword(claims.Username, Request.OldPassword, Request.NewPassword)
	if err != nil {
		code := config.GetStatusCode(err)

		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Failed to change password",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Password changed successfully",
	})
}
