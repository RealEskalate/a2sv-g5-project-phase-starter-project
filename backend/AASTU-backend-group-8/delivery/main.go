package main

import (
	"log"
	"meleket/bootstrap"
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

	userCollection := repository.NewMongoCollection(client.Database("Blog").Collection("Users"))
	blogCollection := repository.NewMongoCollection(client.Database("Blog").Collection("Blogs"))
	tokenCollection := repository.NewMongoCollection(client.Database("Blog").Collection("Tokens"))
	otpCollection := repository.NewMongoCollection(client.Database("Blog").Collection("OTPs"))
	commentCollection := repository.NewMongoCollection(client.Database("Blog").Collection("Comments"))
	likeCollection := repository.NewMongoCollection(client.Database("Blog").Collection("Likes"))

	userRepo := repository.NewUserRepository(userCollection)
	blogRepo := repository.NewBlogRepository(blogCollection)
	tokenRepo := repository.NewTokenRepository(tokenCollection)
	otpRepo := repository.NewOtpRepository(otpCollection)
	commentRepo := repository.NewCommentRepository(commentCollection)
	likeRepo := repository.NewLikeRepository(likeCollection)

	jwtService := infrastructure.NewJWTService(os.Getenv("JWT_SECRET"), "Kal", os.Getenv("JWT_REFRESH_SECRET"))

	userUsecase := usecases.NewUserUsecase(userRepo)
	tokenUsecase := usecases.NewTokenUsecase(tokenRepo)
	blogUsecase := usecases.NewBlogUsecase(blogRepo)
	otpUsecase := usecases.NewOTPUsecase(otpRepo)
	commentUsecase := usecases.NewCommentUsecase(commentRepo)
	likeUsecase := usecases.NewLikeUsecase(likeRepo)

	r := gin.Default()
	routers.InitRoutes(r, blogUsecase, userUsecase, likeUsecase, commentUsecase, tokenUsecase, otpUsecase, jwtService)

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
