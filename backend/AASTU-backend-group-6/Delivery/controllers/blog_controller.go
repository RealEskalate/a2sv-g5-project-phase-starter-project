package controllers

import (
	domain "blogs/Domain"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	BlogUsecase domain.BlogUsecase
	Validator   domain.ValidateInterface
}

// CommentOnBlog implements domain.BlogUsecase.
func (b BlogController) CommentOnBlog(c *gin.Context) {
	panic("unimplemented")
}

// CreateBlog implements domain.BlogUsecase.
func (b BlogController) CreateBlog(c *gin.Context) {
	var blog domain.Blog
	uid, isAuth := c.Get("id")
	if isAuth {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{
			Message: "Authentication failed.",
			Status:  http.StatusUnauthorized,
		})
		return
	}
	userID, isString := uid.(string)
	if !isString {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{
			Message: "Authentication failed.",
			Status:  http.StatusUnauthorized,
		})
		return
	}
	if err := c.ShouldBind(&blog); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}
	if err := b.Validator.ValidateStruct(blog); err != nil {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{
			Message: "Invalid request payload.",
			Status:  http.StatusUnauthorized,
		})
		return
	}

	newBlog, err := b.BlogUsecase.CreateBlog(userID, blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusCreated, domain.SuccessResponse{
		Message: "Blog created successfully.",
		Data:    newBlog,
		Status:  http.StatusCreated,
	})

}

// DeleteBlogByID implements domain.BlogUsecase.
func (b BlogController) DeleteBlogByID(c *gin.Context) {
	panic("unimplemented")
}

// FilterBlogsByTag implements domain.BlogUsecase.
func (b BlogController) FilterBlogsByTag(c *gin.Context) {
	panic("unimplemented")
}

// GetBlogByID implements domain.BlogUsecase.
func (b BlogController) GetBlogByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "id is required"})
		c.Abort()
	}
	blog, err := b.BlogUsecase.GetBlogByID(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		c.Abort()
	} else {
		c.JSON(http.StatusOK, gin.H{"blog": blog})
	}
}

// GetBlogs implements domain.BlogUsecase.
func (b BlogController) GetBlogs(c *gin.Context) {
	pageNo := c.Query("pageNo")
	pageSize := c.Query("pageSize")

	if pageNo == "" {
		pageNo = "0"
	}
	if pageSize == "" {
		pageSize = "0"
	}

	blogs, pagination, err := b.BlogUsecase.GetBlogs(pageNo, pageSize)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		c.Abort()
	} else {
		c.JSON(http.StatusOK, gin.H{"blogs": blogs, "pagination": pagination})
	}
}

// GetMyBlogByID implements domain.BlogUsecase.
func (b BlogController) GetMyBlogByID(c *gin.Context) {
	panic("unimplemented")
}

// GetMyBlogs implements domain.BlogUsecase.
func (b BlogController) GetMyBlogs(c *gin.Context) {
	panic("unimplemented")
}

// SearchBlogByTitleAndAuthor implements domain.BlogUsecase.
func (b BlogController) SearchBlogByTitleAndAuthor(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	x := fmt.Sprintf("title: %s, author: %s", title, author)
	fmt.Println("////////////////////////////")
	fmt.Println(title, author)
	fmt.Println("////////////////////////////")
	c.JSON(200, gin.H{"des blogs": x})
}

// UpdateBlogByID implements domain.BlogUsecase.
func (b BlogController) UpdateBlogByID(c *gin.Context) {
	panic("unimplemented")
}

func NewBlogController(BlogUsecase domain.BlogUsecase, validator domain.ValidateInterface) BlogController {
	return BlogController{
		BlogUsecase: BlogUsecase,
		Validator:   validator,
	}
}
