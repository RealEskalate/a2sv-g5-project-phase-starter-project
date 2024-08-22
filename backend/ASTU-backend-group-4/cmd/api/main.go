package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func getHost() string {
	hostUrl := os.Getenv("HOST_URL")
	if hostUrl != "" {
		return hostUrl
	}

	return "localhost:8000"
}

func main() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// routers.Setup(c.Database(os.Getenv("MONGO_DB")), r)

	r.Run(getHost())
}
