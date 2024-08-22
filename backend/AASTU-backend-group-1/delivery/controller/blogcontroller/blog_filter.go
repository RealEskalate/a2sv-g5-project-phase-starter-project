package blogcontroller

import (
	"blogs/config"
	"log"
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

	if err := ctx.ShouldBindJSON(&filter); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dateFrom, err := time.Parse(time.RFC3339, filter.DateFrom)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid date_from"})
		return
	}

	dateTo, err := time.Parse(time.RFC3339, filter.DateTo)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid date_to"})
		return
	}

	blogs, err := b.BlogUsecase.FilterBlog(filter.Tags, dateFrom, dateTo)
	if err != nil {
		code := config.GetStatusCode(err)

		if code == http.StatusInternalServerError {
			log.Println(err)
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		ctx.JSON(code, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, blogs)
}
