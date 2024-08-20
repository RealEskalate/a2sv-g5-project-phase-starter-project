package blogcontroller

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (b *BlogController) GetBlogs(ctx *gin.Context) {
	pageString := ctx.Query("page")
	sizeString := ctx.Query("size")
	sortBy := ctx.Query("sort_by")
	reverse := ctx.Query("reverse")

	if pageString == "" {
		ctx.JSON(http.StatusBadRequest, "page is required")
		return
	}

	var page, size int
	_, err := fmt.Sscanf(pageString, "%d", &page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "page must be a number")
		return
	}

	_, err = fmt.Sscanf(sizeString, "%d", &size)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "size must be a number")
		return
	}

	if sizeString == "" {
		ctx.JSON(http.StatusBadRequest, "size is required")
		return
	}

	if sortBy == "" {
		sortBy = "date"
	}

	if reverse == "" {
		reverse = "false"
	}

	blogs, err := b.BlogUsecase.GetBlogs(sortBy, page, size, reverse == "true")
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count": len(blogs),
		"blogs": blogs,
	})
}


// GetBlogByID ...

func (b *BlogController) GetBlogByID(ctx *gin.Context) {
	id := ctx.Param("id")
	blog, err := b.BlogUsecase.GetBlogByID(id)
	if err != nil {
		log.Println(err)
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusOK, blog)
}

