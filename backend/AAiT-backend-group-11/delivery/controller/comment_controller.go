package controllers

import (
	"net/http"

	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CommentController struct {
	commentService interfaces.CommentService
}

func NewCommentController(cs interfaces.CommentService) *CommentController {
	return &CommentController{
		commentService: cs,
	}
}

func (cc *CommentController) AddComment(c *gin.Context) {
	var comment entities.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.BlogPostID, _ = primitive.ObjectIDFromHex(c.Param("blogPostId"))
	comment.AuthorID, _ = primitive.ObjectIDFromHex(c.GetString("userId"))

	createdComment, err := cc.commentService.AddComment(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdComment)
}

func (cc *CommentController) DeleteComment(c *gin.Context) {
	commentId := c.Param("commentId")
	err := cc.commentService.DeleteComment(commentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

func (cc *CommentController) GetCommentsByBlogPostId(c *gin.Context) {
	blogPostId := c.Param("blogPostId")

	comments, err := cc.commentService.GetCommentsByBlogPostId(blogPostId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (cc *CommentController) UpdateComment(c *gin.Context) {
	var comment entities.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	comment.ID, _ = primitive.ObjectIDFromHex(c.Param("commentId"))

	updatedComment, err := cc.commentService.UpdateComment(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedComment)
}
