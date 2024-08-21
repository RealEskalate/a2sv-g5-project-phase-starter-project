package blogcontroller

import (
	"blogs/config"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (b *BlogController) GetBlogs(ctx *gin.Context) {
	pageString := ctx.Query("page")
	sizeString := ctx.Query("size")
	sortBy := ctx.Query("sort_by")
	reverse := ctx.Query("reverse")

	if pageString == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "page is required"})
		return
	}

	var page, size int
	_, err := fmt.Sscanf(pageString, "%d", &page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "page must be a number"})
		return
	}

	if sizeString == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "size is required"})
		return
	}

	_, err = fmt.Sscanf(sizeString, "%d", &size)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "size must be a number"})
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
		code := config.GetStatusCode(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		ctx.JSON(code, gin.H{"error": err.Error()})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"count": len(blogs),
		"data":  blogs,
	})
}

func (b *BlogController) GetBlogByID(ctx *gin.Context) {
	idHex := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	blog, err := b.BlogUsecase.GetBlogByID(id.Hex())
	if err != nil {
		code := config.GetStatusCode(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		ctx.JSON(code, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, blog)
}
