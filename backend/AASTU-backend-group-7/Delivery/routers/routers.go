package routers

import (
	config "blogapp/Config"
	"blogapp/Domain"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func Setuprouter() {
	client := config.ConnectDB()
	log.Println(client)

	DataBase := client.Database("Blog")
	BlogCollections := Domain.Collections{

		Users: DataBase.Collection("Users"),
		Blogs: DataBase.Collection("Blogs"),
	}
	
	fmt.Println(BlogCollections)

	router = gin.Default()
}
