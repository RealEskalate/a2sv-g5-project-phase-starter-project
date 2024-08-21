package controllers

import (
	"net/http"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CommentController struct {
	CommentUsecase usecases.CommentUsecaseInterface
}

func (cont *CommentController) GetComments(c *gin.Context) {
	blogID, err := uuid.Parse(c.Param("blog_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	comments, cerr := cont.CommentUsecase.GetComments(blogID)
	if cerr != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"comments": comments})
}
func (cont *CommentController) GetCommentsCount(c *gin.Context) {
	blogID, err := uuid.Parse(c.Param("blog_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	commentCount, cerr := cont.CommentUsecase.GetCommentsCount(blogID)
	if cerr != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"comment count": commentCount})
}

func (cont *CommentController) AddComment(c *gin.Context) {
	var comment domain.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cerr := cont.CommentUsecase.AddComment(comment)
	if cerr != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": cerr.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Comment added successfully"})
}

func (cont *CommentController) UpdateComment(c *gin.Context) {
	commentID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	var comment domain.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requesterID := c.MustGet("id").(uuid.UUID)
	comment.UserID = requesterID
	comment.ID = commentID
	cerr := cont.CommentUsecase.UpdateComment(requesterID, comment)
	if cerr != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

func (cont *CommentController) DelelteComment(c *gin.Context) {
	commentID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
	}
	requesterID := c.MustGet("id").(uuid.UUID)
	requesterRole := c.MustGet("is_admin").(bool)
	if cerr := cont.CommentUsecase.DelelteComment(commentID, requesterID, requesterRole); err != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfuly"})
}
