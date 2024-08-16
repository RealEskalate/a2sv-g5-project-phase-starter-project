package routers

import (
	infrastructure "blogs/Infrastructure"
	repositories "blogs/Repositories"
	"blogs/mongo"

	"github.com/gin-gonic/gin"
)

func NewSignupRoute(config *infrastructure.Config, DB mongo.Database, SignupRoute *gin.RouterGroup) {
	repo := repositories.NewSignupRepository(DB , )
	
	SignupRoute.POST("/auth/create")
	SignupRoute.GET("/auth/google")
	

	// blogRouter.GET("/get")
	// blogRouter.GET("/get/:id")
	// blogRouter.PUT("/update/:id")
	// blogRouter.DELETE("/delete/:id")
	// blogRouter.POST("/comment/:id")

}
