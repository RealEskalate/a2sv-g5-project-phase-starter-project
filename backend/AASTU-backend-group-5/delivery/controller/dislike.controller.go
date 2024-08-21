package controller

import (
	"net/http"

	"github.com/RealEskalate/blogpost/usecase"
	"github.com/gin-gonic/gin"
)

type DislikeController struct {
	DislikeUseCase *usecase.DislikeUseCase
}

func NewDislikeController(dislikeUseCase *usecase.DislikeUseCase) *DislikeController {
	return &DislikeController{
		DislikeUseCase: dislikeUseCase,
	}
}

func (BC *DislikeController) GetDislikes(ctx *gin.Context) {
	post_id := ctx.Param("post_id")
	dislikes, err := BC.DislikeUseCase.GetDislikes(post_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, dislikes)
}

func (BC *DislikeController) CreateDislike(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	post_id := ctx.Param("post_id")

	err := BC.DislikeUseCase.CreateDislike(user_id, post_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "dislike added successfully"})
}

func (BC *DislikeController) ToggleDislike(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	post_id := ctx.Param("post_id")

	err := BC.DislikeUseCase.ToggleDislike(user_id, post_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "dislike toggled successfully"})
}

func (BC *DislikeController) RemoveDislike(ctx *gin.Context) {
	user_id := ctx.Param("user_id")
	post_id := ctx.Param("post_id")

	err := BC.DislikeUseCase.RemoveDislike(user_id, post_id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "dislike removed successfully"})
}
