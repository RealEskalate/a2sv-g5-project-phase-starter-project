package blogcontroller

import (
	"blogs/config"
	"blogs/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (b *BlogController) SearchBlog(ctx *gin.Context) {
	title := ctx.Query("title")
	author := ctx.Query("author")
	tags := ctx.QueryArray("tags")

	if title == "" && author == "" && len(tags) == 0 {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "at least one of title, author, or tags is required",
		})
		return
	}

	if len(tags) == 0 {
		tags = []string{}
	}

	blogs, err := b.BlogUsecase.SearchBlog(title, author, tags)
	if err != nil {
		code := config.GetStatusCode(err)
		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Cannot search blog",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Count:   len(blogs),
		Data:    blogs,
	})
}
