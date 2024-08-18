package controllers

import (
	"meleket/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// CommentController handles comment-related requests
type CommentController struct {
	commentUsecase domain.CommentUsecaseInterface
}

// NewCommentController creates a new instance of CommentController
func NewCommentController(commentUsecase domain.CommentUsecaseInterface) *CommentController {
	return &CommentController{commentUsecase: commentUsecase}
}

// AddComment adds a new comment to a blog post
func (cc *CommentController) AddComment(c *gin.Context) {
	var comment domain.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blogID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	comment.BlogID = blogID
	comment.UserID = c.MustGet("userID").(primitive.ObjectID)

	if err := cc.commentUsecase.AddComment(&comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Comment added successfully", "comment": comment})
}

// GetCommentsByBlogID retrieves all comments for a specific blog post
func (cc *CommentController) GetCommentsByBlogID(c *gin.Context) {
	blogID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	comments, err := cc.commentUsecase.GetCommentsByBlogID(blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// UpdateComment updates an existing comment
func (cc *CommentController) UpdateComment(c *gin.Context) {
	commentID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	var content struct {
		Content string `json:"content"`
	}
	if err := c.ShouldBindJSON(&content); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := cc.commentUsecase.UpdateComment(commentID, content.Content); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

// DeleteComment deletes a comment by its ID
func (cc *CommentController) DeleteComment(c *gin.Context) {
	commentID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID"})
		return
	}

	if err := cc.commentUsecase.DeleteComment(commentID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
