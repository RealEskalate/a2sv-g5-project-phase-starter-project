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

	"AAIT-backend-group-3/internal/repositories/implementation"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	
	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Fatalf("Invalid SMTP_PORT: %v", err)
	}

	dbClient, err := database.NewMongoDBClient(context.Background(), dbName)
	if err != nil {
		log.Fatalf("Error connecting to MongoDB: %v", err)
	}
	fmt.Println("Connected to MongoDB!", dbClient.Database.Name())

	//repositories
	userRepo := repositories.NewMongoUserRepository(dbClient.Database, "users")
	otpRepo := repositories.NewMongoOtpRepository(dbClient.Database, "otps")
	blogRepo := repositories.NewMongoBlogRepository(dbClient.Database, "blogs")
	commentRepo := repositories.NewMongoCommentRepository(dbClient.Database, "comments")

	//services
	emailSvc := services.NewEmailService(smtpHost, smtpPort, userName, passWord)
	passSvc := services.NewPasswordService()
	validationSvc := services.NewValidationService()
	jwtSvc := services.NewJWTService(secretKey)

	//middlewares
	authMiddleware := middlewares.NewAuthMiddleware(jwtSvc)

	//usecases
	userUsecase := usecases.NewUserUsecase(userRepo, passSvc, validationSvc, emailSvc, jwtSvc)
	otpUsecase := usecases.NewOtpUseCase(otpRepo, userRepo, emailSvc, passSvc, "http://localhost:8080")
	blogService := usecases.NewBlogUsecase(blogRepo)
	commentService := usecases.NewCommentUsecase(commentRepo)

	// controllers
	userController := controllers.NewUserController(userUsecase)
	otpController := controllers.NewOTPController(otpUsecase)
	blogController := controllers.NewBlogController(blogService)
	commentController := controllers.NewCommentController(commentService)

	router := gin.New()
	router.Use(gin.Logger())

	// routers
	routers.CreateUserRouter(router, userController, otpController)
	routers.CreateBlogRouter(router, blogController, authMiddleware)
	routers.CreateCommentRouter(router, commentController, authMiddleware)
	if err := router.Run(":" + os.Getenv("PORT")); err!= nil{
		log.Fatal(err)
	}
	fmt.Println("server connnected")
}
