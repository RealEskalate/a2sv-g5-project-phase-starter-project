package main

import (
	"time"

	config "github.com/aait.backend.g5.main/backend/Config"
	routers "github.com/aait.backend.g5.main/backend/Delivery/Routers"
	"github.com/gin-gonic/gin"
)

func main() {

	// Initialize MongoDB connection
	client := config.ConnectDB()
	db := config.GetDatabase(client)

	env := config.NewEnv()
	timeout := time.Duration(env.CONTEXT_TIMEOUT) * time.Second

	gin := gin.Default()

	routers.Setup(env, timeout, *db, gin)
	gin.Run(env.SERVER_ADDRESS)
}
