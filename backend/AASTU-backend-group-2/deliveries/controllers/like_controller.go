package controllers

import (
	"blog_g2/domain"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LikeController struct {
	likeusecase domain.LikeUsecase
}

func NewLikeController(likeusecase domain.LikeUsecase) *LikeController {
	return &LikeController{likeusecase: likeusecase}
}

func (lc LikeController) CreateLike(c *gin.Context) {
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
	err := lc.likeusecase.CreateLike(context.TODO(), userIDString, postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to like the post",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post liked successfully",
	})
}

func (lc LikeController) DeleteLike(c *gin.Context) {

	role, exists := c.Get("role")
	if !exists || (role != "user" && role != "admin") {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	likeid, exists := c.Get("userID")
	likeidstr, ok := likeid.(string)
	if !ok || !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "like not found"})
		return
	}
	err := lc.likeusecase.DeleteLike(context.TODO(), likeidstr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to unlike the post",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Post unliked successfully",
	})
}

func (lc LikeController) GetLikes(c *gin.Context) {

	role, exists := c.Get("role")
	if !exists || (role != "user" && role != "admin") {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
		return
	}
	postID := c.Param("postID")
	likes, err := lc.likeusecase.GetLikes(context.TODO(), postID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to retrieve likes",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"likes": likes,
	})

}
