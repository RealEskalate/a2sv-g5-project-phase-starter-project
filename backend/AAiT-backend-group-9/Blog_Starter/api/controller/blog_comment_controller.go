package controller

import (
	"Blog_Starter/domain"
	"Blog_Starter/utils"
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)

type BlogCommentController struct {
	blogCommentUsecase	domain.CommentUseCase
	timeout				time.Duration
}

func NewBlogCommentController(blogCommentUseCase domain.CommentUseCase, timeout time.Duration) *BlogCommentController {
	return &BlogCommentController{
		blogCommentUsecase: blogCommentUseCase,
		timeout : timeout,
	}
}

func (bc *BlogCommentController) CreateComment(c *gin.Context) {
	var createdComment domain.CommentRequest
	blogID := c.Param("blog_id")
	if err := c.BindJSON(&createdComment); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : "invalid request format"})
		return
	}

	user, err := utils.CheckUser(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	createdComment.UserID = user.UserID
	createdComment.BlogID = blogID
	insertedComment, err := bc.blogCommentUsecase.Create(c, &createdComment)
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
	commentId := c.Param("id")
	deletedComment, err := bc.blogCommentUsecase.Delete(c, commentId)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error" : "internal server error"})
		return
	}
	
	c.IndentedJSON(http.StatusOK, gin.H{"deleted_comment" : deletedComment})
}

func (bc *BlogCommentController) UpdateComment(c *gin.Context) {
	commentID := c.Param("id")
	var updatedComment domain.CommentRequest
	if err := c.BindJSON(&updatedComment); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : "invalid request format"})
		return
	}

	returnedComment, err := bc.blogCommentUsecase.Update(c, updatedComment.Content, commentID)
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