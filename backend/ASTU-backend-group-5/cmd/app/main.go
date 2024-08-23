package main

import (
	"blogApp/internal/config"
	"blogApp/internal/http/routes"
	"blogApp/pkg/mongo"

	"github.com/gin-gonic/gin"
)

func main() {
	// os.Clearenv()
	confs, err := config.Load()
	if err != nil {
		panic(err)
	}
	mongo.ConnectDB(confs.MONGO_URI)
	defer mongo.DisconnectDB()
	mongo.InitializeCollections()

	router := gin.Default()
	// Serve static files from the "upload" directory
	router.Static("/uploads", "./uploads")
	routes.SetUpRoute(router)
	router.Run(":" + confs.GO_PORT)
}
