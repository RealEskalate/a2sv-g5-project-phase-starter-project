package controller

import (
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	"github.com/gin-gonic/gin"
)

// interface for blog controllers
type blogController interface {
	GetBlogs() gin.HandlerFunc
	GetBlog() gin.HandlerFunc
	CreateBlog() gin.HandlerFunc
	UpdateBlog() gin.HandlerFunc
	DeleteBlog() gin.HandlerFunc
	GetComments() gin.HandlerFunc
	CreateComment() gin.HandlerFunc
	GetComment() gin.HandlerFunc
	UpdateComment() gin.HandlerFunc
	DeleteComment() gin.HandlerFunc
	CreateLike() gin.HandlerFunc
}

type BlogController struct {
	BlogUsecase domain.BlogUsecase
	Env         *bootstrap.Env
}

func (bc *BlogController) GetBlogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogs, err := bc.BlogUsecase.GetAllBlogs()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, blogs)
	}
}

func (bc *BlogController) GetBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		blog, err := bc.BlogUsecase.GetBlogByID(blogID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, blog)
	}
}

func (bc *BlogController) CreateBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newBlog domain.Blog
		if err := c.ShouldBindJSON(&newBlog); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		blog, err := bc.BlogUsecase.CreateBlog(newBlog)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, blog)
	}
}

func (bc *BlogController) UpdateBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		var updatedBlog domain.Blog
		if err := c.ShouldBindJSON(&updatedBlog); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		blog, err := bc.BlogUsecase.UpdateBlog(blogID, updatedBlog)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, blog)
	}
}

func (bc *BlogController) DeleteBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		err := bc.BlogUsecase.DeleteBlog(blogID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(204, nil)
	}
}

func (bc *BlogController) GetComments() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		comments, err := bc.BlogUsecase.GetComments(blogID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, comments)
	}
}

func (bc *BlogController) CreateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		var newComment domain.Comment
		if err := c.ShouldBindJSON(&newComment); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		comment, err := bc.BlogUsecase.CreateComment(blogID, newComment)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, comment)
	}
}

func (bc *BlogController) GetComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		commentID := c.Param("comment_id")
		comment, err := bc.BlogUsecase.GetComment(blogID, commentID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, comment)
	}
}

func (bc *BlogController) UpdateComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		commentID := c.Param("comment_id")
		var updatedComment domain.Comment
		if err := c.ShouldBindJSON(&updatedComment); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		comment, err := bc.BlogUsecase.UpdateComment(blogID, commentID, updatedComment)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, comment)
	}
}

func (bc *BlogController) DeleteComment() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		commentID := c.Param("comment_id")
		err := bc.BlogUsecase.DeleteComment(blogID, commentID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(204, nil)
	}
}

func (bc *BlogController) CreateLike() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		userID := c.Query("user_id")
		err := bc.BlogUsecase.LikeBlog(blogID, userID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, nil)
	}
}
