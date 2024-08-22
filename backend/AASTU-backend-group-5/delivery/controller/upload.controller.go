package controller

import (
	"net/http"

	"github.com/RealEskalate/blogpost/usecase"
	"github.com/gin-gonic/gin"
)

type UploadController struct{
	uploadUC usecase.UploadProfileUsecase
}

func NewUploadController(upload_uc usecase.UploadProfileUsecase) *UploadController {
	return &UploadController{
		uploadUC: upload_uc,
	}
}

func (ctrl *UploadController) UplaodImg() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Param("id")
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

		err = ctrl.uploadUC.UploadPicture("assets/uploads/"+file.Filename , id)
		if err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError , gin.H{"error" : err.Error()})
			return
		}
		
		ctx.IndentedJSON(http.StatusOK , gin.H{"message" : "success"})
	}
}
