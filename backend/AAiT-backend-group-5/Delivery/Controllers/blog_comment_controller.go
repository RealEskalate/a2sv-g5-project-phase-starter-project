package controllers

import (
	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

type blogCommentController struct {
	usecase interfaces.BlogCommentUsecase
}

func NewBlogCommentController(usecase interfaces.BlogCommentUsecase) interfaces.BlogCommentController {
	return &blogCommentController{
		usecase: usecase,
	}
}

func (c *blogCommentController) getUserID(ctx *gin.Context) string {
	return ctx.GetString("id")
}

func (c *blogCommentController) getBlogID(ctx *gin.Context) string {
	return ctx.Param("blogID")
}

func (bcc *blogCommentController) AddCommentController(ctx *gin.Context) {
	blogID := bcc.getBlogID(ctx)
	userID := bcc.getUserID(ctx)

	var comment dtos.CommentCreateRequest
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newComment := models.Comment{
		Content: comment.Content,
		BlogID:  blogID,
		UserID:  userID,
	}

	err := bcc.usecase.AddComment(ctx, newComment)
	if err != nil {
		ctx.JSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": gin.H{"comment": "Comment added successfully"}})

}

func (bcc *blogCommentController) GetCommentsController(ctx *gin.Context) {
	blogID := bcc.getBlogID(ctx)

	comments, err := bcc.usecase.GetComments(ctx, blogID)
	if err != nil {
		ctx.JSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"comments": comments})
}

func (bcc *blogCommentController) GetCommentController(ctx *gin.Context) {
	commentID := ctx.Param("commentID")

	comment, err := bcc.usecase.GetComment(ctx, commentID)
	if err != nil {
		ctx.JSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"comment": comment})
}

func (bcc *blogCommentController) UpdateCommentController(ctx *gin.Context) {
	commentID := ctx.Param("commentID")
	userID := bcc.getUserID(ctx)

	var comment dtos.CommentUpdateRequest

	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := bcc.usecase.UpdateComment(ctx, commentID, userID, comment); err != nil {
		ctx.JSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": gin.H{"comment": "Comment updated successfully"}})

}

func (bcc *blogCommentController) DeleteCommentController(ctx *gin.Context) {
	// Implement the logic to delete a comment
	commentID := ctx.Param("commentID")
	userID := bcc.getUserID(ctx)

	if err := bcc.usecase.DeleteComment(ctx, userID, commentID); err != nil {
		ctx.JSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": gin.H{"comment": "Comment deleted successfully"}})

}