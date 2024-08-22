package controllers

import (
	domain "aait-backend-group4/Domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

// LikeController handles the "like" and "dislike" functionalities for blog posts.
type LikeController struct {
	LikeUsecase domain.LikeUsecase
}

// Like adds a like to a blog post.
// It first checks the current like/dislike status of the user for the specified blog post.
// If the user has not interacted with the post, it adds a new like.
// If the user has already liked the post, it returns an error.
// If the user has disliked the post, it removes the dislike and then adds the like.
func (lctrl *LikeController) Like(c *gin.Context) {
	userID := c.Param("user_id")   // Extract user ID from the URL parameter
	blogID := c.Param("blog_id")   // Extract blog ID from the URL parameter

	// Check the current like/dislike status of the user for the specified blog
	likeStatus, likeID, err := lctrl.LikeUsecase.GetStatus(c, userID, blogID)
	if err != nil {
		if err.Error() == "status not found" {
			// If no status is found, it means the user has neither liked nor disliked the blog post
			err = lctrl.LikeUsecase.Like(c, userID, blogID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Like added successfully"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if likeStatus {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Already liked"})
		return
	}

	// If the user has disliked the blog post, remove the dislike
	err = lctrl.LikeUsecase.RemoveDislike(c, likeID, blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Add the like to the blog post
	err = lctrl.LikeUsecase.Like(c, userID, blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Like added successfully"})
}

// Dislike adds a dislike to a blog post.
// It first checks the current like/dislike status of the user for the specified blog post.
// If the user has not interacted with the post, it adds a new dislike.
// If the user has already disliked the post, it returns an error.
// If the user has liked the post, it removes the like and then adds the dislike.
func (lctrl *LikeController) Dislike(c *gin.Context) {
	userID := c.Param("user_id")   // Extract user ID from the URL parameter
	blogID := c.Param("blog_id")   // Extract blog ID from the URL parameter

	// Check the current like/dislike status of the user for the specified blog
	dislikeStatus, dislikeID, err := lctrl.LikeUsecase.GetStatus(c, userID, blogID)
	if err != nil {
		if err.Error() == "status not found" {
			// If no status is found, it means the user has neither liked nor disliked the blog post
			err = lctrl.LikeUsecase.Dislike(c, userID, blogID)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			c.JSON(http.StatusOK, gin.H{"message": "Dislike added successfully"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if dislikeStatus {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Already disliked"})
		return
	}

	// If the user has liked the blog post, remove the like
	err = lctrl.LikeUsecase.RemoveLike(c, dislikeID, blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Add the dislike to the blog post
	err = lctrl.LikeUsecase.Dislike(c, userID, blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dislike added successfully"})
}
