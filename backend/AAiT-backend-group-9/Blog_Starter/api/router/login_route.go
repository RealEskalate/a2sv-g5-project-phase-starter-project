package router

import (
	"time"

	"Blog_Starter/api/controller"
	"Blog_Starter/config"
	"Blog_Starter/domain"
	"Blog_Starter/repository"
	"Blog_Starter/usecase"
	"Blog_Starter/utils/infrastructure"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

// NewLoginRouter sets up the login routes.
func NewLoginRouter(env *config.Env, timeout time.Duration, db *mongo.Client, group *gin.RouterGroup) {
    // Initialize the database
    database := db.Database(env.DBName)
    
    // Initialize repositories
    or := repository.NewOtpRepository(database, domain.CollectionOTP)
    ur := repository.NewUserRepository(database, domain.CollectionUser)
    tm := &infrastructure.NewTokenManager{} // Assuming NewTokenManager returns an implementation of TokenManager
    // Initialize use cases
    UserUsecase := usecase.NewUserUsecase(ur, timeout)
    loginUsecase := usecase.NewLoginUseCase(ur,tm,timeout,env)
    otpUsecase := usecase.NewOtpUsecase(or, timeout)
    
    // Initialize controller
    loginController := controller.NewLoginController(loginUsecase,otpUsecase,UserUsecase)
    
    // Set up routes
    group.POST("/login", loginController.Login)
    group.POST("/forgotpassword", loginController.ForgotPassword)
    group.PATCH("/updatepassword", loginController.UpdatePassword)
}
