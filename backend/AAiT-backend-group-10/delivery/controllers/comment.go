package controllers

import (
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
		c.JSON(400, gin.H{"error": "Invalid ID"})
	}
	comments, err := cont.CommentUsecase.GetComments(blogID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"comments": comments})
}
func (cont *CommentController) GetCommentsCount(c *gin.Context) {
	blogID, err := uuid.Parse(c.Param("blog_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
	}
	commentCount, err := cont.CommentUsecase.GetCommentsCount(blogID)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"comment count": commentCount})
}

func (cont *CommentController) AddComment(c *gin.Context) {
	var comment domain.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := cont.CommentUsecase.AddComment(comment)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(201, gin.H{"message": "Comment added successfully"})
}

func (cont *CommentController) UpdateComment(c *gin.Context) {
	commentID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
	}
	var comment domain.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	requesterID := c.MustGet("id").(uuid.UUID)
	comment.UserID = requesterID
	err = cont.CommentUsecase.UpdateComment(commentID, comment)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Comment updated successfully"})
}

func (cont *CommentController) DelelteComment(c *gin.Context) {
	commentID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid ID"})
	}
	requesterID := c.MustGet("id").(uuid.UUID)
	requesterRole := c.MustGet("is_admin").(bool)
	if err := cont.CommentUsecase.DelelteComment(commentID, requesterID, requesterRole); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Comment deleted successfuly"})
}
