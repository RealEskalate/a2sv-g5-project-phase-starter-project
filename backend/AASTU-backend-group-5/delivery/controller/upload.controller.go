package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UploadController struct{}

func (UploadController) UplaodImg() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		file,err := ctx.FormFile("image")
		if err != nil {
			ctx.IndentedJSON(http.StatusBadRequest , gin.H{"error" : err.Error()})
			return
		}
		err = ctx.SaveUploadedFile(file , "assets/uploads/"+file.Filename)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError , gin.H{"error" : err.Error()})
			return
		}

		ctx.IndentedJSON(http.StatusOK , gin.H{"message" : "success"})
	}
}
