package blogcontroller

import (
	"blogs/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (l *BlogController) UpdateBlogByID(ctx *gin.Context) {
	id := ctx.Param("id")
	var blog domain.Blog
	if err := ctx.ShouldBindJSON(&blog); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claim, ok := ctx.MustGet("claims").(*domain.LoginClaims)

	if !ok {
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	err := l.BlogUsecase.UpdateBlogByID(id, &blog, claim)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusOK, "log updated")
}
