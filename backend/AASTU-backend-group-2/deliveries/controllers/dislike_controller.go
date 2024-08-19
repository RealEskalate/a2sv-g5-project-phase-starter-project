package controllers

import (
	"blog_g2/domain"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DisLikeController struct {
	dislikeusecase domain.DisLikeUsecase
}

func NewDisLikeController(dislikeusecase domain.DisLikeUsecase) *DisLikeController {
	return &DisLikeController{dislikeusecase: dislikeusecase}
}

func (dc DisLikeController) CreateDisLike(c *gin.Context) {
	role, exists := c.Get("role")
	if !exists || (role != "user" && role != "admin") {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	postID := c.Param("postID")
	userID, exists := c.Get("userID")
	userIDString, ok := userID.(string)
	if !ok || !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "user not found"})
		return
	}
	err := dc.dislikeusecase.CreateDisLike(context.TODO(), userIDString, postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to dislike the post",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post disliked successfully",
	})
}

func (dc DisLikeController) DeleteDisLike(c *gin.Context) {

	role, exists := c.Get("role")
	if !exists || (role != "user" && role != "admin") {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	dislikeid, exists := c.Get("userID")
	dislikeidstr, ok := dislikeid.(string)
	if !ok || !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "like not found"})
		return
	}
	err := dc.dislikeusecase.DeleteDisLike(context.TODO(), dislikeidstr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to undislike the post",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Post undisliked successfully",
	})
}

func (dc DisLikeController) GetDisLikes(c *gin.Context) {

	role, exists := c.Get("role")
	if !exists || (role != "user" && role != "admin") {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	postID := c.Param("postID")
	dislikes, err := dc.dislikeusecase.GetDisLikes(context.TODO(), postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve dislikes",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"dislikes": dislikes,
	})
}
