package blog_controller

import (
	"/domain"
	"net/http"
	"your_project/services"

	"github.com/gin-gonic/gin"
)

func (bc *BlogController) FileUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		formFile, _, err := c.Request.FormFile("file")
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				domain.MediaDto{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Select a file to upload"},
				})
			return
		}

		uploadUrl, err := services.NewMediaService().FileUpload(domain.File{File: formFile})
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				domain.MediaDto{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Error uploading file"},
				})
			return
		}

		c.JSON(
			http.StatusOK,
			domain.MediaDto{
				StatusCode: http.StatusOK,
				Message:    "success",
				Data:       map[string]interface{}{"data": uploadUrl},
			})
	}
}

func (bc *BlogController) RemoteUpload() gin.HandlerFunc {
	return func(c *gin.Context) {
		var url domain.Url

		// Validate the request body
		if err := c.BindJSON(&url); err != nil {
			c.JSON(
				http.StatusBadRequest,
				domain.MediaDto{
					StatusCode: http.StatusBadRequest,
					Message:    "error",
					Data:       map[string]interface{}{"data": err.Error()},
				})
			return
		}

		uploadUrl, err := services.NewMediaService().RemoteUpload(url)
		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				domain.MediaDto{
					StatusCode: http.StatusInternalServerError,
					Message:    "error",
					Data:       map[string]interface{}{"data": "Error uploading file"},
				})
			return
		}

		c.JSON(
			http.StatusOK,
			domain.MediaDto{
				StatusCode: http.StatusOK,
				Message:    "success",
				Data:       map[string]interface{}{"data": uploadUrl},
			})
	}
}
