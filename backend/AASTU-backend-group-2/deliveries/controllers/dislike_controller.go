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
	postID := c.Param("postID")
	userID, exists := c.Get("userid")
	userIDString, ok := userID.(string)
	if !ok || !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "user not found"})
		return
	}
	err := dc.dislikeusecase.CreateDisLike(context.TODO(), userIDString, postID)
	if err != nil {
		c.JSON(err.Status(), gin.H{
			"error": err.Message(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Post disliked successfully",
	})
}

func (dc DisLikeController) DeleteDisLike(c *gin.Context) {

	dislikeid, exists := c.Get("userid")
	dislikeidstr, ok := dislikeid.(string)
	if !ok || !exists {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "like not found"})
		return
	}
	err := dc.dislikeusecase.DeleteDisLike(context.TODO(), dislikeidstr)
	if err != nil {
		c.JSON(err.Status(), gin.H{
			"error": err.Message(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Post undisliked successfully",
	})
}

func (dc DisLikeController) GetDisLikes(c *gin.Context) {
	postID := c.Param("postID")
	dislikes, err := dc.dislikeusecase.GetDisLikes(context.TODO(), postID)
	if err != nil {
		c.JSON(err.Status(), gin.H{
			"error": err.Message(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"dislikes": dislikes,
	})
}
