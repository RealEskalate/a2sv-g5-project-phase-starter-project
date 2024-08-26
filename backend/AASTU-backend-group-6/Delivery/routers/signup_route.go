package routers

import (
	controllers "blogs/Delivery/controllers"
	infrastructure "blogs/Infrastructure"
	repositories "blogs/Repositories"
	usecases "blogs/Usecases"
	"blogs/mongo"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func NewSignupRoute(config *infrastructure.Config, DB mongo.Database, SignupRoute *gin.RouterGroup) {

	redisClient := redis.NewClient(&redis.Options{
		Addr: "localhost:6379", 
		Password: "",            
		DB: 0,                   
	})
	
	ratelimiter := infrastructure.NewRateLimiter(redisClient)
	repo := repositories.NewSignupRepository(DB, config.UserCollection)
	passwordService := infrastructure.NewPasswordService()
	unverRepo := repositories.NewUnverifiedUserRepository(DB, config.UnverifiedUserCollection)
	usecase := usecases.NewSignupUseCase(repo, unverRepo, time.Duration(config.ContextTimeout)*time.Second, passwordService , ratelimiter)
	signup := controllers.SignupController{
		SignupUsecase: usecase,
	}

	backgroundTask := infrastructure.NewBackgroundTask(usecase)
	
	go backgroundTask.StartCronJob()
	


	SignupRoute.POST("/signup", signup.Signup)
	SignupRoute.POST("/signup/verify", ratelimiter.RateLimitMiddleware() , signup.VerifyOTP)
	SignupRoute.POST("/reset", signup.ForgotPassword)
	SignupRoute.POST("/resend", signup.ResendOTP)

}
