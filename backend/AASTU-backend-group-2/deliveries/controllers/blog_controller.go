package controllers

import (
	"blog_g2/domain"

	"log"

	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	Blogusecase    domain.BlogUsecase
	Likeusecase    domain.LikeUsecase
	Commentusecase domain.CommentUsecase
	Dislikeusecase domain.DisLikeUsecase
	Aiservice      domain.AIService
}

// Blog-controller constructor
func NewBlogController(Blogmgr domain.BlogUsecase, likemgr domain.LikeUsecase, commentmgr domain.CommentUsecase, dislmgr domain.DisLikeUsecase, aiserv domain.AIService) *BlogController {
	return &BlogController{
		Blogusecase:    Blogmgr,
		Likeusecase:    likemgr,
		Commentusecase: commentmgr,
		Dislikeusecase: dislmgr,
		Aiservice:      aiserv,
	}

}

func (controller *BlogController) CreateBlog(c *gin.Context) {
	// role, exists := c.Get("role")
	// if !exists || (role != "user" && role != "admin") {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	// 	return
	// }

	// userid, exists := c.Get("userid")
	// if !exists {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	// 	return
	// }

	var blog domain.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if blog.Title == "" || blog.Content == "" || blog.Tags == nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Please fill in all fields"})
		return
	}
	// blog.UserID = userid

	blog.Date = time.Now()
	err := controller.Blogusecase.CreateBlog(c, &blog)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, blog)

}
func (controller *BlogController) RetrieveBlog(c *gin.Context) {
	// role, exists := c.Get("role")
	// if !exists || (role != "user" && role != "admin") {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"message": "Unauthorized"})
	// 	return
	// }
	pages, _ := strconv.Atoi(c.Query("page"))
	blogs, err := controller.Blogusecase.RetrieveBlog(c, pages)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, blogs)
}
func (controller *BlogController) UpdateBlog(c *gin.Context) {
	getID := c.Param("id")
	if getID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id can not be empty"})
		return
	}

	var blog domain.Blog
	if err := c.BindJSON(&blog); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid blog request"})
		return
	}

	err := controller.Blogusecase.UpdateBlog(c, blog, getID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "blog succesfully updated"})
}
func (controller *BlogController) DeleteBlog(c *gin.Context) {
	getID := c.Param("id")
	if getID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id can not be empty"})
		return
	}

	err := controller.Blogusecase.DeleteBlog(c, getID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "blog succesfully deleted"})
}
func (controller *BlogController) SearchBlog(c *gin.Context) {
	name := c.DefaultQuery("name", "")
	author := c.DefaultQuery("user", "")

	blogs, err := controller.Blogusecase.SearchBlog(c, name, author)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"blogs": blogs})
}
func (controller *BlogController) FilterBlog(c *gin.Context) {
	tags := c.QueryArray("tags[]")
	date := c.DefaultQuery("date", "")

	log.Println(date)

	log.Println(tags)

	convDate, _ := time.Parse("2006-01-02", date)

	log.Println(convDate)

	blogs, err := controller.Blogusecase.FilterBlog(c, tags, convDate)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"blogs": blogs})
}

func (h *BlogController) GeneratePost(c *gin.Context) {
	var req domain.PostRequest

	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	post, err := h.Aiservice.GeneratePost(req.Title, req.Content)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"title":   post.Title,
		"content": post.Content,
	})
}
