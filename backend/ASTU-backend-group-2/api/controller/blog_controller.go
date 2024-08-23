package controller

import (
	"context"
	"strconv"
	"time"

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
	// GetComments() gin.HandlerFunc
	// CreateComment() gin.HandlerFunc
	// GetComment() gin.HandlerFunc
	// UpdateComment() gin.HandlerFunc
	// DeleteComment() gin.HandlerFunc
	// CreateLike() gin.HandlerFunc
}

type BlogController struct {
	BlogUsecase domain.BlogUsecase
	Env         *bootstrap.Env
}

func (bc *BlogController) GetBlogs() gin.HandlerFunc {
	return func(c *gin.Context) {

		page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
		dateFrom, _ := time.Parse(time.RFC3339, c.Query("date_from"))
		dateTo, _ := time.Parse(time.RFC3339, c.Query("date_to"))
		tags, _ := c.GetQueryArray("tags")
		popularityFrom, _ := strconv.Atoi(c.Query("popularity_from"))
		popularityTo, _ := strconv.Atoi(c.Query("popularity_to"))

		var blogFilter domain.BlogFilter

		blogFilter = domain.BlogFilter{
			Title:          c.Query("title"),
			Tags:           tags,
			DateFrom:       dateFrom,
			DateTo:         dateTo,
			Limit:          10, // 10 pages perfilter
			Pages:          page,
			PopularityFrom: popularityFrom,
			PopularityTo:   popularityTo,
		}

		blogs, pagination, err := bc.BlogUsecase.GetAllBlogs(context.Background(), blogFilter)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		res := domain.PaginatedResponse{
			Data:     blogs,
			MetaData: pagination,
		}

		c.JSON(200, res)
	}
}

func (bc *BlogController) GetBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		blog, err := bc.BlogUsecase.GetBlogByID(context.Background(), blogID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, blog)
	}
}

func (bc *BlogController) CreateBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		var newBlog domain.BlogIn
		if err := c.ShouldBindJSON(&newBlog); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		blog, err := bc.BlogUsecase.CreateBlog(context.Background(), &newBlog)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(201, blog)
	}
}

func (bc *BlogController) BatchCreateBlog() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}

func (bc *BlogController) UpdateBlog() gin.HandlerFunc {
	return func(c *gin.Context) {
		blogID := c.Param("id")
		var updatedBlog domain.BlogIn
		if err := c.ShouldBindJSON(&updatedBlog); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		blog, err := bc.BlogUsecase.UpdateBlog(context.Background(), blogID, &updatedBlog)
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
		err := bc.BlogUsecase.DeleteBlog(context.TODO(), blogID)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(204, nil)
	}
}
func (bc *BlogController) GetByTags() gin.HandlerFunc {
	return func(c *gin.Context) {
		tags := c.Query("tags")
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.Query("page"), 10, 64)

		blogs, pagination, err := bc.BlogUsecase.GetByTags(context.TODO(), []string{tags}, limit, page)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"blogs": blogs, "pageination": pagination})
	}
}

func (bc *BlogController) GetbyPopularity() gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
		blogs, pagination, err := bc.BlogUsecase.GetByPopularity(context.Background(), limit, page)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"blogs": blogs, "pageination": pagination})
	}
}
func (bc *BlogController) Search() gin.HandlerFunc {

	return func(c *gin.Context) {
		searchTerm := c.Query("search")
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
		blogs, pagination, err := bc.BlogUsecase.Search(context.Background(), searchTerm, limit, page)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"blogs": blogs, "pageination": pagination})
	}
}
func (bc *BlogController) SortByDate() gin.HandlerFunc {
	return func(c *gin.Context) {
		limit, _ := strconv.ParseInt(c.Query("limit"), 10, 64)
		page, _ := strconv.ParseInt(c.Query("page"), 10, 64)
		blogs, pagination, err := bc.BlogUsecase.SortByDate(context.Background(), limit, page)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"blogs": blogs, "pageination": pagination})
	}
}
