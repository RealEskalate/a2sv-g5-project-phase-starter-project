package gin

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/auth"
	blogDomain "github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/blog"
	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/pkg/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validationErr validator.ValidationErrors

type BlogController struct {
	blogUseCase blogDomain.BlogUseCase
}

func NewBlogController(blogUseCase blogDomain.BlogUseCase) *BlogController {
	return &BlogController{
		blogUseCase: blogUseCase,
	}
}

func (bc *BlogController) CreateBlog(c *gin.Context) {
	var blog blogDomain.CreateBlogRequest
	if err := c.ShouldBindJSON(&blog); err != nil {
		if err.Error() == "EOF" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON body"})
			return
		}
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.Value("user_id").(string)
	createdBlog, err := bc.blogUseCase.CreateBlog(c.Request.Context(), userID, blog)

	if err != nil {
		if errors.As(err, &validationErr) {
			c.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.ReturnErrorResponse(err))
		} else if errors.Is(err, auth.ErrNoUserWithId) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else {
			log.Default().Println("Error trying to create blog", err, "blog", blog)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusCreated, createdBlog)
}

func (bc *BlogController) UpdateBlog(c *gin.Context) {
	var blog blogDomain.UpdateBlogRequest
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.Value("user_id").(string)
	blogID := c.Param("id")
	updatedBlog, err := bc.blogUseCase.UpdateBlog(c.Request.Context(), blogID, userID, blog)

	if err != nil {
		if errors.As(err, &validationErr) {
			c.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.ReturnErrorResponse(err))
		} else if errors.Is(err, auth.ErrNoUserWithId) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else if errors.Is(err, blogDomain.ErrBlogNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else if errors.Is(err, blogDomain.ErrInvalidID) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else {
			log.Default().Println("Error trying to update blog", err, "blog", blog)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, updatedBlog)
}

func (bc *BlogController) DeleteBlog(c *gin.Context) {
	blogID := c.Param("id")
	userID := c.Value("user_id").(string)

	err := bc.blogUseCase.DeleteBlog(c.Request.Context(), userID, blogID)
	if err != nil {
		if errors.Is(err, auth.ErrNoUserWithId) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else if errors.Is(err, blogDomain.ErrBlogNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else {
			log.Default().Println("Error trying to delete blog", err, "ID:", blogID)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

func (bc *BlogController) GetBlogByID(c *gin.Context) {
	blogID := c.Param("id")

	blog, err := bc.blogUseCase.GetBlogByID(c.Request.Context(), blogID)
	if err != nil {
		if errors.Is(err, blogDomain.ErrBlogNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else {
			log.Default().Println("Error trying to get blog by ID", err, "ID:", blogID)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, blog)
}

func (bc *BlogController) GetBlogs(c *gin.Context) {
	var filterQuery blogDomain.FilterQuery
	popularity, err := strconv.ParseFloat(c.Query("popularity"), 32)
	filterQuery.CreatedAtFrom = c.Query("created_at_from")
	filterQuery.CreatedAtTo = c.Query("created_at_to")
	filterQuery.Tags = c.QueryArray("tags")
	filterQuery.Popularity = float32(popularity)

	var pagination infrastructure.PaginationRequest
	pagination.Limit, _ = strconv.Atoi(c.Query("limit"))
	pagination.Page, _ = strconv.Atoi(c.Query("page"))

	blogs, err := bc.blogUseCase.GetBlogs(c.Request.Context(), filterQuery, pagination)
	if err != nil {
		log.Default().Println("Error trying to get blogs", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

func (bc *BlogController) SearchBlogs(c *gin.Context) {
	query := c.Query("query")

	var pagination infrastructure.PaginationRequest
	pagination.Limit, _ = strconv.Atoi(c.Query("limit"))
	pagination.Page, _ = strconv.Atoi(c.Query("page"))

	blogs, err := bc.blogUseCase.SearchBlogs(c.Request.Context(), query, pagination)
	if err != nil {
		log.Default().Println("Error trying to search blogs", err)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.JSON(http.StatusOK, blogs)
}

func (bc *BlogController) GetCommentsByBlogID(c *gin.Context) {
	blogID := c.Param("id")
	var pagination infrastructure.PaginationRequest
	pagination.Limit, _ = strconv.Atoi(c.Query("limit"))
	pagination.Page, _ = strconv.Atoi(c.Query("page"))

	comments, err := bc.blogUseCase.GetCommentsByBlogID(c.Request.Context(), blogID, pagination)
	if err != nil {
		if errors.Is(err, blogDomain.ErrBlogNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else {
			log.Default().Println("Error trying to get comments by blog ID", err, "ID:", blogID)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.JSON(http.StatusOK, comments)
}

func (bc *BlogController) CreateComment(c *gin.Context) {
	userID := c.Value("user_id").(string)
	blogID := c.Param("id")

	var comment blogDomain.CreateCommentRequest
	if err := c.ShouldBindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := bc.blogUseCase.CreateComment(c.Request.Context(), userID, blogID, comment)
	if err != nil {
		if errors.As(err, &validationErr) {
			c.AbortWithStatusJSON(http.StatusBadRequest, infrastructure.ReturnErrorResponse(err))
		} else if errors.Is(err, auth.ErrNoUserWithId) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else if errors.Is(err, blogDomain.ErrBlogNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else {
			log.Default().Println("Error trying to create comment", err, "comment", comment)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.Status(http.StatusCreated)
}

func (bc *BlogController) DeleteComment(c *gin.Context) {
	commentID := c.Param("comment_id")
	userID := c.Value("user_id").(string)

	err := bc.blogUseCase.DeleteComment(c.Request.Context(), commentID, userID)
	if err != nil {
		if errors.Is(err, auth.ErrNoUserWithId) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else if errors.Is(err, blogDomain.ErrCommentNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else {
			log.Default().Println("Error trying to delete comment", err, "ID:", commentID)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

func (bc *BlogController) LikeBlog(c *gin.Context) {
	userID := c.Value("user_id").(string)
	blogID := c.Param("id")

	err := bc.blogUseCase.LikeBlog(c.Request.Context(), userID, blogID)
	if err != nil {
		if errors.Is(err, auth.ErrNoUserWithId) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else if errors.Is(err, blogDomain.ErrBlogNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else {
			log.Default().Println("Error trying to like blog", err, "ID:", blogID)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

func (bc *BlogController) DislikeBlog(c *gin.Context) {
	userID := c.Value("user_id").(string)
	blogID := c.Param("id")

	err := bc.blogUseCase.DislikeBlog(c.Request.Context(), userID, blogID)
	if err != nil {
		if errors.Is(err, auth.ErrNoUserWithId) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else if errors.Is(err, blogDomain.ErrBlogNotFound) {
			c.AbortWithStatusJSON(http.StatusNotFound, err)
		} else {
			log.Default().Println("Error trying to dislike blog", err, "ID:", blogID)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		}
		return
	}

	c.Status(http.StatusNoContent)
}

func (bc *BlogController) UnLikeBlog(c *gin.Context) {
	userID := c.Value("user_id").(string)
	blogID := c.Param("id")

	err := bc.blogUseCase.UnLikeBlog(c.Request.Context(), userID, blogID)
	if err != nil {
		log.Default().Println("Error trying to unlike blog", err, "ID:", blogID)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Status(http.StatusNoContent)
}

func (bc *BlogController) UnDislikeBlog(c *gin.Context) {
	userID := c.Value("user_id").(string)
	blogID := c.Param("id")

	err := bc.blogUseCase.UnDislikeBlog(c.Request.Context(), userID, blogID)
	if err != nil {
		log.Default().Println("Error trying to undislike blog", err, "ID:", blogID)
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}

	c.Status(http.StatusNoContent)
}
