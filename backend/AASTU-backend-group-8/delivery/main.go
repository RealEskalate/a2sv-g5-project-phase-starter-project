package main

import (
	"context"
	"log"
	"meleket/bootstrap"
	"meleket/infrastructure"
	"meleket/repository"
	"meleket/routers"
	"meleket/usecases"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

	userCollection := client.Database("Blog").Collection("Users")
	blogCollection := client.Database("Blog").Collection("Blogs")
	tokenCollection := client.Database("Blog").Collection("Tokens")
	otpCollection := client.Database("Blog").Collection("OTPs")

	userMockCollection := repository.NewMongoCollection(userCollection)
	blogMockCollection := repository.NewMongoCollection(blogCollection)
	tokenMockCollection := repository.NewMongoCollection(tokenCollection)
	otpMockCollection := repository.NewMongoCollection(otpCollection)

	userRepo := repository.NewUserRepository(userMockCollection)
	blogRepo := repository.NewBlogRepository(blogMockCollection)
	tokenRepo := repository.NewTokenRepository(tokenMockCollection)
	otpRepo := repository.NewOTPRepository(otpMockCollection)
	
	jwtService := infrastructure.NewJWTService(os.Getenv("JWT_SECRET"),"kal", os.Getenv("JWT_REFRESH_SECRET"))
	
	userUsecase := usecases.NewUserUsecase(userRepo,jwtService)
	tokenUsecase := usecases.NewTokenRepository(tokenRepo)
	blogUsecase := usecases.NewBlogRepository(blogRepo)
	otpUsecase := usecases.NewotpRepository(otpRepo)

	// passwordService := infrastructure.NewPasswordService()


	r := gin.Default()
	routers.InitRoutes(r, blogUsecase, userUsecase, tokenUsecase, otpUsecase, jwtService)

	if err := r.Run(); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}