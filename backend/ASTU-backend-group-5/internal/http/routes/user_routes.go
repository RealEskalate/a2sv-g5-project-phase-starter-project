package routes

import (
    "blogApp/internal/http/handlers"
    "blogApp/internal/repository/mongodb"
    "blogApp/internal/usecase/user"
    "blogApp/pkg/mongo"
    "log"
    "os"

    "github.com/gin-gonic/gin"
    "github.com/joho/godotenv"
)

func SetupRouter() *gin.Engine {
    err := godotenv.Load(".env")
    if err != nil {
        log.Fatalf("Error loading .env file: %s", err)
    }
    uri := os.Getenv("MONGO_URI")
    // Connect to the database
    client, err := mongo.ConnectDB(uri)
    if err != nil {
        log.Fatal(err)
    }
    db := client.Database("blogDB")
    if err != nil {
        log.Fatal(err)
    }

    // Get the collection
    userRepo := &mongodb.UserRepositoryMongo{Collection: db.Collection("users")}
    loginUseCase := user.NewLoginUseCase(userRepo)
    loginHandler := handlers.NewLoginHandler(loginUseCase)

    // Setup Gin router
    r := gin.Default()

    // Register login route
    r.POST("/api/v1/auth/login", loginHandler.Login)

    return r
}