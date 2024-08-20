package controllers

import (
	domain "aait-backend-group4/Domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type LikeController struct {
	likeUsecase domain.LikeUsecase
}

// NewLikeController creates a new instance of LikeController
func NewLikeController(router *gin.Engine, likeUsecase domain.LikeUsecase) {
	controller := &LikeController{
		likeUsecase: likeUsecase,
	}

	// Define the routes
	likes := router.Group("/likes")
	{
		likes.POST("/", controller.Like)
		likes.DELETE("/:likeID", controller.RemoveLike)
		likes.GET("/user/:userID", controller.GetLikesByUser)
		likes.GET("/blog/:blogID", controller.GetLikesByBlog)
		likes.POST("/dislike", controller.Dislike)
		likes.DELETE("/dislike/:dislikeID", controller.RemoveDislike)
	}
}

// Like handles the request to like a blog post
func (lc *LikeController) Like(c *gin.Context) {
	var request struct {
		UserID string `json:"user_id" binding:"required"`
		BlogID string `json:"blog_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := lc.likeUsecase.Like(c.Request.Context(), request.UserID, request.BlogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog liked successfully"})
}

// Dislike handles the request to dislike a blog post
func (lc *LikeController) Dislike(c *gin.Context) {
	var request struct {
		UserID string `json:"user_id" binding:"required"`
		BlogID string `json:"blog_id" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := lc.likeUsecase.Dislike(c.Request.Context(), request.UserID, request.BlogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog disliked successfully"})
}

// RemoveLike handles the request to remove a like from a blog post
func (lc *LikeController) RemoveLike(c *gin.Context) {
	likeID := c.Param("likeID")

	if likeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid like ID"})
		return
	}

	err := lc.likeUsecase.RemoveLike(c.Request.Context(), likeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Like removed successfully"})
}

// RemoveDislike handles the request to remove a dislike from a blog post
func (lc *LikeController) RemoveDislike(c *gin.Context) {
	dislikeID := c.Param("dislikeID")

	if dislikeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dislike ID"})
		return
	}

	err := lc.likeUsecase.RemoveDislike(c.Request.Context(), dislikeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dislike removed successfully"})
}

// GetLikesByUser handles the request to get likes by a specific user
func (lc *LikeController) GetLikesByUser(c *gin.Context) {
	userID := c.Param("userID")
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
		return
	}

	likes, err := lc.likeUsecase.GetLikesByUser(c.Request.Context(), userID, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, likes)
}

// GetLikesByBlog handles the request to get likes by a specific blog
func (lc *LikeController) GetLikesByBlog(c *gin.Context) {
	blogID := c.Param("blogID")
	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page"})
		return
	}

	likes, err := lc.likeUsecase.GetLikesByBlog(c.Request.Context(), blogID, limit, page)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, likes)
}
