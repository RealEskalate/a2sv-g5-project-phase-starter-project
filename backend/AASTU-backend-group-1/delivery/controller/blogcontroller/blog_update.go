package blogcontroller

import (
	"blogs/config"
	"blogs/domain"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (b *BlogController) UpdateBlogByID(ctx *gin.Context) {
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

	var blogUpdate struct {
		Title   string   `json:"title"`
		Content string   `json:"content"`
		Tags    []string `json:"tags"`
	}

	if err := ctx.ShouldBindJSON(&blogUpdate); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   err.Error(),
		})
		return
	}

	if blogUpdate.Title == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "title cannot be empty",
		})
		return
	}

	if blogUpdate.Content == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid request",
			Error:   "content cannot be empty",
		})
		return
	}

	claim, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Error:   "cannot get claims",
		})
		return
	}

	blog := &domain.Blog{
		Title:   blogUpdate.Title,
		Content: blogUpdate.Content,
		Tags:    blogUpdate.Tags,
	}

	newBlog, err := b.BlogUsecase.UpdateBlogByID(id.Hex(), blog, claim)
	if err != nil {
		code := config.GetStatusCode(err)
		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Cannot update blog",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Success",
		Data:    newBlog,
	})
}
