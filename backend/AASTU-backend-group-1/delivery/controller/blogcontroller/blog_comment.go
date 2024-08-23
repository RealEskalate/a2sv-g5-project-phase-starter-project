package blogcontroller

import (
	"blogs/config"
	"blogs/domain"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (l *BlogController) AddComment(ctx *gin.Context) {
	idHex := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Request",
			Error:   "The id is invalid",
		})
		return
	}

	var comment struct {
		Content string `bson:"content" json:"content"`
	}

	if err := ctx.ShouldBindJSON(&comment); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Request",
			Error:   err.Error(),
		})
		return
	}

	if comment.Content == "" {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Request",
			Error:   "Content cannot be empty",
		})
		return
	}

	claim, ok := ctx.MustGet("claims").(*domain.LoginClaims)
	if !ok {
		log.Println("Error getting claims")
		ctx.JSON(http.StatusInternalServerError, domain.APIResponse{
			Status:  http.StatusInternalServerError,
			Message: "Internal server error",
			Error:   "Error getting claims",
		})
		return
	}

	newcomment := domain.Comment{
		BlogID:  id,
		Author:  claim.Username,
		Content: comment.Content,
		Date:    time.Now(),
	}

	err = l.BlogUsecase.AddComment(&newcomment)
	if err != nil {
		code := config.GetStatusCode(err)
		log.Println(err)

		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Error adding comment",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusCreated,
		Message: "Comment added",
	})
}

func (l *BlogController) GetBlogComments(ctx *gin.Context) {
	idHex := ctx.Param("id")
	id, err := primitive.ObjectIDFromHex(idHex)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Request",
			Error:   "The id is invalid",
		})
		return
	}

	comments, err := l.BlogUsecase.GetBlogComments(id.Hex())
	if err != nil {
		code := config.GetStatusCode(err)
		log.Println(err)

		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Error getting comments",
			Error:   "Error getting comments",
		})
		return
	}

	if len(comments) == 0 {
		comments = []*domain.Comment{}
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status:  http.StatusOK,
		Message: "Comments retrieved",
		Data:    comments,
	})
}
