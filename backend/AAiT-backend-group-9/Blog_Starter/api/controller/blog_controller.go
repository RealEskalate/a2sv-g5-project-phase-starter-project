package controller

import (
	"Blog_Starter/domain"
	"Blog_Starter/utils"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

type BlogController struct {
	blogUseCase domain.BlogUseCase
}

func NewBlogController(blogUseCase domain.BlogUseCase) *BlogController {
	return &BlogController{
		blogUseCase: blogUseCase,
	}
}

// CreateBlog godoc
func (bc *BlogController) CreateBlog(c *gin.Context) {
    var blog domain.BlogCreate
    err := c.ShouldBindJSON(&blog)
    if err != nil {
        c.JSON(http.StatusBadRequest, domain.Response{
            Success: false,
            Message: err.Error(),
        })
        return
    }
  	err = blog.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
    user, err := utils.CheckUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, domain.Response{
            Success: false,
            Message: err.Error(),
        })
        return
    }
    blog.UserID = user.UserID


	blogModel, err := bc.blogUseCase.CreateBlog(c, &blog)
	if err != nil {
		if err.Error() == "content length should be greater than 10" {
			c.JSON(http.StatusBadRequest, domain.Response{
				Success: false,
				Message: err.Error(),
			})
		} else if err.Error() == "user not found" {
			c.JSON(http.StatusNotFound, domain.Response{
				Success: false,
				Message: err.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, domain.Response{
				Success: false,
				Message: err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusCreated, domain.Response{
		Success: true,
		Message: "Blog created successfully",
		Data:    blogModel,
	})
}

// GetBlogByID godoc
func (bc *BlogController) GetBlogByID(c *gin.Context) {
	blogID := c.Param("blog_id")
	blog, err := bc.blogUseCase.GetBlogByID(c, blogID)
	if err != nil {
		if err.Error() == "invalid blog id" {
			c.JSON(http.StatusBadRequest, domain.Response{
				Success: false,
				Message: err.Error(),
			})
		} else if err.Error() == "blog not found" {
			c.JSON(http.StatusNotFound, domain.Response{
				Success: false,
				Message: err.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, domain.Response{
				Success: false,
				Message: err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Success: true,
		Message: "Blog retrieved successfully",
		Data:    blog,
	})
}

// GetAllBlog godoc
func (bc *BlogController) GetAllBlog(c *gin.Context) {
	skipStr := c.Query("skip")
	limitStr := c.Query("limit")
	skip, _ := strconv.ParseInt(skipStr, 10, 64)
	limit, _ := strconv.ParseInt(limitStr, 10, 64)
	sortBy := c.Query("sort_by")

	blogs, paginationMetadata, err := bc.blogUseCase.GetAllBlog(c, skip, limit, sortBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Success: true,
		Message: "Blogs retrieved successfully",
		Data:    gin.H{"blogs": blogs, "metadata": paginationMetadata},
	})
}

// UpdateBlog godoc
func (bc *BlogController) UpdateBlog(c *gin.Context) {
	blogID := c.Param("blog_id")
	var blog domain.BlogUpdate
	err := c.ShouldBindJSON(&blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	err = blog.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	user, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}
	blog.UserID = user.UserID

	blogModel, err := bc.blogUseCase.UpdateBlog(c, &blog, blogID)
	if err != nil {
		if err.Error() == "blog not found" {
			c.JSON(http.StatusNotFound, domain.Response{
				Success: false,
				Message: err.Error(),
			})
		} else if err.Error() == "user is not the owner of the blog" {
			c.JSON(http.StatusUnauthorized, domain.Response{
				Success: false,
				Message: err.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, domain.Response{
				Success: false,
				Message: err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusOK, domain.Response{
		Success: true,
		Message: "Blog updated successfully",
		Data:    blogModel,
	})
}

// DeleteBlog godoc
func (bc *BlogController) DeleteBlog(c *gin.Context) {
	blogID := c.Param("blog_id")
	user, err := utils.CheckUser(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, domain.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err = bc.blogUseCase.DeleteBlog(c, blogID, user.UserID, user.Role)
	if err != nil {
		if err.Error() == "blog not found" {
			c.JSON(http.StatusNotFound, domain.Response{
				Success: false,
				Message: err.Error(),
			})
		} else if err.Error() == "user is not the owner of the blog" {
			c.JSON(http.StatusUnauthorized, domain.Response{
				Success: false,
				Message: err.Error(),
			})
		} else {
			c.JSON(http.StatusInternalServerError, domain.Response{
				Success: false,
				Message: err.Error(),
			})
		}
		return
	}
	c.JSON(http.StatusNoContent, domain.Response{
		Success: true,
		Message: "Blog deleted successfully",
	})
}

func (bc *BlogController) FilterBlog(c *gin.Context) {
	var filterReq domain.BlogFilterRequest
	if err := c.BindJSON(&filterReq); err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid request format",
		})
		return
	}

	filtrationResponse, err := bc.blogUseCase.FilterBlogs(c, &filterReq)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, domain.Response{
				Success: false,
				Message: "No matches found",
			})
			return
		}
		if err.Error() == "invalid request format" {
			c.JSON(http.StatusBadRequest, domain.Response{
				Success: false,
				Message: err.Error(),
			})
			return
		}
		c.JSON(http.StatusInternalServerError, domain.Response{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Success: true,
		Message: "Blogs filtered successfully",
		Data:    filtrationResponse,
	})
}

func (bc *BlogController) SearchBlog(c *gin.Context) {
	author := c.Query("author")
	title := c.Query("title")
	if title == "" && author == "" {
		c.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: "Invalid request format",
		})
		return
	}

	var searchRequest domain.BlogSearchRequest
	searchRequest.Author = author
	searchRequest.Title = title

	err := searchRequest.Validate()
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	searchResult, err := bc.blogUseCase.SearchBlogs(c, &searchRequest)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusNotFound, domain.Response{
				Success: false,
				Message: "No matches found",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, domain.Response{
			Success: false,
			Message: "Internal server error",
		})
		return
	}

	c.JSON(http.StatusOK, domain.Response{
		Success: true,
		Message: "Blogs searched successfully",
		Data:    searchResult,
	})
}
