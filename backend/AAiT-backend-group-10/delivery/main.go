package main

import (
	"context"
	"log"
	"os"

	"aait.backend.g10/repositories"
	"aait.backend.g10/usecases"
	"aait.backend.g10/delivery/controllers"
	"aait.backend.g10/delivery/router"
	"aait.backend.g10/infrastructures"
	"github.com/go-redis/redis/v8"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	clientOption := options.Client().ApplyURI(os.Getenv("DB_URI"))
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(os.Getenv("DB_NAME"))
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})
	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}

	email := os.Getenv("EMAIL")
	password := os.Getenv("EMAIL_PASSWORD")
	username := os.Getenv("EMAIL_USERNAME")
	host := os.Getenv("EMAIL_HOST")

	cacheRepo := infrastructures.NewCacheRepo(redisClient, context.Background())

	aiService := infrastructures.NewAIService(os.Getenv("GEMINI_API_KEY"))
	jwtService := infrastructures.JwtService{JwtSecret: os.Getenv("JWT_SECRET")}
	pwdService := infrastructures.HashingService{}
	emailService := infrastructures.EmailService{
		AppEmail:    email,
		AppPass:     password,
		AppUsername: username,
		AppHost:     host,
	}

	userRepo := repositories.NewUserRepository(db, os.Getenv("USER_COLLECTION"))
	userUseCase := usecases.NewUserUseCase(userRepo)
	userController := controllers.NewUserController(userUseCase)

	authUsecases := usecases.NewAuthUsecase(userRepo, &jwtService, &pwdService, &emailService)
	authController := controllers.NewAuthController(authUsecases, controllers.GoogleOAuthConfig)

	commentRepo := repositories.NewCommentRepository(db, os.Getenv("COMMENT_COLLECTION_NAME"))
	commentController := &controllers.CommentController{
		CommentUsecase: usecases.NewCommentUsecase(commentRepo, userRepo, cacheRepo),
	}

	likeRepo := repositories.NewLikeRepository(db, os.Getenv("LIKE_COLLECTION_NAME"))
	likeController := &controllers.LikeController{
		LikeUseCase: usecases.NewLikeUseCase(likeRepo, cacheRepo),
	}

	blogRepo := repositories.NewBlogRepository(db, os.Getenv("BLOG_COLLECTION"))
	blogUseCase := usecases.NewBlogUseCase(blogRepo, userRepo, likeRepo, commentRepo, aiService, cacheRepo)
	blogController := controllers.NewBlogController(blogUseCase)

	router.NewRouter(
		db,
		redisClient,
		router.RouterControllers{
			CommentController: commentController,
			LikeController:    likeController,
			AuthController:    authController,
			UserController:    userController,
			BlogController:    blogController,
		},
		router.RouterServices{
			JwtService: &jwtService,
		},
	)
}
