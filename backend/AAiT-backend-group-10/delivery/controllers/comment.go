package controllers

import (
	"aait.backend.g10/domain"
	"aait.backend.g10/usecases/interfaces"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type CommentController struct {
	CommentUsecase interfaces.CommentUsecaseInterface
}

func (cont *CommentController) GetComments(c *gin.Context) {
	blogID := c.Param("blog_id")
	blog_id, _ := uuid.Parse(blogID)
	comments, err := cont.CommentUsecase.GetComments(blog_id)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"comments": comments})
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
	var comment domain.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := cont.CommentUsecase.UpdateComment(comment)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Comment updated successfully"})
}

func (cont *CommentController) DelelteComment(c *gin.Context) {
	blogID := c.Param("blog_id")
	userID := c.Param("user_id")
	blog_id, _ := uuid.Parse(blogID)
	user_id, _ := uuid.Parse(userID)
	if err := cont.CommentUsecase.DelelteComment(blog_id, user_id); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Comment deleted successfuly"})
}
