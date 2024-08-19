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

func NewBlogController(BlogUsecase domain.BlogUsecase, validator domain.ValidateInterface) BlogController {
	return BlogController{
		BlogUsecase: BlogUsecase,
		Validator:   validator,
	}
}

// CommentOnBlog implements domain.BlogUsecase.
func (b BlogController) CommentOnBlog(c *gin.Context) {
	panic("unimplemented")
}

// CreateBlog implements domain.BlogUsecase.
func (b BlogController) CreateBlog(c *gin.Context) {
	var blog domain.Blog
	userID := c.GetString("user_id")
	role := c.GetString("role")
	if userID == "" || role == "" {
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
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid request payload.",
			Status:  http.StatusBadRequest,
		})
		return
	}

	newBlog, err := b.BlogUsecase.CreateBlog(userID, blog, role)
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
	blogID := c.Param("id")
	if blogID == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Blog id required.",
			Status:  http.StatusBadRequest,
		})
		return
	}
	userID := c.GetString("user_id")
	role := c.GetString("role")
	fmt.Println(userID, role)
	if userID == "" || role == "" {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{
			Message: "Authentication failed.",
			Status:  http.StatusUnauthorized,
		})
		return
	}
	err := b.BlogUsecase.DeleteBlogByID(userID, blogID, role)
	if err != (domain.ErrorResponse{}) {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Blog deleted successfully.",
		Status:  http.StatusOK,
	})

}

// FilterBlogsByTag implements domain.BlogUsecase.
func (b BlogController) FilterBlogsByTag(c *gin.Context) {
	panic("unimplemented")
}

// GetBlogByID implements domain.BlogUsecase.
func (b BlogController) GetBlogByID(c *gin.Context) {
	blog_id := c.Param("id")
	if blog_id == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "user id is required",
			Status:  http.StatusBadRequest,
		})
		c.Abort()
	}

	blog, err := b.BlogUsecase.GetBlogByID(blog_id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		c.Abort()
	} else {
		c.JSON(http.StatusOK, domain.SuccessResponse{
			Status:  http.StatusOK,
			Data:    blog,
			Message: "blog",
		})
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
	blog_id := c.Param("id")
	if blog_id == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "blog id is required",
			Status:  http.StatusBadRequest,
		})
		c.Abort()
	}
	user_id := c.GetString("user_id")
	role := c.GetString("role")
	if user_id == "" || role == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "User id is required",
			Status:  http.StatusBadRequest,
		})
	}

	blog, err := b.BlogUsecase.GetMyBlogByID(user_id, blog_id, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		c.Abort()
	} else {
		c.JSON(http.StatusOK, domain.SuccessResponse{
			Status:  http.StatusOK,
			Data:    blog,
			Message: "blog",
		})
	}
}

// GetMyBlogs implements domain.BlogUsecase.
func (b BlogController) GetMyBlogs(c *gin.Context) {
	pageNo := c.Query("pageNo")
	pageSize := c.Query("pageSize")
	if pageNo == "" {
		pageNo = "0"
	}
	if pageSize == "" {
		pageSize = "0"
	}

	// user_id, user_id_existes := c.Get("id")
	user_id := c.GetString("user_id")
	// role := c.GetString("role")
	if user_id == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "User id is required",
			Status:  http.StatusBadRequest,
		})
	}

	myBlogs, pagination, err := b.BlogUsecase.GetMyBlogs(user_id, pageNo, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		c.Abort()
	} else {
		c.JSON(http.StatusOK, gin.H{"my_blogs": myBlogs, "pagination": pagination})
	}
}

// SearchBlogByTitleAndAuthor implements domain.BlogUsecase.
func (b BlogController) SearchBlogByTitleAndAuthor(c *gin.Context) {
	title := c.Query("title")
	author := c.Query("author")
	pageNo := c.Query("pageNo")
	pageSize := c.Query("pageSize")
	if pageNo == "" {
		pageNo = "0"
	}
	if pageSize == "" {
		pageSize = "0"
	}
	blogs, pagination, err := b.BlogUsecase.SearchBlogByTitleAndAuthor(title, author, pageNo, pageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Blogs fetched successfully.",
		Data: map[string]interface{}{
			"Blogs":      blogs,
			"Pagination": pagination,
		},
		Status: http.StatusOK,
	})

}

// UpdateBlogByID implements domain.BlogUsecase.
func (b BlogController) UpdateBlogByID(c *gin.Context) {

	blog_id := c.Param("id")
	if blog_id == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "blog id is required",
			Status:  http.StatusBadRequest,
		})
		c.Abort()
	}
	user_id := c.GetString("user_id")
	role := c.GetString("role")
	if user_id == "" || role == "" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "User id is required",
			Status:  http.StatusBadRequest,
		})
	}

	var blog domain.Blog
	if err := c.ShouldBind(&blog); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		c.Abort()
	}
	updatedBlog, err := b.BlogUsecase.UpdateBlogByID(user_id, blog_id, blog, role)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
	} else {
		c.JSON(http.StatusAccepted, gin.H{"updated_blog": updatedBlog})
	}
}
