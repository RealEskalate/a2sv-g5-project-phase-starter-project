package delivery

import (
	"AAiT-backend-group-6/domain"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	BlogUsecase domain.BlogUseCase
}

func NewBlogController(bu domain.BlogUseCase) *BlogController {
	return &BlogController{
		BlogUsecase: bu,
	}
}

func (bc *BlogController) CreateBlog(c *gin.Context) {
	var blog domain.Blog
	err := c.BindJSON(&blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
	}
	err = bc.BlogUsecase.CreateBlog(&blog)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Blog Created Successfully"})
}

func (bc *BlogController) GetBlog(c *gin.Context) {
	id := c.Param("id")
	blog, err := bc.BlogUsecase.GetBlog(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Blog Not Found"})
	}
	c.JSON(http.StatusOK, blog)
}

func (bc *BlogController) GetBlogs(c *gin.Context) {
	var pagination domain.Pagination
	err := c.BindQuery(&pagination)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
	}
	blogs, err := bc.BlogUsecase.GetBlogs(&pagination)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	c.JSON(http.StatusOK, blogs)
}

// update blog should be partially updated
func (bc *BlogController) UpdateBlog(c *gin.Context) {
	id := c.Param("id")
	var blog domain.Blog
	err := c.BindJSON(&blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
	}
	err = bc.BlogUsecase.UpdateBlog(&blog, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blog Updated Successfully"})
}

func (bc *BlogController) DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	err := bc.BlogUsecase.DeleteBlog(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blog Deleted Successfully"})
}

func (bc *BlogController) LikeBlog(c *gin.Context) {
	blogID := c.Param("blog_id")
	userID := c.Param("user_id")
	err := bc.BlogUsecase.LikeBlog(blogID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blog Liked Successfully"})
}

func (ctrl *BlogController) UnlikeBlog(c *gin.Context) {
	blogID := c.Param("blogID")
	userID := c.Param("userID")

	err := ctrl.BlogUsecase.UnlikeBlog(blogID, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Blog unliked successfully"})
}

// CommentBlog handles HTTP requests to add a comment to a blog post
func (ctrl *BlogController) CommentBlog(c *gin.Context) {
	blogID := c.Param("blogID")

	var comment domain.Comment
	if err := c.BindJSON(&comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := ctrl.BlogUsecase.CommentBlog(blogID, &comment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
}
