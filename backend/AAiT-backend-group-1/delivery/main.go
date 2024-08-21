package main

import (
	"context"
	"log"
	"os"

	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/delivery/controllers"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/delivery/router"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/infrastructure"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/infrastructure/mail"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/repository"
	"github.com/RealEskalate/a2sv-g5-project-phase-starter-project/aait-backend-group-1/usecases"
	"github.com/joho/godotenv"
)

func main() {
	
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	
	geminiApiKey := os.Getenv("GEMINI_API_KEY")

	databaseUri := os.Getenv("DATABASE_URI")
	databaseName := os.Getenv("DATABASE_NAME")
	databaseService := infrastructure.NewDatabaseService(databaseUri, databaseName)

	accessSecret := os.Getenv("ACCESS_SECRET")
	refreshSecret := os.Getenv("REFRESH_SECRET")
	jwtService := infrastructure.NewJWTTokenService(accessSecret, refreshSecret, databaseService.GetCollection("tokens"))

	cacheDbUri := os.Getenv("CACHE_DB_URI")
	cacheDbPassword := os.Getenv("CACHE_DB_PASSWORD")
	cacheService := infrastructure.NewCacheService(cacheDbUri, cacheDbPassword, 0)

	passwordService := infrastructure.NewPasswordService()

	emailService := mail.NewEmailService()
	
	sessionRepo := repository.NewSessionRespository(databaseService.GetCollection("sessions"))

	userRepo := repository.NewUserRespository(databaseService.GetCollection("users"))
	userUC := usecases.NewUserUseCase(userRepo, sessionRepo, passwordService, jwtService, emailService, cacheService)
	userController := controllers.NewUserController(userUC)
	
	blogRepo := repository.NewBlogRepository(databaseService.GetCollection("blogs"), context.TODO())
	blogUC := usecases.NewBlogUseCase(blogRepo)
	blogController := controllers.NewBlogController(blogUC)

	blogAssistantUC := usecases.NewBlogAssistantUsecase(geminiApiKey)
	blogAssistantController := controllers.NewBlogAssistantController(blogAssistantUC)

	authMiddleware := infrastructure.NewMiddlewareService(jwtService, cacheService)

	r := router.SetupRouter(userController, blogController, blogAssistantController, authMiddleware)
	if err := r.Run(); err != nil {
		log.Fatal(err)
	}

}
