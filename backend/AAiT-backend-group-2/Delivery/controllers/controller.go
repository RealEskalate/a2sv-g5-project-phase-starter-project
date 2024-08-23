package controllers

import (
	domain "AAiT-backend-group-2/Domain"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BlogController struct {
	blogUseCase    domain.BlogUseCase
	commentUseCase domain.CommentUsecase
}

func NewBlogController(bu domain.BlogUseCase, cu domain.CommentUsecase) *BlogController {
	return &BlogController{
		blogUseCase:    bu,
		commentUseCase: cu,
	}
}

func (bc *BlogController) CreateBlog(c *gin.Context) {
	var req domain.RequestBlog
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	author := c.GetString("userID")
	if author == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	err := bc.blogUseCase.CreateBlog(c.Request.Context(), &req, author)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Blog created successfully"})
}

func (bc *BlogController) GetAllBlogs(c *gin.Context) {

	pageStr := c.DefaultQuery("page", "1")
	pageSizeStr := c.DefaultQuery("pageSize", "10")
	sortBy := c.DefaultQuery("sortBy", "created_at")
	sortOrder := c.DefaultQuery("sortOrder", "desc")

	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	pageSize, err := strconv.Atoi(pageSizeStr)
	if err != nil || pageSize < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page size"})
		return
	}

	blogs, total, err := bc.blogUseCase.GetAllBlogs(c.Request.Context(), page, pageSize, sortBy, sortOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":      blogs,
		"total":     total,
		"page":      page,
		"pageSize":  pageSize,
		"sortBy":    sortBy,
		"sortOrder": sortOrder,
	})
}

func (bc *BlogController) GetBlogByID(c *gin.Context) {
	id := c.Param("id")
	blog, err := bc.blogUseCase.GetBlogByID(c.Request.Context(), id)
	if err != nil {
		if err.Error() == "blog not found" {
			c.JSON(http.StatusNotFound, gin.H{"error": "Blog not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, blog)
}

func (bc *BlogController) UpdateBlog(c *gin.Context) {
	id := c.Param("id")
	author := c.GetString("userID")

	if author == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
	}

	var req domain.RequestBlog
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	err := bc.blogUseCase.UpdateBlog(c.Request.Context(), &req, author, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blog updated successfully"})
}

func (bc *BlogController) DeleteBlog(c *gin.Context) {
	id := c.Param("id")
	author := c.GetString("userID")
	role := c.GetString("role")

	if author == "" && (role == "" || role != "admin") {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
	}
	err := bc.blogUseCase.DeleteBlog(c.Request.Context(), author, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

func (bc *BlogController) FilterBlogs(c *gin.Context) {
	tags := c.QueryArray("tags")
	startDate := c.Query("startDate")
	endDate := c.Query("endDate")
	sortBy := c.DefaultQuery("sortBy", "created_at")

	blogs, err := bc.blogUseCase.FilterBlogs(c.Request.Context(), tags, startDate, endDate, sortBy)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": blogs})
}

func (bc *BlogController) CreateComment(c *gin.Context) {
	var req domain.Comment
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	author := c.GetString("userID")
	if author == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	req.Author = author
	req.BlogID = c.Param("id")

	req.ID = uuid.New().String()

	err := bc.commentUseCase.CreateComment(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "Comment added successfully"})
}

func (bc *BlogController) GetCommentsByBlogID(c *gin.Context) {
	blogID := c.Param("id")
	comments, err := bc.commentUseCase.GetCommentsByBlogID(c.Request.Context(), blogID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"comments": comments})
}

func (bc *BlogController) UpdateComment(c *gin.Context) {
	id := c.Param("comment_id")
	author := c.GetString("userID")
	if author == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	var req domain.Comment
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	req.ID = id
	req.Author = author

	err := bc.commentUseCase.UpdateComment(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment updated successfully"})
}

func (bc *BlogController) DeleteComment(c *gin.Context) {
	id := c.Param("comment_id")
	author := c.GetString("userID")
	if author == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	err := bc.commentUseCase.DeleteComment(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Comment deleted successfully"})
}
