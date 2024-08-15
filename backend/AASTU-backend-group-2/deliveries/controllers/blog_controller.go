package controllers

import (
	"blog_g2/domain"
	"log"
	"net/http"
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

}
func (controller *BlogController) RetrieveBlog(c *gin.Context) {

}
func (controller *BlogController) UpdateBlog(c *gin.Context) {

}
func (controller *BlogController) DeleteBlog(c *gin.Context) {

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
