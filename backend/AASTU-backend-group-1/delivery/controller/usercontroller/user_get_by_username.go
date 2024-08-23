package usercontroller

import (
	"blogs/config"
	"blogs/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (uc *UserController) GetUserByUsername(ctx *gin.Context) {
	username := ctx.Param("username")
	err := config.IsValidUsername(username)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid username",
		})
		return
	}

	user, err := uc.UserUsecase.GetUserByUsername(username)
	if err != nil {
		code := http.StatusInternalServerError

		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Failed to get user",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    user,
	})
}
