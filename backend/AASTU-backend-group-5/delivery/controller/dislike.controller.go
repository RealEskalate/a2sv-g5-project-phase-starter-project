package controller

import (
	"net/http"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/gin-gonic/gin"
)

type DislikeController struct {
	DislikeUseCase domain.DisLike_Usecase_interface
}

func NewDislikeController(dislikeUseCase domain.DisLike_Usecase_interface) *DislikeController {
	return &DislikeController{
		DislikeUseCase: dislikeUseCase,
	}
}

func (BC *DislikeController) GetDislikes(ctx *gin.Context) {
	postID := ctx.Param("post_id")
	dislikes, err := BC.DislikeUseCase.GetDislikes(postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, dislikes)
}

func (BC *DislikeController) CreateDislike(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	postID := ctx.Param("post_id")

	err := BC.DislikeUseCase.CreateDislike(userID, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "dislike added successfully"})
}

func (BC *DislikeController) ToggleDislike(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	postID := ctx.Param("post_id")

	err := BC.DislikeUseCase.ToggleDislike(userID, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "dislike toggled successfully"})
}

func (BC *DislikeController) RemoveDislike(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	postID := ctx.Param("post_id")

	err := BC.DislikeUseCase.RemoveDislike(userID, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "dislike removed successfully"})
}
