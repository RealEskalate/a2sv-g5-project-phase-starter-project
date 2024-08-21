package main

import (
	"log"
	"meleket/bootstrap"
	"meleket/delivery/controllers"
	"meleket/delivery/routers"
	"meleket/infrastructure"
	"meleket/repository"
	"meleket/usecases"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	client, err := bootstrap.InitMongoDB()
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Define the collections
	userCollection := client.Database("Blog").Collection("Users")
	blogCollection := client.Database("Blog").Collection("Blogs")
	tokenCollection := client.Database("Blog").Collection("Tokens")
	otpCollection := client.Database("Blog").Collection("OTPs")
	profileCollection := client.Database("Blog").Collection("Profile")
	commentCollection := client.Database("Blog").Collection("Comments")
	likeDislikeCollection := client.Database("Blog").Collection("LikesDislikes")

	// Repositories
	userRepo := repository.NewUserRepository(repository.NewMongoCollection(userCollection))
	blogRepo := repository.NewBlogRepository(repository.NewMongoCollection(blogCollection))
	tokenRepo := repository.NewTokenRepository(repository.NewMongoCollection(tokenCollection))
	otpRepo := repository.NewOtpRepository(repository.NewMongoCollection(otpCollection))
	profileRepo := repository.NewProfileRepository(repository.NewMongoCollection(profileCollection))
	commentRepo := repository.NewCommentRepository(repository.NewMongoCollection(commentCollection))
	likeDislikeRepo := repository.NewLikeDislikeRepository(likeDislikeCollection)

	// Services
	jwtService := infrastructure.NewJWTService(os.Getenv("JWT_SECRET"), "Kal", os.Getenv("JWT_REFRESH_SECRET"))

	// Usecases
	userUsecase := usecases.NewUserUsecase(userRepo, tokenRepo, jwtService)
	tokenUsecase := usecases.NewTokenUsecase(tokenRepo, jwtService)
	blogUsecase := usecases.NewBlogUsecase(blogRepo)
	otpUsecase := usecases.NewOTPUsecase(otpRepo, userRepo)
	profileUsecase := usecases.NewProfileUsecase(profileRepo)
	commentUsecase := usecases.NewCommentUsecase(commentRepo)
	likeDislikeUsecase := usecases.NewLikeDislikeUsecase(likeDislikeRepo)

	// Handlers
	profileHandler := controllers.NewProfileHandler(profileUsecase)

	// Routers
	r := gin.Default()
	profileRouter := routers.NewProfileRouter(profileHandler, r)
	profileRouter.InitProfileRoutes(r.Group("/api/v1"))
	routers.InitRoutes(r, blogUsecase, userUsecase, tokenUsecase, jwtService, likeDislikeUsecase, commentUsecase, tokenUsecase, otpUsecase)

	// Start the server
	if err := r.Run(os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
