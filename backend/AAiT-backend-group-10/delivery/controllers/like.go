package controllers

import (
	"aait.backend.g10/domain"
	"aait.backend.g10/usecases"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LikeController struct {
	LikeUseCase usecases.LikeUsecaseInterface
}

func (cont *LikeController) LikeBlog(c *gin.Context) {
	var like domain.Like
	if err := c.BindJSON(&like); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := cont.LikeUseCase.LikeBlog(like)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Like added successfully"})
}

func (cont *LikeController) DeleteLike(c *gin.Context) {
	var like domain.Like
	if err := c.BindJSON(&like); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	err := cont.LikeUseCase.DeleteLike(like)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Like deleted successfully"})
}

func (cont *LikeController) BlogLikeCount(c *gin.Context) {
	blogID, err := uuid.Parse(c.Param("blog_id"))
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid blog ID"})
		return
	}
	count, err := cont.LikeUseCase.BlogLikeCount(blogID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"like count": count})
}
