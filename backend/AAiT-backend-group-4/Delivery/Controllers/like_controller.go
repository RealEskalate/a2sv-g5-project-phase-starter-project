package controllers

import (
	domain "aait-backend-group4/Domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LikeController struct {
	LikeUsecase domain.LikeUsecase
}

// LikeBlog handles the request to like a blog post.
// - It retrieves the user ID and blog ID from the request context or parameters.
// - Calls the Like use case to increment the like count and update the blog's feedback.
// - Returns success or error responses based on the operation outcome.
func (lc *LikeController) LikeBlog(c *gin.Context) {
	blogID := c.Param("blog_id")
	userID := c.Request.Header.Get("userID")

	if err := lc.LikeUsecase.Like(c, userID, blogID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog liked successfully"})
}

// DislikeBlog handles the request to dislike a blog post.
// - It retrieves the user ID and blog ID from the request context or parameters.
// - Calls the Dislike use case to increment the dislike count and update the blog's feedback.
// - Returns success or error responses based on the operation outcome.
func (lc *LikeController) DislikeBlog(c *gin.Context) {
	blogID := c.Param("blog_id")
	userID := c.Request.Header.Get("userID")

	if err := lc.LikeUsecase.Dislike(c, userID, blogID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog disliked successfully"})
}

// RemoveLike handles the request to remove a like from a blog post.
// - It retrieves the like ID from the request parameters.
// - Calls the RemoveLike use case to decrement the like count and update the blog's feedback.
// - Returns success or error responses based on the operation outcome.
func (lc *LikeController) RemoveLike(c *gin.Context) {
	likeID := c.Param("like_id")

	if err := lc.LikeUsecase.RemoveLike(c, likeID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Like removed successfully"})
}

// RemoveDislike handles the request to remove a dislike from a blog post.
// - It retrieves the dislike ID from the request parameters.
// - Calls the RemoveDislike use case to decrement the dislike count and update the blog's feedback.
// - Returns success or error responses based on the operation outcome.
func (lc *LikeController) RemoveDislike(c *gin.Context) {
	dislikeID := c.Param("dislike_id")

	if err := lc.LikeUsecase.RemoveDislike(c, dislikeID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Dislike removed successfully"})
}

// GetLikesByUser handles the request to retrieve all likes by a user.
// - It retrieves the user ID from the request context or parameters.
// - Calls the GetLikesByUser use case to fetch all likes associated with the user.
// - Returns the list of likes or an error response.
func (lc *LikeController) GetLikesByUser(c *gin.Context) {
	userID := c.Param("user_id")

	likes, err := lc.LikeUsecase.GetLikesByUser(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, likes)
}

// GetLikesByBlog handles the request to retrieve all likes for a blog post.
// - It retrieves the blog ID from the request context or parameters.
// - Calls the GetLikesByBlog use case to fetch all likes associated with the blog post.
// - Returns the list of likes or an error response.
func (lc *LikeController) GetLikesByBlog(c *gin.Context) {
	blogID := c.Param("blog_id")

	likes, err := lc.LikeUsecase.GetLikesByBlog(c, blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, likes)
}
