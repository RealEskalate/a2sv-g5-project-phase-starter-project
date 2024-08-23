package main

import (
	"context"
	"log"

	"aait.backend.g10/delivery/config"
	"aait.backend.g10/delivery/controllers"
	"aait.backend.g10/delivery/router"
	"aait.backend.g10/infrastructures"
	"aait.backend.g10/repositories"
	"aait.backend.g10/usecases"
	"github.com/go-redis/redis/v8"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// load environment variables
	err := config.LoadEnvironmentVariables()
	if err != nil {
		log.Fatal(err.Error())
		return
	}


	clientOption := options.Client().ApplyURI(config.ENV.DB_URI)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}
	db := client.Database(config.ENV.DB_NAME)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     config.ENV.REDIS_ADDR,
		Password: config.ENV.REDIS_PASSWORD,
		DB:       0,
	})
	_, err = redisClient.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal(err)
	}

	aiService := infrastructures.NewAIService(config.ENV.GEMINI_API_KEY)
	jwtService := infrastructures.JwtService{JwtSecret: config.ENV.JWT_SECRET}
	pwdService := infrastructures.HashingService{}
	emailService := infrastructures.EmailService{
		AppEmail:    config.ENV.EMAIL,
		AppPass:     config.ENV.EMAIL_PASSWORD,
		AppUsername: config.ENV.EMAIL_USERNAME,
		AppHost:     config.ENV.EMAIL_HOST,
	}

	cacheRepo := infrastructures.NewCacheRepo(redisClient, context.Background())

	userRepo := repositories.NewUserRepository(db, config.ENV.USER_COLLECTION)
	userUseCase := usecases.NewUserUseCase(userRepo)
	userController := controllers.NewUserController(userUseCase)

	authUsecases := usecases.NewAuthUsecase(userRepo, &jwtService, &pwdService, &emailService)
	authController := controllers.NewAuthController(authUsecases, config.GoogleOAuthConfig)

	commentRepo := repositories.NewCommentRepository(db, config.ENV.COMMENT_COLLECTION_NAME)
	commentController := &controllers.CommentController{
		CommentUsecase: usecases.NewCommentUsecase(commentRepo, userRepo, cacheRepo),
	}

	likeRepo := repositories.NewLikeRepository(db, config.ENV.LIKE_COLLECTION_NAME)
	likeController := &controllers.LikeController{
		LikeUseCase: usecases.NewLikeUseCase(likeRepo, cacheRepo),
	}

	blogRepo := repositories.NewBlogRepository(db, config.ENV.BLOG_COLLECTION)
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
