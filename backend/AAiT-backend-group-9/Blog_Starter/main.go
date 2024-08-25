package main

import (
	"log"
	"time"

	route "Blog_Starter/api/router"
	"Blog_Starter/config"
	"Blog_Starter/repository"
	"Blog_Starter/usecase"

	"Blog_Starter/domain"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	env := config.NewEnv()

	// Initialize MongoDB client
	db := config.NewMongoDatabase(env)
	defer config.CloseMongoDBConnection(db)
	database := db.Database(env.DBName)

	// Set up the Gin router

	router := gin.Default()
	router.Static("/upload", "./upload")


	ur := repository.NewUserRepository(database, domain.CollectionUser)
	timeout := time.Duration(env.ContextTimeout) * time.Second
	userUsecase := usecase.NewUserUsecase(ur, timeout)
	// Start the routine to clean up expired users every 24 hours
	userUsecase.StartExpiredUserCleanup(24 * time.Hour)
	

	// Global CORS configuration
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"*"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type"}
	router.Use(cors.New(config))

	// Set up routes
	route.Setup(env, timeout, db, router)

	// Run the server
	if err := router.Run("localhost:" + env.DBPort); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
