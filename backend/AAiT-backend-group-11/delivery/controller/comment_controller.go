package controller

import (
	"net/http"

	"backend-starter-project/domain/entities"
	"backend-starter-project/domain/interfaces"
	"backend-starter-project/domain/dto"

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
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Failed to bind comment data",
			Error:   err.Error(),
		})
		return
	}

	comment.BlogPostID, _ = primitive.ObjectIDFromHex(c.Param("blogId"))
	comment.AuthorID, _ = primitive.ObjectIDFromHex(c.GetString("userId"))

	createdComment, err := cc.commentService.AddComment(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to add comment",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Comment added successfully",
		Data:    createdComment,
	})
}

func (cc *CommentController) DeleteComment(c *gin.Context) {
	commentId := c.Param("id")
	err := cc.commentService.DeleteComment(commentId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to delete comment",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Comment deleted successfully",
	})
}

func (cc *CommentController) GetCommentsByBlogPostId(c *gin.Context) {
	blogPostId := c.Param("blogId")

	comments, err := cc.commentService.GetCommentsByBlogPostId(blogPostId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to retrieve comments",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Comments retrieved successfully",
		Data:    comments,
	})
}

func (cc *CommentController) UpdateComment(c *gin.Context) {
	var comment entities.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: "Failed to bind comment data",
			Error:   err.Error(),
		})
		return
	}

	comment.ID, _ = primitive.ObjectIDFromHex(c.Param("id"))

	updatedComment, err := cc.commentService.UpdateComment(&comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: "Failed to update comment",
			Error:   err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "Comment updated successfully",
		Data:    updatedComment,
	})
}
