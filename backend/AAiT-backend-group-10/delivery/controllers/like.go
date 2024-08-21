package controllers

import (
	"net/http"

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	cerr := cont.LikeUseCase.LikeBlog(like)
	if cerr != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": cerr.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Like added successfully"})
}

func (cont *LikeController) DeleteLike(c *gin.Context) {
	var like domain.Like
	if err := c.BindJSON(&like); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	requester_id := c.MustGet("id").(uuid.UUID)
	like.UserID = requester_id
	cerr := cont.LikeUseCase.DeleteLike(like)
	if cerr != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": cerr.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Like deleted successfully"})
}

func (cont *LikeController) BlogLikeCount(c *gin.Context) {
	blogID, err := uuid.Parse(c.Param("blog_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}
	count, cerr := cont.LikeUseCase.BlogLikeCount(blogID)
	if cerr != nil {
		c.JSON(cerr.StatusCode, gin.H{"error": cerr.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"like count": count})
}
