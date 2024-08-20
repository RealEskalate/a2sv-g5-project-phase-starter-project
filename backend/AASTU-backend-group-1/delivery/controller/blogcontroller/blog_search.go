package blogcontroller

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (b *BlogController) SearchBlog(ctx *gin.Context) {

	var search struct {
		Title  string   `json:"title"`
		Author string   `json:"author"`
		Tags   []string `json:"tags"`
	}
	   search.Title = ctx.Query("title")
	   search.Author = ctx.Query("author")
	   search.Tags = ctx.QueryArray("tags")


	if search.Title == "" || search.Author == "" || len(search.Tags) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": errors.New("title, author and tags are required")})
		return
	}

	blogs, err := b.BlogUsecase.SearchBlog(search.Title, search.Author, search.Tags)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}

