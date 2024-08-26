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
	commentCollection := repository.NewMongoCollection(client.Database("Blog").Collection("Comments"))
	// likeCollection := repository.NewMongoCollection(client.Database("Blog").Collection("Likes"))

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
	commentRepo := repository.NewCommentRepository(commentCollection)
	// likeRepo := repository.NewLikeRepository(likeCollection)

	jwtService := infrastructure.NewJWTService(os.Getenv("JWT_SECRET"), "Kal", os.Getenv("JWT_REFRESH_SECRET"))

	aiService := infrastructure.NewAIService()
	aiUsecase := usecases.NewAIUsecase(aiService)

	userUsecase := usecases.NewUserUsecase(userRepo, tokenRepo, jwtService)
	tokenUsecase := usecases.NewTokenUsecase(tokenRepo, jwtService)
	blogUsecase := usecases.NewBlogUsecase(blogRepo)
	otpUsecase := usecases.NewOTPUsecase(otpRepo, userRepo)
	profileUsecase := usecases.NewProfileUsecase(profileRepo)
	commentUsecase := usecases.NewCommentUsecase(commentRepo)
	// likeUsecase := usecases.NewLikeUsecase(likeRepo)

	//controllers
	aiHandler := controllers.NewAIHandler(aiUsecase)
	profileHandler := controllers.NewProfileHandler(profileUsecase)
	promoteDemoteHandler := controllers.NewPromoteDemoteController(userUsecase)

	r := gin.Default()

	routers.AIRouter(r, aiHandler, jwtService)
	routers.PromoteDemoteRouter(r, promoteDemoteHandler, jwtService)

	profileRouter := routers.NewProfileRouter(profileHandler, r)
	profileRouter.InitProfileRoutes(r.Group("/api/v1"))
	routers.InitRoutes(r, blogUsecase, userUsecase, tokenUsecase, jwtService, commentUsecase, tokenUsecase, otpUsecase)
	// routers.InitRoutes(r, blogUsecase, userUsecase, tokenUsecase, otpUsecase)

	if err := r.Run(os.Getenv("PORT")); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
