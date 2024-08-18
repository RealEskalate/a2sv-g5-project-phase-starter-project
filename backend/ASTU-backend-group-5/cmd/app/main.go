package main

import (
	"blogApp/internal/config"
	"blogApp/internal/http/routes"
	"blogApp/pkg/mongo"

	"github.com/gin-gonic/gin"
)

func main() {
	confs, err := config.Load()
	if err != nil {
		panic(err)
	}
	mongo.ConnectDB(confs.MONGO_URI)
	defer mongo.DisconnectDB()
	mongo.InitializeCollections()

	userCollection := mongo.UserCollection

	router := gin.Default()

	routes.SetUpRoute(router, userCollection)
	router.Run(confs.APP_DOMAIN)
}
