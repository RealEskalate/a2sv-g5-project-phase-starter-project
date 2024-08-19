package Helper

import (
	"AAiT-backend-group-8/Domain"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

func GetSearchParams(ctx *gin.Context) *Domain.SearchCriteria {
	Title := ctx.Query("title")
	Author := ctx.Query("author")
	Tags := ctx.QueryArray("tags")
	StartDateStr := ctx.Query("startDate")
	EndDateStr := ctx.Query("endDate")
	MinViewsStr := ctx.Query("minViews")
	SortBy := ctx.Query("sortBy")
	Order := ctx.Query("order")
	PageStr := ctx.Query("page")
	PageSizeStr := ctx.Query("pageSize")

	var err error
	var Page int
	var PageSize int
	var StartDate time.Time
	var EndDate time.Time
	var MinViews int

	Page, err = strconv.Atoi(PageStr)
	if err != nil {
		Page = 1
	}

	MinViews, err = strconv.Atoi(MinViewsStr)

	if err != nil {
		MinViews = 0
	}

	PageSize, err = strconv.Atoi(PageSizeStr)
	if err != nil {
		PageSize = 10
	}

	StartDate, err = time.Parse(time.Layout, StartDateStr)
	if err != nil {
		StartDate = time.Time{}
	}

	EndDate, err = time.Parse(time.Layout, EndDateStr)
	if err != nil {
		EndDate = time.Now()
	}

	return &Domain.SearchCriteria{
		Title:     Title,
		Author:    Author,
		Tags:      Tags,
		StartDate: StartDate,
		EndDate:   EndDate,
		MinViews:  MinViews,
		SortBy:    SortBy,
		Order:     Order,
		Page:      Page,
		PageSize:  PageSize,
	}
}
