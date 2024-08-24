package main

import (
	"AAIT-backend-group-3/internal/delivery/controllers"
	"AAIT-backend-group-3/internal/delivery/routers"
	"AAIT-backend-group-3/internal/infrastructures/database"
	"AAIT-backend-group-3/internal/infrastructures/middlewares"
	"AAIT-backend-group-3/internal/infrastructures/services"
	"AAIT-backend-group-3/internal/usecases"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	repositories "AAIT-backend-group-3/internal/repositories/implementation"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/google/generative-ai-go/genai"
	"google.golang.org/api/option"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbName := os.Getenv("DB_NAME")
	secretKey := os.Getenv("SECRET_KEY")
	smtpPortStr := os.Getenv("SMTP_PORT")
	userName := os.Getenv("USERNAME")
	smtpHost := os.Getenv("SMTP_HOST")
	passWord := os.Getenv("PASSWORD")
	geminiApiKey := os.Getenv("GEMINI_API_KEY")
	
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Fatalf("Invalid SMTP_PORT: %v", err)
	}

	dbClient, err := database.NewMongoDBClient(context.Background(), dbName)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	fmt.Println("Connected to MongoDB!", dbClient.Database.Name())

	// create a new gemini client
	gemini_client, err := genai.NewClient(context.Background(), option.WithAPIKey(geminiApiKey))
	if err != nil {
		log.Fatal(err)
	}


	//services
	emailSvc := services.NewEmailService(smtpHost, smtpPort, userName, passWord)
	passSvc := services.NewPasswordService()
	validationSvc := services.NewValidationService()
	jwtSvc := services.NewJWTService(secretKey)
	cacheSvc := services.NewCacheService("localhost:6379", "", 0)
	cloudSvc := services.NewCloudinaryService(os.Getenv("CLOUDINARY_NAME"), os.Getenv("CLOUDINARY_API_KEY"), os.Getenv("CLOUDINARY_API_SECRET"), os.Getenv("CLOUDINARY_UPLOAD_FOLDER"),)
	aiService := services.NewAiService(gemini_client)

	//repositories
	userRepo := repositories.NewMongoUserRepository(dbClient.Database, "users", cacheSvc)
	otpRepo := repositories.NewMongoOtpRepository(dbClient.Database, "otps")
	blogRepo := repositories.NewMongoBlogRepository(dbClient.Database, "blogs",cacheSvc)
	commentRepo := repositories.NewMongoCommentRepository(dbClient.Database, "comments")
	tagRep := repositories.NewMongoTagRepository(dbClient.Database, "tags")

	//middlewares
	authMiddleware := middlewares.NewAuthMiddleware(jwtSvc, cacheSvc)


	//usecases
	userUsecase := usecases.NewUserUsecase(userRepo, passSvc, validationSvc, emailSvc, jwtSvc, cloudSvc)
	otpUsecase := usecases.NewOtpUseCase(otpRepo, userRepo, emailSvc, passSvc, "http://localhost:8080", validationSvc)
	blogService := usecases.NewBlogUsecase(blogRepo, tagRep)
	commentService := usecases.NewCommentUsecase(commentRepo)
	aiHelperUsecase := usecases.NewAiHelperUsecase(aiService)


	// controllers
	userController := controllers.NewUserController(userUsecase)
	otpController := controllers.NewOTPController(otpUsecase)
	blogController := controllers.NewBlogController(blogService)
	commentController := controllers.NewCommentController(commentService)
	aiHelperController := controllers.NewAiHelperController(aiHelperUsecase)


	router := gin.New()
	router.Use(gin.Logger())


	// routers
	routers.CreateUserRouter(router, userController, otpController, authMiddleware)
	routers.CreateBlogRouter(router, blogController, authMiddleware)
	routers.CreateCommentRouter(router, commentController, authMiddleware)
	routers.CreateAiHelperRouter(router, aiHelperController)
	if err := router.Run(":" + os.Getenv("PORT")); err!= nil{
		log.Fatal(err)
	}
}
