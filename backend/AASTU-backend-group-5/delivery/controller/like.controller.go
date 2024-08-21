package controller

import (
	"net/http"

	"github.com/RealEskalate/blogpost/usecase"
	"github.com/gin-gonic/gin"
)

type LikeController struct {
	LikeUseCase *usecase.LikeUseCase
}

func NewLikeController(likeUseCase *usecase.LikeUseCase) *LikeController {
	return &LikeController{
		LikeUseCase: likeUseCase,
	}
}

func (BC *LikeController) GetLikes(ctx *gin.Context) {
	post_id := ctx.Param("post_id")
	likes, err := BC.LikeUseCase.GetLikes(post_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, likes)
}

func (BC *LikeController) CreateLike(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	post_id := ctx.Param("post_id")

	err := BC.LikeUseCase.CreateLike(user_id, post_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "like added successfully"})
}

func (BC *LikeController) ToggleLike(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	post_id := ctx.Param("post_id")

	err := BC.LikeUseCase.ToggleLike(user_id, post_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "like toggled successfully"})
}

func (BC *LikeController) RemoveLike(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	post_id := ctx.Param("post_id")

	err := BC.LikeUseCase.RemoveLike(user_id, post_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "like removed successfully"})
}
