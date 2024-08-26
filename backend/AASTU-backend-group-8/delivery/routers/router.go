package routers

import (
	"meleket/infrastructure"
	"meleket/repository"
	"meleket/usecases"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InitRoutes(r *gin.Engine, client *mongo.Client) {

	r.MaxMultipartMemory = 8 << 20

	jwtService := infrastructure.NewJWTService(os.Getenv("JWT_SECRET"), "Kal", os.Getenv("JWT_REFRESH_SECRET"))

	aiService := infrastructure.NewAIService()
	aiUsecase := usecases.NewAIUsecase(aiService)
	AIRouter(r, aiUsecase, jwtService)

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

	userUsecase := usecases.NewUserUsecase(userRepo, tokenRepo, jwtService)
	otpUsecase := usecases.NewOTPUsecase(otpRepo, userRepo)

	NewUserRouter(r, userUsecase, jwtService, otpUsecase)

	tokenUsecase := usecases.NewTokenUsecase(tokenRepo, jwtService)
	NewRefreshTokenRouter(r, userUsecase, tokenUsecase, jwtService)

	blogUsecase := usecases.NewBlogUsecase(blogRepo)
	NewBlogRouter(r, blogUsecase, jwtService)

	profileUsecase := usecases.NewProfileUsecase(profileRepo)
	NewProfileRoutes(r, profileUsecase, jwtService)

}
