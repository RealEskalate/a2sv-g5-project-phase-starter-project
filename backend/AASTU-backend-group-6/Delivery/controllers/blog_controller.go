package controllers

import (
	domain "blogs/Domain"
	"net/http"
	"strings"
	"time"

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

// ReactOnBlog implements domain.BlogUsecase.
func (b BlogController) ReactOnBlog(c *gin.Context) {
	blog_id := c.Param("id")
	if blog_id == ":id" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Blog ID required.",
			Status:  http.StatusBadRequest,
		})
		return
	}
	reactionType := c.Query("isLiked")
	if strings.ToLower(reactionType) != "true" && strings.ToLower(reactionType) != "false" {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Valid reaction type required.",
			Status:  http.StatusBadRequest,
		})
		return
	}
	userID := c.GetString("user_id")
	if userID == "" {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{
			Message: "Authentication failed.",
			Status:  http.StatusUnauthorized,
		})
		return
	}
	err := b.BlogUsecase.ReactOnBlog(userID, reactionType, blog_id)
	if err != (domain.ErrorResponse{}) {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, domain.SuccessResponse{
		Message: "Reaction saved successfully",
		Status:  http.StatusOK,
	})
}

// CommentOnBlog implements domain.BlogUsecase.
func (b BlogController) CommentOnBlog(c *gin.Context) {
	var comment domain.Comment
	userID := c.GetString("user_id")
	if userID == ""{
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{
			Message: "Authentication failed.",
			Status:  http.StatusUnauthorized,
		})
		return
	}
	if err := c.ShouldBind(&comment); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusBadRequest,
		})
		return
	}
	if err := b.Validator.ValidateStruct(comment); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Message: "Invalid request payload.",
			Status:  http.StatusBadRequest,
		})
		return
	}
	err := b.BlogUsecase.CommentOnBlog(userID, comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
		return
	}
	c.JSON(http.StatusCreated, domain.SuccessResponse{
		Message: "Comment created successfully",
		Status:  http.StatusCreated,
	})
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
	if userID == "" || role == "" {
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{
			Message: "Authentication failed.",
			Status:  http.StatusUnauthorized,
		})
		return
	}
	err := b.BlogUsecase.DeleteBlogByID(userID, blogID)
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
	pageNo := c.Query("pageNo")
	pageSize := c.Query("pageSize")

	if pageNo == "" {
		pageNo = "1"
	}
	if pageSize == "" {
		pageSize = "1"
	}

	startDate := c.Query("startDate")
	endDate := c.Query("endDate")

	if startDate == "" || endDate == "" {
		if !(startDate == "" && endDate == "") {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "start and end date must be set together",
			})
			c.Abort()
		} else {
			startDate = time.Unix(0, 0).Format(time.RFC3339)
			endDate = time.Now().Format(time.RFC3339)
		}
	}

	tagsParam := c.Query("tags")
	var tags []string
	if tagsParam != "" {
		tags = strings.Split(tagsParam, ",")
		if len(tags) == 0 {
			c.JSON(http.StatusBadRequest, domain.ErrorResponse{
				Status:  http.StatusBadRequest,
				Message: "tags should not empty",
			})
			c.Abort()
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{
			Status:  http.StatusBadRequest,
			Message: "tags should not empty",
		})
		c.Abort()
		return
	}
	popularity := c.Query("popularity")
	blogs, pagination, err := b.BlogUsecase.FilterBlogsByTag(tags, pageNo, pageSize, startDate, endDate, popularity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Status:  http.StatusInternalServerError,
			Message: err.Error(),
		})
	} else {
		data := make(map[string]interface{}, 2)
		data["blogs"] = blogs
		data["pagination"] = pagination

		c.JSON(http.StatusAccepted, domain.SuccessResponse{
			Status:  http.StatusAccepted,
			Data:    data,
			Message: "blogs",
		})
	}
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
	blog, err := b.BlogUsecase.GetBlogByID(blog_id, false)
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
	popularity := c.Query("popularity")

	if pageNo == "" {
		pageNo = "1"
	}
	if pageSize == "" {
		pageSize = "1"
	}

	blogs, pagination, err := b.BlogUsecase.GetBlogs(pageNo, pageSize, popularity)
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

	blog, err := b.BlogUsecase.GetMyBlogByID(user_id, blog_id)
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
	popularity := c.Query("popularity")
	if pageNo == "" {
		pageNo = "1"
	}
	if pageSize == "" {
		pageSize = "1"
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

	myBlogs, pagination, err := b.BlogUsecase.GetMyBlogs(user_id, pageNo, pageSize, popularity)
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
		pageNo = "1"
	}
	if pageSize == "" {
		pageSize = "1"
	}

	popularity := c.Query("popularity")
	blogs, pagination, err := b.BlogUsecase.SearchBlogByTitleAndAuthor(title, author, pageNo, pageSize, popularity)

	if err != (domain.ErrorResponse{}) {
		c.JSON(err.Status, domain.ErrorResponse{
			Message: err.Message,
			Status:  err.Status,
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
	updatedBlog, err := b.BlogUsecase.UpdateBlogByID(user_id, blog_id, blog)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{
			Message: err.Error(),
			Status:  http.StatusInternalServerError,
		})
	} else {
		c.JSON(http.StatusAccepted, gin.H{"updated_blog": updatedBlog})
	}
}
