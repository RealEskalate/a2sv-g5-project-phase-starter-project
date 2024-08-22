package blogcontroller

import (
	"blogs/config"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (b *BlogController) SearchBlog(ctx *gin.Context) {
	title := ctx.Query("title")
	author := ctx.Query("author")
	tags := ctx.QueryArray("tags")

	if title == "" && author == "" && len(tags) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "title, author or tags is required"})
		return
	}

	if len(tags) == 0 {
		tags = []string{}
	}

	blogs, err := b.BlogUsecase.SearchBlog(title, author, tags)
	if err != nil {
		code := config.GetStatusCode(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		ctx.JSON(code, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{"counts": len(blogs), "data": blogs})
}
