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

func NewBlogController(bu domain.BlogUseCase, redisClient redis.Client) *BlogController {
	return &BlogController{
		BlogUsecase: bu,
		RedisClient: redisClient,
	}
}

func (bc *BlogController) CreateBlog(cxt *gin.Context) {
	var blog domain.Blog
	err := cxt.ShouldBindJSON(&blog)
	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
		return
	}
	createdBlog, err := bc.BlogUsecase.CreateBlog(cxt, &blog)
	// println(createdBlog, err.Error())

	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return
	}

	cxt.JSON(http.StatusCreated, gin.H{"message": "Blog Created Successfully", "Blog": createdBlog})

}

func (bc *BlogController) GetBlog(c *gin.Context) {
	// we need to use redis
	id := c.Param("id")

	// check if the blog exists in redis
	cachedBlog, err := bc.RedisClient.Get(c, id)
	if err == rd.Nil {

		blog, err := bc.BlogUsecase.GetBlog(c, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
		// marshal the blog
		blogM, err := json.Marshal(blog)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
		// cache the blog
		err = bc.RedisClient.Set(c, id, blogM, 0)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
			return
		}
		c.JSON(http.StatusOK, blog)
		return

	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		return

	} else {
		c.JSON(http.StatusOK, cachedBlog)
		return
	}

}

func (bc *BlogController) GetBlogs(c *gin.Context) {
	var pagination domain.Pagination
	err := c.BindQuery(&pagination)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
		return
	}

	cacheKey := fmt.Sprintf("blogs:%d:%d", pagination.Page, pagination.PageSize)

	cachedBlogs, err := bc.RedisClient.Get(c, cacheKey)
	if err == rd.Nil {
		blogs, err := bc.BlogUsecase.GetBlogs(c, &pagination)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		}
		// cache the blog for 5 minutes
		err = bc.RedisClient.Set(c, cacheKey, blogs, 300)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
		}
		c.JSON(http.StatusOK, blogs)

	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	} else {
		c.JSON(http.StatusOK, cachedBlogs)
	}

}

// update blog should be partially updated
func (bc *BlogController) UpdateBlog(cxt *gin.Context) {
	id := cxt.Param("id")
	var blog domain.Blog
	err := cxt.BindJSON(&blog)
	if err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"message": "Invalid Request"})
	}
	err = bc.BlogUsecase.UpdateBlog(cxt, &blog, id)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	cxt.JSON(http.StatusOK, gin.H{"message": "Blog Updated Successfully"})
}

func (bc *BlogController) DeleteBlog(cxt *gin.Context) {
	id := cxt.Param("id")
	err := bc.BlogUsecase.DeleteBlog(cxt, id)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	cxt.JSON(http.StatusOK, gin.H{"message": "Blog Deleted Successfully"})
}

func (bc *BlogController) LikeBlog(cxt *gin.Context) {
	blogID := cxt.Param("blog_id")
	userID := cxt.Param("user_id")
	err := bc.BlogUsecase.LikeBlog(cxt, blogID, userID)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server Error"})
	}
	cxt.JSON(http.StatusOK, gin.H{"message": "Blog Liked Successfully"})
}

func (ctrl *BlogController) UnlikeBlog(cxt *gin.Context) {
	blogID := cxt.Param("blogID")
	userID := cxt.Param("userID")

	err := ctrl.BlogUsecase.UnlikeBlog(cxt, blogID, userID)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"message": "Blog unliked successfully"})
}

// CommentBlog handles HTTP requests to add a comment to a blog post
func (ctrl *BlogController) CommentBlog(cxt *gin.Context) {
	blogID := cxt.Param("blogID")

	var comment domain.Comment
	if err := cxt.BindJSON(&comment); err != nil {
		cxt.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err := ctrl.BlogUsecase.CommentBlog(cxt, blogID, &comment)
	if err != nil {
		cxt.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	cxt.JSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
}
