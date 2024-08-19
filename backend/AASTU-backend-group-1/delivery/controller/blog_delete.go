package controller

import (
	"blogs/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (l *BlogController) DeleteLogByID(ctx *gin.Context)  {
	id := ctx.Param("id")

	claims, ok :=ctx.MustGet("claims").(*domain.LoginClaims)

	if !ok {
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return 
	}

	err := l.BlogUsecase.DeleteBlogByID(id, claims)

	if err != nil {

		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusOK, "log deleted")

	
}