package controller

import (
	"Blog_Starter/domain"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
)

type BlogCommentController struct {
	blogCommentUsecase	domain.CommentUseCase
	ctx          		context.Context
}

func NewBlogCommentController(blogCommentUseCase domain.CommentUseCase, ctx context.Context) *BlogCommentController {
	return &BlogCommentController{
		blogCommentUsecase: blogCommentUseCase,
		ctx:           ctx,
	}
}

func (bc *BlogCommentController) CreateComment(c *gin.Context) {
	var createdComment domain.CommentRequest
	if err := c.BindJSON(&createdComment); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : "invalid request format"})
		return
	}

	insertedComment, err := bc.blogCommentUsecase.Create(bc.ctx, &createdComment)
	if err != nil {
		if err.Error() == "comment content too short" {
			c.IndentedJSON(http.StatusBadRequest,gin.H{"error" : err.Error()})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error" : "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"created_comment" : insertedComment})
}

func (bc *BlogCommentController) DeleteCommment(c *gin.Context) {
	commentId := c.Param("comment_id")
	deletedComment, err := bc.blogCommentUsecase.Delete(bc.ctx, commentId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error" : "internal server error"})
		return
	}
	
	c.IndentedJSON(http.StatusOK, gin.H{"deleted_comment" : deletedComment})
}

func (bc *BlogCommentController) UpdateComment(c *gin.Context) {
	var updatedComment domain.CommentRequest
	if err := c.BindJSON(&updatedComment); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : "invalid request format"})
		return
	}

	returnedComment, err := bc.blogCommentUsecase.Update(bc.ctx, updatedComment.Content, updatedComment.CommentID)
	if err != nil {
		if err.Error() == "content too short" {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : "comment content too short"})
			return
		}
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error" : "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"updated_comment" : returnedComment})
}