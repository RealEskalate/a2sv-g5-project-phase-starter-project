package controllers

import (
	"blog_g2/domain"
	"net/http"

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
	getID := c.Param("id")
	if getID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id can not be empty"})
	}

	var blog domain.Blog
	if err := c.BindJSON(&blog); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid blog request"})
	}

	err := controller.Blogusecase.UpdateBlog(c, blog, getID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "blog succesfully updated"})
}
func (controller *BlogController) DeleteBlog(c *gin.Context) {
	getID := c.Param("id")
	if getID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id can not be empty"})
	}

	err := controller.Blogusecase.DeleteBlog(c, getID)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "blog succesfully deleted"})
}
func (controller *BlogController) SearchBlog(c *gin.Context) {

}
func (controller *BlogController) FilterBlog(c *gin.Context) {

}
