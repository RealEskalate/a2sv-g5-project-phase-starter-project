package routers

import (
	infrastructure "blogs/Infrastructure"
	repositories "blogs/Repositories"
	usecases "blogs/Usecases"
	controllers "blogs/Delivery/controllers"
	"blogs/mongo"
	"time"

	"github.com/gin-gonic/gin"
)

func NewSignupRoute(config *infrastructure.Config, DB mongo.Database, SignupRoute *gin.RouterGroup) {
	repo := repositories.NewSignupRepository(DB , config.UserCollection)
	
	usecase := usecases.NewSignupUseCase(repo , time.Duration(config.ContextTimeout) * time.Second)
	signup := controllers.SignupController{
		SignupUsecase: usecase,
	}

	SignupRoute.POST("/signup", signup.Signup)
	// otp verifyer route
	SignupRoute.POST("/signup/verify" , signup.VerifyOTP)
	// Google Auth	
	SignupRoute.GET("/auth/google" , signup.GoogleAuth)
	
	
	// SignupRoute.POST("/auth/signup" )
	// SignupRoute.GET("/auth/google")
	

	// blogRouter.GET("/get")
	// blogRouter.GET("/get/:id")
	// blogRouter.PUT("/update/:id")
	// blogRouter.DELETE("/delete/:id")
	// blogRouter.POST("/comment/:id")

}
