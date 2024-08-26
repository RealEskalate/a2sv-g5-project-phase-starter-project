package blogcontroller

import (
	"blogs/config"
	"blogs/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (b *BlogController) DeleteBlogByID(ctx *gin.Context) {
	id := ctx.Param("id")

	claims, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		log.Println("Error getting claims")
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Error:   "Error getting claims",
		})

		return
	}

	err := b.BlogUsecase.DeleteBlogByID(id, claims)
	if err != nil {
		code := config.GetStatusCode(err)
		log.Println(err)

		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Error deleting blog",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
