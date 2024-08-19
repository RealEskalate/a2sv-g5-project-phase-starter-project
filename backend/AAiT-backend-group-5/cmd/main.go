package main

import (
	"time"

	config "github.com/aait.backend.g5.main/backend/Config"
	routers "github.com/aait.backend.g5.main/backend/Delivery/Routers"
	"github.com/gin-gonic/gin"
)

func main() {

	env := config.NewEnv()
	timeout := time.Duration(env.CONTEXT_TIMEOUT) * time.Second

	// Initialize MongoDB connection
	client := config.ConnectDB(env)
	db := config.GetDatabase(client, env)

	gin := gin.Default()

	routers.Setup(env, timeout, *db, gin)
	gin.Run(env.SERVER_ADDRESS)
}
