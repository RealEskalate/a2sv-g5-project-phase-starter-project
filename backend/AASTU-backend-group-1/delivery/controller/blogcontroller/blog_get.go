package blogcontroller

import (
	"blogs/config"
	"blogs/domain"
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
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "page is required",
		})
		return
	}

	var page, size int
	_, err := fmt.Sscanf(pageString, "%d", &page)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "page must be a number",
		})
		return
	}

	if sizeString == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "size is required",
		})
		return
	}

	_, err = fmt.Sscanf(sizeString, "%d", &size)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "size must be a number",
		})
		return
	}

	if sortBy == "" {
		sortBy = "date"
	}

	if reverse == "" {
		reverse = "false"
	}

	blogs, total, err := b.BlogUsecase.GetBlogs(sortBy, page, size, reverse == "true")
	if err != nil {
		code := config.GetStatusCode(err)
		log.Println(err)

		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Cannot get blogs",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data: gin.H{
			"blogs":       blogs,
			"total_pages": total,
		},
	})
}

func (b *BlogController) GetBlogByID(ctx *gin.Context) {
	idHex := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "invalid id",
		})
		return
	}

	blog, err := b.BlogUsecase.GetBlogByID(id.Hex())
	if err != nil {
		code := config.GetStatusCode(err)
		log.Println(err)

		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Cannot get blog",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    blog,
	})
}
