package controller

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/RealEskalate/blogpost/domain"
	"github.com/gin-gonic/gin"
)

type PopularityController struct {
	popularityUsecase domain.BlogPopularityUsecase
}

func NewPopularityController(popularityUsecase domain.BlogPopularityUsecase) *PopularityController {
	return &PopularityController{popularityUsecase: popularityUsecase}
}

func (bc *PopularityController) GetPopularBlogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		sortByStr := c.Query("sort_by")
		sortOrderStr := c.Query("sort_order")

		sortByValues := strings.Split(sortByStr, ",")
		sortOrderValues := strings.Split(sortOrderStr, ",")

		var sortBys []domain.SortBy
		for _, sb := range sortByValues {
			sortBy := domain.SortBy(sb)
			if isValidSortBy(sortBy) {
				sortBys = append(sortBys, sortBy)
			} else {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sort_by parameter"})
				return
			}
		}

		var sortOrders []domain.SortOrder
		for _, so := range sortOrderValues {
			sortOrder, err := strconv.Atoi(so)
			if err != nil {
				sortOrder = int(domain.SortOrderDescending)
			}
			if sortOrder != int(domain.SortOrderAscending) && sortOrder != int(domain.SortOrderDescending) {
				c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid sort_order parameter"})
				return
			}
			sortOrders = append(sortOrders, domain.SortOrder(sortOrder))
		}

		blogs, err := bc.popularityUsecase.GetSortedPopularBlogs(sortBys, sortOrders)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, blogs)
	}
}

func isValidSortBy(sortBy domain.SortBy) bool {
	switch sortBy {
	case domain.SortByLikeCount, domain.SortByCommentCount, domain.SortByPublishDate, domain.SortByEngagement, domain.SortByDislikeCount:
		return true
	}
	return false
}
