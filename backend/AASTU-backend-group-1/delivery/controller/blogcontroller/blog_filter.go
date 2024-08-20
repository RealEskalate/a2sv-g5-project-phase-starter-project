package blogcontroller

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// BlogController ...


func (b *BlogController) FilterBlog(ctx *gin.Context) {
	var filter struct {
		Tags     []string `json:"tags"`
		DateFrom string   `json:"dateFrom"`
		DateTo   string   `json:"dateTo"`
	}

	if err := ctx.ShouldBindJSON(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dateFrom, err := time.Parse(time.RFC3339, filter.DateFrom)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid dateFrom")
		return
	}

	dateTo, err := time.Parse(time.RFC3339, filter.DateTo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, "invalid dateTo")
		return
	}

	blogs, err := b.BlogUsecase.FilterBlog(filter.Tags, dateFrom, dateTo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, "internal server error")
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}