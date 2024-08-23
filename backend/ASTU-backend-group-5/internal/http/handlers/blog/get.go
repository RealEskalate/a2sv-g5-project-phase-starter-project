package blog

import (
	"blogApp/internal/domain"
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func (h *BlogHandler) GetBlogByIDHandler(c *gin.Context) {
	claims, _ := GetClaims(c)
	userId := ""
	if claims != nil {
		userId = claims.UserID
	}

	blogId := c.Param("id")
	blog, err := h.UseCase.GetBlogByID(context.Background(), blogId, userId)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	if blog == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		return
	}

	c.JSON(http.StatusOK, blog)
}

func (h *BlogHandler) GetAllBlogsHandler(c *gin.Context) {
	blogs, err := h.UseCase.GetAllBlogs(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// func (h *BlogHandler) FilterBlogsHandler(c *gin.Context) {
// 	var filter domain.BlogFilter
// 	if err := c.ShouldBindQuery(&filter); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	blogs, err := h.UseCase.FilterBlogs(context.Background(), filter)
// 	if err != nil {
// 		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
// 		return
// 	}

// 	c.JSON(http.StatusOK, blogs)
// }

func (h *BlogHandler) PaginateBlogsHandler(c *gin.Context) {
	var filter domain.BlogFilter
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blogs, err := h.UseCase.PaginateBlogs(context.Background(), filter, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	nextPageURL := fmt.Sprintf("%s?page=%d&pageSize=%d", c.Request.URL.Path, page+1, pageSize)
	previousPageURL := ""
	if page > 1 {
		previousPageURL = fmt.Sprintf("%s?page=%d&pageSize=%d", c.Request.URL.Path, page, pageSize)
	}
	c.JSON(http.StatusOK, gin.H{"blogs": blogs, "nextPageURL": nextPageURL, "previousPageURL": previousPageURL})
}

func (h *BlogHandler) GetCommentsByBlogIDHandler(c *gin.Context) {
	blogID := c.Param("id")

	comments, err := h.UseCase.GetCommentsByBlogID(context.Background(), blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (h *BlogHandler) GetLikesByBlogIDHandler(c *gin.Context) {
	blogID := c.Param("id")

	likes, err := h.UseCase.GetLikesByBlogID(context.Background(), blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, likes)
}

func (h *BlogHandler) GetViewsByBlogIDHandler(c *gin.Context) {
	blogID := c.Param("id")

	views, err := h.UseCase.GetViewsByBlogID(context.Background(), blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, views)
}

func (h *BlogHandler) GetAllTagsHandler(c *gin.Context) {
	tags, err := h.UseCase.GetAllTags(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tags)
}

func (h *BlogHandler) GetTagByIDHandler(c *gin.Context) {
	id := c.Param("id")
	tag, err := h.UseCase.GetTagByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tag)
}
