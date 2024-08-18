package controllers

import (
	"aait.backend.g10/domain"
	"aait.backend.g10/usecases"
	"github.com/gin-gonic/gin"
)

type LikeCOntroller struct {
	LikeUseCase usecases.LikeUsecaseInterface
}

func (cont *LikeCOntroller) LikeBlog(c *gin.Context) {
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

func (cont *LikeCOntroller) DeleteLike(c *gin.Context) {
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
