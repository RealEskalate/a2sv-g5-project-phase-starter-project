package controllers

import (
	"net/http"

	"aait.backend.g10/domain"
	"aait.backend.g10/usecases"
	"aait.backend.g10/usecases/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type LikeController struct {
	LikeUseCase usecases.LikeUsecaseInterface
}

func (cont *LikeController) LikeBlog(c *gin.Context) {
	var like domain.Like
	if err := c.BindJSON(&like); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	reacterId, err := uuid.Parse(c.MustGet("id").(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}
	like.ReacterID = reacterId
	cerr := cont.LikeUseCase.LikeBlog(like)
	if cerr != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": cerr.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Like added successfully"})
}

func (cont *LikeController) DeleteLike(c *gin.Context) {
	var like dto.UnlikeDto
	if err := c.BindJSON(&like); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requester_id, err := uuid.Parse(c.MustGet("id").(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return 
	}
	like.ReacterID = requester_id
	cerr := cont.LikeUseCase.DeleteLike(like)
	if cerr != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": cerr.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Like deleted successfully"})
}