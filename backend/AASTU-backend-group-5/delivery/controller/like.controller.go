package controller

import (
	"net/http"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LikeController struct {
	LikeUsecase domain.Like_Usecase_interface
	UserUsecase domain.User_Usecase_interface
}

func NewLikeController(likeUsecase domain.Like_Usecase_interface, userUsecase domain.User_Usecase_interface) *LikeController {
	return &LikeController{
		LikeUsecase: likeUsecase,
		UserUsecase: userUsecase,
	}
}

func (lc *LikeController) GetLikes() gin.HandlerFunc {
	return func(c *gin.Context) {
		postID := c.Param("post_id")

		_, err := primitive.ObjectIDFromHex(postID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID format. Please provide a valid post ID."})
			return
		}

		likes, err := lc.LikeUsecase.GetLikes(postID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve likes. Please try again: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Likes retrieved successfully!", "likes": likes})
	}
}

func (lc *LikeController) CreateLike() gin.HandlerFunc {
	return func(c *gin.Context) {
		postID := c.Param("post_id")

		_, err := primitive.ObjectIDFromHex(postID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID format. Please provide a valid post ID."})
			return
		}

		claims, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Please log in to like this post."})
			return
		}

		userClaims := claims.(*domain.Claims)
		userID := userClaims.UserID

		if err := lc.LikeUsecase.CreateLike(userID, postID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to like the post. Please try again: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Post liked successfully!"})
	}
}

func (lc *LikeController) DeleteLike() gin.HandlerFunc {
	return func(c *gin.Context) {
		likeID := c.Param("like_id")

		_, err := primitive.ObjectIDFromHex(likeID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid like ID format. Please provide a valid like ID."})
			return
		}

		if err := lc.LikeUsecase.DeleteLike(likeID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete like. Please try again: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Like deleted successfully!"})
	}
}
