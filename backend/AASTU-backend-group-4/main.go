package main

import (
	"blog-api/delivery/controller/user_controller"
	"blog-api/infrastructure/auth"
	"blog-api/infrastructure/bootstrap"
	"blog-api/infrastructure/email"
	"blog-api/repository/refresh_token_repository"
	"blog-api/repository/reset_token_repository"
	"blog-api/repository/user_repository"
	"blog-api/usecase/user_usecase"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()
	defer app.CloseDBConnection()
	env := app.Env

	db := app.Mongo.Database(env.DBName)

	userCollection := db.Collection("users")
	refreshTokenCollection := db.Collection("refresh-tokens")
	resetTokenCollection := db.Collection("reset-tokens")

	userRepo := user_repository.NewUserRepository(userCollection)
	refreshTokenRepo := refresh_token_repository.NewRefreshTokenRepository(refreshTokenCollection)
	resetTokenRepo := reset_token_repository.NewResetTokenRepository(resetTokenCollection)

	authService := auth.NewAuthService(refreshTokenRepo, resetTokenRepo, env.AccessTokenSecret, env.RefreshTokenSecret, env.ResetTokenSecret, env.AccessTokenExpiryHour, env.RefreshTokenExpiryHour, env.ResetTokenExpiryHour)

	emailService := email.NewEmailService(env.SMTPServer, env.SMTPPort, env.SMTPUser, env.SMTPPassword, env.FromAddress)

	userUsecase := user_usecase.NewUserUsecase(userRepo, authService, emailService, time.Duration(env.ContextTimeout))

	userController := user_controller.NewUserController(userUsecase, authService, env)

	router := gin.Default()

	// Register routes

	router.POST("/signup", userController.SignUp)
	router.POST("/login", userController.Login)
	router.POST("/refresh", userController.RefreshTokens)
	router.GET("/logout", auth.JwtAuthMiddleware(env.AccessTokenSecret), userController.Logout)
	router.POST("/forgot-password", userController.ForgotPassword)
	router.POST("/reset-password", userController.ResetPassword)
	// Run the server
	router.Run(env.ServerAddress)

	fmt.Println("Work correctly", env.DBName)
}
