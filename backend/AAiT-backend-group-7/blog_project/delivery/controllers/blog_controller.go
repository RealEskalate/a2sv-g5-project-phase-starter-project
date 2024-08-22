package controllers

import (
	"blog_project/domain"
	"strconv"

	"github.com/gin-gonic/gin"
)

type blogController struct {
	BlogUsecase domain.IBlogUsecase
}

func NewBlogController(blogUsecase domain.IBlogUsecase) domain.IBlogController {
	return &blogController{BlogUsecase: blogUsecase}
}

func (bc *blogController) GetAllBlogs(c *gin.Context) {
	sortOrder := c.DefaultQuery("sort", "DESC") // Default to "DESC" if not specified
	page := c.DefaultQuery("page", "1")
	limit := c.DefaultQuery("limit", "10")

	// Convert page and limit to integers
	pageInt, err := strconv.Atoi(page)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid page number"})
		return
	}

	limitInt, err := strconv.Atoi(limit)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid limit number"})
		return
	}

	blogs, err := bc.BlogUsecase.GetAllBlogs(c, sortOrder, pageInt, limitInt)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, blogs)
}

func (bc *blogController) CreateBlog(c *gin.Context) {
	var blog domain.Blog
	err := c.BindJSON(&blog)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	newBlog, err := bc.BlogUsecase.CreateBlog(c, blog)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, newBlog)
}

func (bc *blogController) UpdateBlog(c *gin.Context) {
	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var blog domain.Blog
	err = c.BindJSON(&blog)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	updatedBlog, err := bc.BlogUsecase.UpdateBlog(c, idInt, blog)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, updatedBlog)

}

func (bc *blogController) DeleteBlog(c *gin.Context) {

	id := c.Param("id")

	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	err = bc.BlogUsecase.DeleteBlog(c, idInt)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"message": "Blog deleted successfully"})

}

func (bc *blogController) LikeBlog(c *gin.Context) {
	blogID := c.Param("blog_id")
	authorID := c.Param("author_id")

	blogIDInt, err := strconv.Atoi(blogID)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	authorIDInt, err := strconv.Atoi(authorID)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	likedBlog, err := bc.BlogUsecase.LikeBlog(c, blogIDInt, authorIDInt)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, likedBlog)

}

func (bc *blogController) DislikeBlog(c *gin.Context) {
	blogID := c.Param("blog_id")
	authorID := c.Param("author_id")

	blogIDInt, err := strconv.Atoi(blogID)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	authorIDInt, err := strconv.Atoi(authorID)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	dislikedBlog, err := bc.BlogUsecase.DislikeBlog(c, blogIDInt, authorIDInt)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, dislikedBlog)

}

func (bc *blogController) AddComment(c *gin.Context) {

	blogID := c.Param("blog_id")
	authorID := c.Param("author_id")

	blogIDInt, err := strconv.Atoi(blogID)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	authorIDInt, err := strconv.Atoi(authorID)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	var comment domain.Comment
	err = c.BindJSON(&comment)

	comment.UserID = authorIDInt

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	commentedBlog, err := bc.BlogUsecase.AddComent(c, blogIDInt, authorIDInt, comment.Content)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, commentedBlog)

}

func (bc *blogController) Search(c *gin.Context) {
	author := c.Query("author")
	tags := c.QueryArray("tags")
	title := c.Query("title")

	blogs, err := bc.BlogUsecase.Search(c, author, tags, title)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, blogs)

}

type BlogContent struct {
	Content string `json:"content"`
}

func (bc *blogController) AiRecommendation(c *gin.Context) {
	var content BlogContent

	err := c.ShouldBindJSON(&content)
	// print("content is", content)
	if err != nil {
		print("error is here")
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	recommendation, err := bc.BlogUsecase.AiRecommendation(c, content.Content)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, recommendation)

}
