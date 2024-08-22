package controller

import (
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/redis"
	"encoding/json"
	"fmt"
	"net/http"

	rd "github.com/go-redis/redis/v8"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	BlogUsecase domain.BlogUseCase
	RedisClient redis.Client
}

func NewBlogController(bu domain.BlogUseCase,redisClient redis.Client) *BlogController {
	return &BlogController{
		BlogUsecase: bu,
		RedisClient: redisClient,
	}
}

func (bc *BlogController) CreateBlog(c *gin.Context) {
	var blog domain.Blog
	err := c.ShouldBindJSON(&blog)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"message":"Invalid Request"})
		return
	}
	err = bc.BlogUsecase.CreateBlog(&blog)

	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message":"Internal Server Error"})
		return
	}

	c.JSON(http.StatusCreated,gin.H{"message":"Blog Created Successfully"})
}

func (bc *BlogController) GetBlog(c *gin.Context) {
	// we need to use redis
	id := c.Param("id")

	// check if the blog exists in redis
	cachedBlog, err := bc.RedisClient.Get(c,id)
	if err == rd.Nil{

		blog, err := bc.BlogUsecase.GetBlog(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"message":"Internal Server Error"})
			return
		}
		// marshal the blog
		blogM, err := json.Marshal(blog)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"message":"Internal Server Error"})
			return
		}
		// cache the blog
		err = bc.RedisClient.Set(c,id,blogM,0)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"message":"Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK,blog)
		return

	}else if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message":"Internal Server Error"})
		return

	}else{
		c.JSON(http.StatusOK,cachedBlog)
		return
	}

}



func (bc *BlogController) GetBlogs(c *gin.Context) {
	var pagination domain.Pagination
	err := c.BindQuery(&pagination)

	if err != nil{
		c.JSON(http.StatusBadRequest,gin.H{"message":"Invalid Request"})
		return
	}

	cacheKey := fmt.Sprintf("blogs:%d:%d",pagination.Page,pagination.PageSize)

	cachedBlogs, err := bc.RedisClient.Get(c,cacheKey)
	if err == rd.Nil{
		blogs, err := bc.BlogUsecase.GetBlogs(&pagination)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"message":"Internal Server Error"})
		}
		// cache the blog for 5 minutes
		err = bc.RedisClient.Set(c,cacheKey,blogs,300)
		if err != nil {
			c.JSON(http.StatusInternalServerError,gin.H{"message":"Internal Server Error"})
		}
		c.JSON(http.StatusOK,blogs)

	}else if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message":"Internal Server Error"})
	}else{
		c.JSON(http.StatusOK,cachedBlogs)
	}

}

// update blog should be partially updated
func (bc *BlogController) UpdateBlog(c *gin.Context) {
	id := c.Param("id")
	var blog domain.Blog
	err := c.BindJSON(&blog)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
		return
	}

	// Update the blog in the database
	err = bc.BlogUsecase.UpdateBlog(&blog, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	// Invalidate the cache for this blog
	cacheKey := id
	bc.RedisClient.Del(c, cacheKey)

	c.JSON(http.StatusOK, gin.H{"message": "Blog Updated Successfully"})
}


func (bc *BlogController) DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	err := bc.BlogUsecase.DeleteBlog(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message":"Internal Server Error"})
	}

	// Invalidate the cache for this blog
	cacheKey := id
	bc.RedisClient.Del(c, cacheKey)
	c.JSON(http.StatusOK,gin.H{"message":"Blog Deleted Successfully"})
}

func (bc *BlogController) LikeBlog(c *gin.Context) {
	blogID := c.Param("blog_id")
	userID := c.Param("user_id")
	err := bc.BlogUsecase.LikeBlog(blogID,userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError,gin.H{"message":"Internal Server Error"})
	}
	c.JSON(http.StatusOK,gin.H{"message":"Blog Liked Successfully"})
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
