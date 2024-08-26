package blogcontroller

import (
	"blogs/config"
	"blogs/domain"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func (b *BlogController) FilterBlog(ctx *gin.Context) {
	var filter struct {
		Tags     []string `json:"tags"`
		DateFrom string   `json:"date_from"`
		DateTo   string   `json:"date_to"`
	}

	err := ctx.ShouldBindJSON(&filter)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, domain.APIResponse{
			Status:  http.StatusBadRequest,
			Message: "Invalid Request",
			Error:   err.Error(),
		})
		return
	}

	var dateFrom, dateTo time.Time
	if filter.DateFrom == "" {
		dateFrom = time.Time{}
	} else {
		var err error
		dateFrom, err = time.Parse(time.RFC3339, filter.DateFrom)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, domain.APIResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid Request",
				Error:   "invalid date_from",
			})
			return
		}
	}

	if filter.DateTo == "" {
		dateTo = time.Now()
	} else {
		var err error
		dateTo, err = time.Parse(time.RFC3339, filter.DateTo)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, domain.APIResponse{
				Status:  http.StatusBadRequest,
				Message: "Invalid Request",
				Error:   "invalid date_to",
			})
			return
		}
	}

	if len(filter.Tags) == 0 {
		filter.Tags = []string{}
	}

	blogs, err := b.BlogUsecase.FilterBlog(filter.Tags, dateFrom, dateTo)
	if err != nil {
		code := config.GetStatusCode(err)

		ctx.JSON(code, domain.APIResponse{
			Status:  code,
			Message: "Cannot filter blogs",
			Error:   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, domain.APIResponse{
		Status: http.StatusOK,
		Count:  len(blogs),
		Data:   blogs,
	})
}
