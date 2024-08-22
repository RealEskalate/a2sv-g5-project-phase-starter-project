package controller

import (
	"net/http"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/gin-gonic/gin"
)

type LikeController struct {
	LikeUseCase domain.Like_Usecase_interface
}

func NewLikeController(likeUseCase domain.Like_Usecase_interface) *LikeController {
	return &LikeController{
		LikeUseCase: likeUseCase,
	}
}

func (LC *LikeController) GetLikes(ctx *gin.Context) {
	postID := ctx.Param("post_id")
	likes, err := LC.LikeUseCase.GetLikes(postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, likes)
}

func (LC *LikeController) CreateLike(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	postID := ctx.Param("post_id")

	err := LC.LikeUseCase.CreateLike(userID, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "like added successfully"})
}

func (LC *LikeController) ToggleLike(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	postID := ctx.Param("post_id")

	err := LC.LikeUseCase.ToggleLike(userID, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "like toggled successfully"})
}

func (LC *LikeController) RemoveLike(ctx *gin.Context) {
	userID := ctx.Param("user_id")
	postID := ctx.Param("post_id")

	err := LC.LikeUseCase.RemoveLike(userID, postID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"status": "like removed successfully"})
}
