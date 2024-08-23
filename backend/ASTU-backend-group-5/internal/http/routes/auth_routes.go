package routes

import (
	"blogApp/internal/http/handlers/token"
	"blogApp/internal/repository/mongodb"
	"blogApp/internal/usecase"
	localmongo "blogApp/pkg/mongo"

	"github.com/gin-gonic/gin"
)

func RegisterVerificationRoutes(router *gin.Engine) {

	// userRepo := &mongodb.UserRepositoryMongo{Collection: collection}
	// userUsecase := user.NewUserUsecase(userRepo)
	userHandler := InstantaiteUserHandler()
	authRoutes := router.Group("/api/v1/auth")

	{
		authRoutes.POST("/login", userHandler.Login)
		authRoutes.POST("/register", userHandler.Register)
		authRoutes.GET("/google_redirect", userHandler.GoogleCallback)
		authRoutes.POST("/verify/request", userHandler.RequestVerifyEmail)
		authRoutes.GET("/verify/confirm", userHandler.VerifyEmail) //I used this naming to make things clear
		authRoutes.POST("/reset-password/request", userHandler.ResetPasswordRequest)
		authRoutes.POST("/reset-password/confirm", userHandler.ResetPassword)

		//logout and refresh
		tokenRepo := mongodb.NewMongoTokenRepository(localmongo.TokenCollection)
		tokenUsecase := usecase.NewTokenUsecase(tokenRepo)
		TokenHandler := token.NewTokenHandler(*tokenUsecase)

		r2 := authRoutes.Group("/")
		r2.POST("/logout", TokenHandler.LogOut)
		r2.POST("/refresh", TokenHandler.RefreshToken)
	}
}
