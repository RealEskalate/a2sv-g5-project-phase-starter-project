package usercontroller

import (
	"blogs/config"
	"blogs/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *UserController) PromoteUser(ctx *gin.Context) {
	claims, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Error:   "cannot get claims",
		})
		return
	}

	var request struct {
		Username string `json:"username"`
		Promoted bool   `json:"promoted"`
	}

	err := ctx.ShouldBind(&request)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	if request.Username == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Username is required",
		})
		return
	}

	promotionWord := "demoted"
	if request.Promoted {
		promotionWord = "promoted"
	}

	err = u.UserUsecase.PromoteUser(request.Username, request.Promoted, claims)
	if err != nil {
		code := config.GetStatusCode(err)
		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Failed to " + promotionWord + " user",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Successfully " + promotionWord + " user",
	})
}
