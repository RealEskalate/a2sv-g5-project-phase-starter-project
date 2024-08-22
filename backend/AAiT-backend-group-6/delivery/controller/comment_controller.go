package controller

import (
	"AAiT-backend-group-6/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentController struct {
	CommentUseCase domain.CommentUseCase
}

// CreateComment creates a new comment for a blog
func (c *CommentController) CreateComment(ctx *gin.Context) {
	var comment domain.Comment
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Extract BlogID from URL parameter
	blogID := ctx.Param("blog_id")

	// Get UserID from authentication context
	userID, exists := ctx.Get("user_id") // Adjust based on your authentication method
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	// Set BlogID and Author field in the comment
	comment.BlogID, _ = primitive.ObjectIDFromHex(blogID)
	comment.Author.ID = userID.(primitive.ObjectID) // Assuming UserID is of type primitive.ObjectID

	if err := c.CommentUseCase.CreateComment(ctx, &comment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

// GetComment retrieves a comment by ID
func (c *CommentController) GetComment(ctx *gin.Context) {
	commentID := ctx.Param("comment_id")

	comment, err := c.CommentUseCase.GetComment(ctx, commentID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if comment == nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "Comment not found"})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

// UpdateComment updates a comment by ID
func (c *CommentController) UpdateComment(ctx *gin.Context) {
	var comment domain.Comment
	if err := ctx.ShouldBindJSON(&comment); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	commentID := ctx.Param("comment_id")
	comment.ID, _ = primitive.ObjectIDFromHex(commentID) // Convert string ID to ObjectID

	if err := c.CommentUseCase.UpdateComment(ctx, &comment); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, comment)
}

// DeleteComment deletes a comment by ID
func (c *CommentController) DeleteComment(ctx *gin.Context) {
	commentID := ctx.Param("comment_id")

	if err := c.CommentUseCase.DeleteComment(ctx, commentID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Comment deleted"})
}
