package gin

import (
	"errors"
	"log"
	"net/http"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/auth"
	blogDomain "github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/blog"
	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/pkg/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

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
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.Value("userID").(string)
	createdBlog, err := bc.blogUseCase.CreateBlog(c.Request.Context(), userID, blog)

	if err != nil {
		if (errors.As(err, validator.ValidationErrors{})) {
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

	userID := c.Value("userID").(string)
	blogID := c.Param("id")
	updatedBlog, err := bc.blogUseCase.UpdateBlog(c.Request.Context(), blogID, userID, blog)

	if err != nil {
		if (errors.As(err, validator.ValidationErrors{})) {
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
	userID := c.Value("userID").(string)

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
