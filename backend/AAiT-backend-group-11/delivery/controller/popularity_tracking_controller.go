package controller

import (
	"backend-starter-project/domain/dto"
	"backend-starter-project/domain/interfaces"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PopularityTrackingController struct {
	popularityTrackingService interfaces.PopularityTrackingService
}

func NewPopularityTrackingController(popularityTrackingService interfaces.PopularityTrackingService) *PopularityTrackingController {
	return &PopularityTrackingController{
		popularityTrackingService: popularityTrackingService,
	}
}


func (ptc *PopularityTrackingController) LikeBlogPost(c *gin.Context) {
	id := c.Param("id")

	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Error: "User not found",
		})
		return
	}

	err := ptc.popularityTrackingService.LikeBlogPost(id, userId.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
	})
}

func (ptc *PopularityTrackingController) DislikeBlogPost(c *gin.Context) {
	id := c.Param("id")

	userId, ok := c.Get("userId")
	if !ok {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Error: "User not found",
		})
		return
	}

	err := ptc.popularityTrackingService.DislikeBlogPost(id, userId.(string))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Error: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
	})
}