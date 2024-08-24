package main

import (
	"blog-api/delivery/controller/blog_controller"
	"blog-api/delivery/controller/user_controller"
	"blog-api/delivery/router"
	"blog-api/infrastructure"
	"blog-api/infrastructure/auth"
	"blog-api/infrastructure/bootstrap"
	"blog-api/infrastructure/email"
	"blog-api/repository/blog_repository"
	"blog-api/repository/comment_repository"
	"blog-api/repository/like_repository"
	"blog-api/repository/refresh_token_repository"
	"blog-api/repository/reset_token_repository"
	"blog-api/repository/user_repository"
	"blog-api/usecase/blog_usecase"
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

	blogRepo := blog_repository.NewBlogRepository(db.Collection("blogs"))
	commRepo := comment_repository.NewCommentRepository(db.Collection("comments"))
	likeRepo := like_repository.NewLikeRepository(db.Collection("likes"))
	aiService := infrastructure.NewGenAIService()

	blogUsecase := blog_usecase.NewBlogUsecase(blogRepo, commRepo, likeRepo, aiService, time.Duration(env.ContextTimeout))
	blogController := blog_controller.NewBlogController(blogUsecase)

	r := gin.Default()
	router.SetRouter(r, blogController, userController, env)
	r.Run(env.ServerAddress)

	fmt.Println("Work correctly", env.DBName)
}
