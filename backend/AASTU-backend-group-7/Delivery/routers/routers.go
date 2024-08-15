package routers

import (
	config "blogapp/Config"
	"blogapp/Domain"
	"context"
	"log"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Setuprouter() *gin.Engine {
	// Initialize the database
	client := config.ConnectDB()
	log.Println(client)
	DataBase := client.Database("Blog")

	// Defer the closing of the database
	defer func() {
		if err := client.Disconnect(context.TODO()); err != nil {
			log.Fatal(err)
		}
	}()

	// Initialize the collections
	BlogCollections := Domain.Collections{

		Users: DataBase.Collection("Users"),
		Blogs: DataBase.Collection("Blogs"),
	}

	// Initialize the router
	router = gin.Default()
	return router
}
