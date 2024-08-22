package Infrastructure

import (
	"AAiT-backend-group-8/Domain"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const dateLayout = "2006-01-02" // example layout, adjust as needed

func GetSearchParams(ctx *gin.Context) *Domain.SearchCriteria {
	Title := ctx.Query("title")
	Author := ctx.Query("author")
	Tags := ctx.QueryArray("tags")
	StartDateStr := ctx.Query("startDate")
	EndDateStr := ctx.Query("endDate")
	MinViewsStr := ctx.Query("minViews")
	SortBy := ctx.Query("sortBy")
	PageStr := ctx.Query("page")
	PageSizeStr := ctx.Query("pageSize")

	var err error
	var Page int
	var PageSize int
	var StartDate time.Time
	var EndDate time.Time
	var MinViews int

	// Parse page number with default value
	Page, err = strconv.Atoi(PageStr)
	if err != nil || Page <= 0 {
		Page = 1
	}

	// Parse page size with default value
	PageSize, err = strconv.Atoi(PageSizeStr)
	if err != nil || PageSize <= 0 {
		PageSize = 10
	}

	// Parse min views with default value
	MinViews, err = strconv.Atoi(MinViewsStr)
	if err != nil {
		MinViews = 0
	}

	// Parse start date with default value
	StartDate, err = time.Parse(dateLayout, StartDateStr)
	if err != nil {
		StartDate = time.Time{} // zero time
	}

	// Parse end date with default value
	EndDate, err = time.Parse(dateLayout, EndDateStr)
	if err != nil {
		EndDate = time.Now() // current time
	}

	return &Domain.SearchCriteria{
		Title:     Title,
		Author:    Author,
		Tags:      Tags,
		StartDate: StartDate,
		EndDate:   EndDate,
		MinViews:  MinViews,
		SortBy:    SortBy,
		Page:      Page,
		PageSize:  PageSize,
	}
}
