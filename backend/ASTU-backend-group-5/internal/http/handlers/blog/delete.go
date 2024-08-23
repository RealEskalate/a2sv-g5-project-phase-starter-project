package blog

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *BlogHandler) DeleteBlogHandler(c *gin.Context) {
	id := c.Param("id")
	userClaims, err := GetClaims(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := h.UseCase.DeleteBlog(context.Background(), id, userClaims.UserID, userClaims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, gin.H{"message": "Blog deleted successfully"})
}

func (h *BlogHandler) DeleteCommentHandler(c *gin.Context) {

	userClaims, err := GetClaims(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	commentId := c.Param("id")
	if err := h.UseCase.DeleteComment(context.Background(), commentId, userClaims.UserID, userClaims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}

func (h *BlogHandler) DeleteLikeHandler(c *gin.Context) {
	userClaims, err := GetClaims(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	likeId := c.Param("id")

	if err := h.UseCase.RemoveLike(context.Background(), likeId, userClaims.UserID, userClaims.Role); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Like deleted successfully"})
}

func (h *BlogHandler) DeleteTagHandler(c *gin.Context) {
	id := c.Param("id")
	if err := h.UseCase.DeleteTag(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
