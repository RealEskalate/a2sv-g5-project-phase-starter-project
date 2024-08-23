package controllers

import (
	"group3-blogApi/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LikeController struct {
	likeUsecase domain.LikeUsecase
}

func NewLikeController(likeUsecase domain.LikeUsecase) *LikeController {
	return &LikeController{
		likeUsecase: likeUsecase,
	}
}

func (c *LikeController) LikeBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	userId := ctx.GetString("user_id")
	Type := "like"
	err := c.likeUsecase.LikeBlog(userId, id, Type)
	if err.Message != "" {
		ctx.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Blog liked successfully",
	})
}

func (c *LikeController) DisLikeBlog(ctx *gin.Context) {
	id := ctx.Param("id")
	userId := ctx.GetString("user_id")
	Type := "dislike"

	err := c.likeUsecase.DisLikeBlog(userId, id, Type)
	if err.Message != "" {
		ctx.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "Blog disliked successfully",
	})
}
