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

	userCollection := client.Database("Blog").Collection("Users")
	blogCollection := client.Database("Blog").Collection("Blogs")
	tokenCollection := client.Database("Blog").Collection("Tokens")
	otpCollection := client.Database("Blog").Collection("OTPs")
	profileCollection := client.Database("Blog").Collection("profile")

	userMockCollection := repository.NewMongoCollection(userCollection)
	blogMockCollection := repository.NewMongoCollection(blogCollection)
	tokenMockCollection := repository.NewMongoCollection(tokenCollection)
	otpMockCollection := repository.NewMongoCollection(otpCollection)
	profileMockCollection := repository.NewMongoCollection(profileCollection)

	userRepo := repository.NewUserRepository(userMockCollection)
	blogRepo := repository.NewBlogRepository(blogMockCollection)
	tokenRepo := repository.NewTokenRepository(tokenMockCollection)
	otpRepo := repository.NewOtpRepository(otpMockCollection)
	profileRepo := repository.NewProfileRepository(profileMockCollection)

	jwtService := infrastructure.NewJWTService(os.Getenv("JWT_SECRET"), "Kal", os.Getenv("JWT_REFRESH_SECRET"))

	userUsecase := usecases.NewUserUsecase(userRepo, tokenRepo, jwtService)
	tokenUsecase := usecases.NewTokenUsecase(tokenRepo, jwtService)
	blogUsecase := usecases.NewBlogUsecase(blogRepo)
	otpUsecase := usecases.NewOTPUsecase(otpRepo, userRepo)
	profileUsecase := usecases.NewProfileUsecase(profileRepo)

	profileHandler := controllers.NewProfileHandler(profileUsecase)

	// passwordService := infrastructure.NewPasswordService()

	r := gin.Default()
	profileRouter := routers.NewProfileRouter(profileHandler, r)
	profileRouter.InitProfileRoutes(r.Group("/api/v1"))
	routers.InitRoutes(r, blogUsecase, userUsecase, tokenUsecase, otpUsecase, jwtService)
	// routers.InitRoutes(r, blogUsecase, userUsecase, tokenUsecase, otpUsecase)

	if err := r.Run(os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
