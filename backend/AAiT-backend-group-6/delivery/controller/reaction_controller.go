package controller

import (
	"AAiT-backend-group-6/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ReactionController struct {
	ReactionUseCase domain.ReactionUsecase
}

func NewLikeController(reactionUseCase domain.ReactionUsecase) *ReactionController {
	return &ReactionController{
		ReactionUseCase: reactionUseCase,
	}
}

func (lc *ReactionController) LikeBlog(ctx *gin.Context) {
	blogID := ctx.Param("blog_id")

	// Get the user ID from the authentication context (adjust as per your auth system)
	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	blogObjectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}
	userObjectID, err := primitive.ObjectIDFromHex(userIDStr)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := lc.ReactionUseCase.LikeBlog(ctx, userObjectID, blogObjectID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Blog liked successfully"})
}

func (lc *ReactionController) UnLikeBlog(ctx *gin.Context) {
	blogID := ctx.Param("blog_id")

	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	blogObjectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}
	userObjectID, err := primitive.ObjectIDFromHex(userIDStr)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := lc.ReactionUseCase.UnLikeBlog(ctx, userObjectID, blogObjectID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Blog unliked successfully"})
}

func (lc *ReactionController) DeleteLike(ctx *gin.Context) {
	blogID := ctx.Param("blog_id")

	userID, exists := ctx.Get("user_id")
	if !exists {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	blogObjectID, err := primitive.ObjectIDFromHex(blogID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	userIDStr, ok := userID.(string)
	if !ok {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}
	userObjectID, err := primitive.ObjectIDFromHex(userIDStr)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID"})
		return
	}

	if err := lc.ReactionUseCase.DeleteLike(ctx, userObjectID, blogObjectID); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Like deleted successfully"})
}
