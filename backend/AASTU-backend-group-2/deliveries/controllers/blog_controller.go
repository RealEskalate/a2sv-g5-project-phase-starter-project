package controllers

import (
	"blog_g2/domain"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	Blogusecase domain.BlogUsecase
}

// Blog-controller constructor
func NewBlogController(Blogmgr domain.BlogUsecase) *BlogController {
	return &BlogController{
		Blogusecase: Blogmgr,
	}

}

func (controller *BlogController) CreateBlog(c *gin.Context) {
	// role, exists := c.Get("role")
	// if !exists || (role != "user" && role != "admin") {
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
	blog.Date = time.Now()
	err := controller.Blogusecase.CreateBlog(c, blog)
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

}
func (controller *BlogController) DeleteBlog(c *gin.Context) {

}
func (controller *BlogController) SearchBlog(c *gin.Context) {

}
func (controller *BlogController) FilterBlog(c *gin.Context) {

}
