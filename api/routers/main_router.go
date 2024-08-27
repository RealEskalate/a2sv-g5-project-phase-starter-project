package routers

import (
	"loan-management/config"
	"loan-management/database"
	"log"

	"github.com/gin-gonic/gin"
)

func Run() {
	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}
	db, err := database.NewMongoDatabase(config)
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	AddUserRoutes(router, db)
	router.Run(config.Server.Port)
}
