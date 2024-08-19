package controller

import (
	"net/http"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type DislikeController struct {
	DislikeUsecase domain.DisLike_Usecase_interface
	UserUsecase    domain.User_Usecase_interface
}

func NewDislikeController(dislikeUsecase domain.DisLike_Usecase_interface, userUsecase domain.User_Usecase_interface) *DislikeController {
	return &DislikeController{
		DislikeUsecase: dislikeUsecase,
		UserUsecase:    userUsecase,
	}
}

func (dc *DislikeController) GetDisLikes() gin.HandlerFunc {
	return func(c *gin.Context) {
		postID := c.Param("post_id")

		_, err := primitive.ObjectIDFromHex(postID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID format. Please provide a valid post ID."})
			return
		}

		dislikes, err := dc.DislikeUsecase.GetDisLikes(postID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve dislikes. Please try again: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Dislikes retrieved successfully!", "dislikes": dislikes})
	}
}

func (dc *DislikeController) CreateDisLike() gin.HandlerFunc {
	return func(c *gin.Context) {
		postID := c.Param("post_id")

		_, err := primitive.ObjectIDFromHex(postID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid post ID format. Please provide a valid post ID."})
			return
		}

		claims, exists := c.Get("user")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized. Please log in to dislike this post."})
			return
		}

		userClaims := claims.(*domain.Claims)
		userID := userClaims.UserID

		if err := dc.DislikeUsecase.CreateDisLike(userID, postID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to dislike the post. Please try again: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Post disliked successfully!"})
	}
}

func (dc *DislikeController) DeleteDisLike() gin.HandlerFunc {
	return func(c *gin.Context) {
		dislikeID := c.Param("dislike_id")

		_, err := primitive.ObjectIDFromHex(dislikeID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid dislike ID format. Please provide a valid dislike ID."})
			return
		}

		if err := dc.DislikeUsecase.DeleteDisLike(dislikeID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete dislike. Please try again: " + err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "Dislike deleted successfully!"})
	}
}
