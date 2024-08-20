package blogcontroller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (b *BlogController) SearchBlog(ctx *gin.Context) {
	var search struct {
		Title  string   `json:"title"`
		Author string   `json:"author"`
		Tags   []string `json:"tags"`
	}

	if err := ctx.ShouldBindJSON(&search); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blogs, err := b.BlogUsecase.SearchBlog(search.Title, search.Author, search.Tags)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

