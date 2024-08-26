package controllers

import (
	"blog_g2/infrastructure"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/labstack/echo"
)

func (h *BlogController) FileUpload(c *gin.Context) {
	//upload
	formHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			infrastructure.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": "Select a file to upload"},
			})
		return
	}

	//get file from header
	formFile, err := formHeader.Open()
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			infrastructure.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": err.Error()},
			})
		return
	}

	uploadUrl, err := h.Medcont.FileUpload(infrastructure.File{File: formFile})
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			infrastructure.MediaDto{
				StatusCode: http.StatusInternalServerError,
				Message:    "error",
				Data:       &echo.Map{"data": err.Error()},
			})
		return
	}

	c.JSON(
		http.StatusOK,
		infrastructure.MediaDto{
			StatusCode: http.StatusOK,
			Message:    "success",
			Data:       &echo.Map{"data": uploadUrl},
		})
}
