package controllers

import (
	"meleket/domain"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LikeController struct {
	likeUsecase domain.LikeUsecaseInterface
}

func NewLikeController(likeUsecase domain.LikeUsecaseInterface) *LikeController {
	return &LikeController{likeUsecase: likeUsecase}
}

func (lc *LikeController) AddLike(c *gin.Context) {
	// Get the blog ID from the URL parameter
	blogIDParam := c.Param("id")
	blogID, err := primitive.ObjectIDFromHex(blogIDParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	// Get the user ID from the context (e.g., from JWT or middleware)
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID not found in context"})
		return
	}

	// Convert userID to primitive.ObjectID
	userObjectID, ok := userID.(primitive.ObjectID)
	if !ok {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Add the like using the blog ID and user ID
	err = lc.likeUsecase.AddLike(blogID, userObjectID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Like added successfully"})
}

func (lc *LikeController) RemoveLike(c *gin.Context) {
	likeID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid like ID"})
		return
	}

	err = lc.likeUsecase.RemoveLike(likeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Like removed successfully"})
}

func (lc *LikeController) GetLikesByBlogID(c *gin.Context) {
	blogID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid blog ID"})
		return
	}

	likes, err := lc.likeUsecase.GetLikesByBlogID(blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, likes)
}
