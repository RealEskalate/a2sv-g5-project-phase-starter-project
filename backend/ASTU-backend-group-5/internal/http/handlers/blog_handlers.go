package handlers

import (
	"blogApp/internal/domain"
	"blogApp/internal/usecase/blog"
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	UseCase blog.BlogUseCase
}

func NewBlogHandler(useCase blog.BlogUseCase) *BlogHandler {
	return &BlogHandler{UseCase: useCase}
}

// CreateBlogHandler creates a new blog post
func (h *BlogHandler) CreateBlogHandler(c *gin.Context) {
	var blog domain.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UseCase.CreateBlog(context.Background(), &blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, blog)
}

// GetBlogByIDHandler retrieves a blog post by its ID
func (h *BlogHandler) GetBlogByIDHandler(c *gin.Context) {
	id := c.Param("id")
	blog, err := h.UseCase.GetBlogByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blog)
}

// UpdateBlogHandler updates an existing blog post
func (h *BlogHandler) UpdateBlogHandler(c *gin.Context) {
	id := c.Param("id")
	var blog domain.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UseCase.UpdateBlog(context.Background(), id, &blog); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blog)
}

// DeleteBlogHandler deletes a blog post by its ID
func (h *BlogHandler) DeleteBlogHandler(c *gin.Context) {
	id := c.Param("id")
	if err := h.UseCase.DeleteBlog(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetAllBlogsHandler retrieves all blog posts
func (h *BlogHandler) GetAllBlogsHandler(c *gin.Context) {
	blogs, err := h.UseCase.GetAllBlogs(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// FilterBlogsHandler filters blogs based on criteria
func (h *BlogHandler) FilterBlogsHandler(c *gin.Context) {
	var filter domain.BlogFilter
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	blogs, err := h.UseCase.FilterBlogs(context.Background(), filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

// PaginateBlogsHandler paginates blog posts
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

	c.JSON(http.StatusOK, blogs)
}

// AddTagToBlogHandler adds a tag to a blog post
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

// RemoveTagFromBlogHandler removes a tag from a blog post
func (h *BlogHandler) RemoveTagFromBlogHandler(c *gin.Context) {
	blogID := c.Param("id")
	tagID := c.Param("tagId")

	if err := h.UseCase.RemoveTagFromBlog(context.Background(), blogID, tagID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Tag removed successfully"})
}

// AddCommentHandler adds a comment to a blog post
func (h *BlogHandler) AddCommentHandler(c *gin.Context) {
	var comment domain.Comment
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UseCase.AddComment(context.Background(), &comment); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

// GetCommentsByBlogIDHandler retrieves comments by blog post ID
func (h *BlogHandler) GetCommentsByBlogIDHandler(c *gin.Context) {
	blogID := c.Param("id")

	comments, err := h.UseCase.GetCommentsByBlogID(context.Background(), blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comments)
}

// AddLikeHandler adds a like to a blog post
func (h *BlogHandler) AddLikeHandler(c *gin.Context) {
	var like domain.Like
	if err := c.ShouldBindJSON(&like); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UseCase.AddLike(context.Background(), &like); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, like)
}

// GetLikesByBlogIDHandler retrieves likes by blog post ID
func (h *BlogHandler) GetLikesByBlogIDHandler(c *gin.Context) {
	blogID := c.Param("id")

	likes, err := h.UseCase.GetLikesByBlogID(context.Background(), blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, likes)
}

// AddViewHandler adds a view to a blog post
func (h *BlogHandler) AddViewHandler(c *gin.Context) {
	var view domain.View
	if err := c.ShouldBindJSON(&view); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UseCase.AddView(context.Background(), &view); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, view)
}

// GetViewsByBlogIDHandler retrieves views by blog post ID
func (h *BlogHandler) GetViewsByBlogIDHandler(c *gin.Context) {
	blogID := c.Param("id")

	views, err := h.UseCase.GetViewsByBlogID(context.Background(), blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, views)
}

// Tags Operations
// GetAllTagsHandler retrieves all tags
func (h *BlogHandler) GetAllTagsHandler(c *gin.Context) {
	tags, err := h.UseCase.GetAllTags(context.Background())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tags)
}

// CreateTagHandler creates a new tag
func (h *BlogHandler) CreateTagHandler(c *gin.Context) {
	var tag domain.BlogTag
	if err := c.ShouldBindJSON(&tag); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UseCase.CreateTag(context.Background(), &tag); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, tag)
}

// UpdateTagHandler updates an existing tag
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

// DeleteTagHandler deletes a tag by its ID
func (h *BlogHandler) DeleteTagHandler(c *gin.Context) {
	id := c.Param("id")
	if err := h.UseCase.DeleteTag(context.Background(), id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}

// GetTagByIDHandler retrieves a tag by its ID
func (h *BlogHandler) GetTagByIDHandler(c *gin.Context) {
	id := c.Param("id")
	tag, err := h.UseCase.GetTagByID(context.Background(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tag)
}


