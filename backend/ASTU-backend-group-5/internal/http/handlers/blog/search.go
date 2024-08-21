package blog

import (
	"blogApp/internal/domain"
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (h *BlogHandler) SearchBlogsHandler(c *gin.Context) {
	var filter domain.BlogFilter

	keyword := c.Query("keyword")
	if keyword != "" {
		filter.Keyword = &keyword
	}

	authorID := c.Query("author")
	if authorID != "" {
		id, err := primitive.ObjectIDFromHex(authorID)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid author ID"})
			return
		}
		filter.AuthorID = &id
	}

	tags := c.QueryArray("tags")
	if len(tags) > 0 {
		filter.Tags = tags
	}

	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	pageSize, _ := strconv.Atoi(c.DefaultQuery("pageSize", "10"))

	blogs, totalBlogs, err := h.UseCase.SearchBlogs(context.Background(), filter, page, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalPages := (totalBlogs + pageSize - 1) / pageSize
	var nextPageURL string
	if page < totalPages {
		nextPageURL = fmt.Sprintf("%s?keyword=%s&author=%s&tags=%s&page=%d&pageSize=%d",
			c.Request.URL.Path, keyword, authorID, strings.Join(tags, ","), page+1, pageSize)
	}

	// Return the results
	c.JSON(http.StatusOK, gin.H{
		"blogs":       blogs,
		"currentPage": page,
		"pageSize":    pageSize,
		"totalPages":  totalPages,
		"totalBlogs":  totalBlogs,
		"nextPageURL": nextPageURL,
	})
}
