package controllers

import (
	"blog_g2/domain"
	"blog_g2/infrastructure"

	"log"

	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogController struct {
	Blogusecase    domain.BlogUsecase
	Likeusecase    domain.LikeUsecase
	Commentusecase domain.CommentUsecase
	Dislikeusecase domain.DisLikeUsecase
	Aiservice      domain.AIService
	Medcont        infrastructure.MediaUpload
}

// Blog-controller constructor
func NewBlogController(Blogmgr domain.BlogUsecase, likemgr domain.LikeUsecase, commentmgr domain.CommentUsecase, dislmgr domain.DisLikeUsecase, aiserv domain.AIService, med infrastructure.MediaUpload) *BlogController {
	return &BlogController{
		Blogusecase:    Blogmgr,
		Likeusecase:    likemgr,
		Commentusecase: commentmgr,
		Dislikeusecase: dislmgr,
		Aiservice:      aiserv,
		Medcont:        med,
	}

}

func (controller *BlogController) CreateBlog(c *gin.Context) {
	userid := c.GetString("userid")

	var blog domain.Blog
	if err := c.ShouldBindJSON(&blog); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	if blog.Title == "" || blog.Content == "" || blog.Tags == nil {
		c.JSON(http.StatusBadGateway, gin.H{"message": "Please fill in all fields"})
		return
	}

	blog.UserID, _ = primitive.ObjectIDFromHex(userid)
	blog.Date = time.Now()
	blog.Likes = 0
	blog.DisLikes = 0
	blog.Comments = 0

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

	pages, _ := strconv.Atoi(c.Query("page"))
	sortby := c.DefaultQuery("sortby", "")
	sortdire := c.DefaultQuery("sortdir", "")
	blogs, err := controller.Blogusecase.RetrieveBlog(c, pages, sortby, sortdire)
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
	admin := c.GetBool("isadmin")
	userid := c.GetString("userid")

	if getID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id can not be empty"})
		return
	}

	var blog domain.Blog
	if err := c.BindJSON(&blog); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": "invalid blog request"})
		return
	}

	err := controller.Blogusecase.UpdateBlog(c, blog, getID, admin, userid)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.IndentedJSON(http.StatusAccepted, gin.H{"message": "blog succesfully updated"})
}

func (controller *BlogController) DeleteBlog(c *gin.Context) {
	getID := c.Param("id")
	admin := c.GetBool("isadmin")
	userid := c.GetString("userid")

	if getID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "id can not be empty"})
		return
	}

	err := controller.Blogusecase.DeleteBlog(c, getID, admin, userid)
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
