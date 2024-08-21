package gin

import (
	"errors"
	"net/http"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/auth"
	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/blog"
	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/pkg/infrastructure"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type BlogController struct {
	blogUseCase blog.BlogUseCase
}

func NewBlogController(blogUseCase blog.BlogUseCase) *BlogController {
	return &BlogController{
		blogUseCase: blogUseCase,
	}
}

func (bc *BlogController) CreateBlog(c *gin.Context) {
	var blog blog.CreateBlogRequest
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
			c.AbortWithStatusJSON(http.StatusInternalServerError, err)
		}
		return
	}

	c.JSON(http.StatusCreated, createdBlog)
}
