package blog

import (
	"blogApp/internal/domain"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *BlogHandler) UpdateBlogHandler(c *gin.Context) {
	id := c.Param("id")

	userClaims, err := GetClaims(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var updateBlog domain.UpdateBlogDTO
	if err := c.ShouldBindJSON(&updateBlog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var blog domain.Blog
	blog.Title = updateBlog.Title
	blog.Content = updateBlog.Content
	blog.Tags = updateBlog.Tags
	blog.AuthorName = updateBlog.AuthorName

	if err := h.UseCase.UpdateBlog(context.Background(), id, &blog, userClaims.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blog)
}

func (h *BlogHandler) AddTagToBlogHandler(c *gin.Context) {
	blogID := c.Param("id")
	var tag domain.BlogTag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UseCase.AddTagToBlog(context.Background(), blogID, tag); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tag added successfully"})
}

func (h *BlogHandler) RemoveTagFromBlogHandler(c *gin.Context) {
	blogID := c.Param("id")
	tagID := c.Param("tagId")

	if err := h.UseCase.RemoveTagFromBlog(context.Background(), blogID, tagID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tag removed successfully"})
}

func (h *BlogHandler) UpdateTagHandler(c *gin.Context) {
	id := c.Param("id")
	var tag domain.BlogTag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UseCase.UpdateTag(context.Background(), id, &tag); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tag)
}
