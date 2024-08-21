package controller

import (
	"net/http"
	"strconv"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/gin-gonic/gin"
)

type PopularityController struct {
	popularityUsecase domain.BlogPopularityUsecase
}

func NewPopularityController(popularityUsecase domain.BlogPopularityUsecase) *PopularityController {
	return &PopularityController{popularityUsecase: popularityUsecase}
}

func (bc *PopularityController) GetPopularBlogs(c *gin.Context) {
	sortBy := c.Query("sort_by")
	sortOrderStr := c.Query("sort_order")

	sortOrder, err := strconv.Atoi(sortOrderStr)
	if err != nil {
		sortOrder = -1
	}

	blogs, err := bc.popularityUsecase.GetSortedPopularBlogs(sortBy, sortOrder)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, blogs)
}
