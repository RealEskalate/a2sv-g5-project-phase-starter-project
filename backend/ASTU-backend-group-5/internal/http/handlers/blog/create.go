package blog

import (
	"blogApp/internal/ai"
	"blogApp/internal/config"
	"blogApp/internal/domain"
	"blogApp/internal/usecase/blog"
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	UseCase blog.BlogUseCase
}

func NewBlogHandler(useCase blog.BlogUseCase) *BlogHandler {
	return &BlogHandler{UseCase: useCase}
}

func (h *BlogHandler) CreateBlogHandler(c *gin.Context) {
	var blog domain.CreateBlogDTO
	claims, err := GetClaims(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// blog.Author = claims.UserID
	dbBlog := domain.Blog{
		Title:   blog.Title,
		Content: blog.Content,
		Tags:    blog.Tags,
	}
	pass := false
	Config, err := config.Load()
	if err == nil {
		if Config.MODERATE_BLOG_BEFORE_CREATE == "TRUE" {
			grade, message, err := ai.ModerateBlog(blog.Content, blog.Title)
			if err != nil {
				pass = true
			}
			if grade < 50 {
				c.JSON(http.StatusBadRequest, gin.H{"error": message})
				return
			}
			pass = true
		}
	}
	if pass {
		if err := h.UseCase.CreateBlog(context.Background(), &dbBlog, claims.UserID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, blog)
	}
}

func (h *BlogHandler) AddCommentHandler(c *gin.Context) {

	var comment domain.Comment

	userClaims, err := GetClaims(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.UseCase.AddComment(context.Background(), &comment, userClaims.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (h *BlogHandler) AddLikeHandler(c *gin.Context) {
	var like domain.Like

	userClaims, err := GetClaims(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&like); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UseCase.AddLike(context.Background(), &like, userClaims.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, like)
}

func (h *BlogHandler) AddViewHandler(c *gin.Context) {
	var view domain.View

	userClaims, err := GetClaims(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := c.ShouldBindJSON(&view); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.UseCase.AddView(context.Background(), &view, userClaims.UserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, view)
}

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
