package controller

import (
	"net/http"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentController struct {
	CommentUsecase domain.Comment_Usecase_interface
}

func NewCommentController(commentUsecase domain.Comment_Usecase_interface) *CommentController {
	return &CommentController{
		CommentUsecase: commentUsecase,
	}
}

func (cc *CommentController) GetComments() gin.HandlerFunc {
	return func(c *gin.Context) {
		postID := c.Param("post_id")

		_, err := primitive.ObjectIDFromHex(postID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID format. Please provide a valid post ID."})
			return
		}

		comments, err := cc.CommentUsecase.GetComments(postID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve comments. Please try again: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Comments retrieved successfully!", "comments": comments})
	}
}

func (cc *CommentController) CreateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		postID := c.Param("post_id")

		_, err := primitive.ObjectIDFromHex(postID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID format. Please provide a valid post ID."})
			return
		}

		var requestBody struct {
			UserID string `json:"user_id"`
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body."})
			return
		}

		if err := cc.CommentUsecase.CreateComment(postID, requestBody.UserID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to add comment. Please try again: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Comment added successfully!"})
	}
}

func (cc *CommentController) UpdateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		commentID := c.Param("comment_id")

		_, err := primitive.ObjectIDFromHex(commentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID format. Please provide a valid comment ID."})
			return
		}

		if err := cc.CommentUsecase.UpdateComment(commentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update comment. Please try again: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully!"})
	}
}

func (cc *CommentController) DeleteComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		commentID := c.Param("comment_id")

		_, err := primitive.ObjectIDFromHex(commentID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid comment ID format. Please provide a valid comment ID."})
			return
		}

		if err := cc.CommentUsecase.DeleteComment(commentID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete comment. Please try again: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully!"})
	}
}
