package routers

import (
	controllers "blogs/Delivery/controllers"
	infrastructure "blogs/Infrastructure"
	repositories "blogs/Repositories"
	usecases "blogs/Usecases"
	"blogs/mongo"
	"time"

	"github.com/gin-gonic/gin"
)

func NewSignupRoute(config *infrastructure.Config, DB mongo.Database, SignupRoute *gin.RouterGroup) {

	repo := repositories.NewSignupRepository(DB, config.UserCollection)
	passwordService := infrastructure.NewPasswordService()

	usecase := usecases.NewSignupUseCase(repo, time.Duration(config.ContextTimeout)*time.Second , passwordService)
	signup := controllers.SignupController{
		SignupUsecase: usecase,
	}

	SignupRoute.POST("/signup", signup.Signup)
	SignupRoute.POST("/signup/verify", signup.VerifyOTP)
	SignupRoute.POST("/reset", signup.ForgotPassword)
	// SignupRoute.POST("/resendotp", signup.ForgotPassword)
	
}
