package main

import (
	"AAiT-backend-group-6/bootstrap"
	"time"

	"time"

	"github.com/gin-gonic/gin"

	"github.com/gin-gonic/gin"
)


func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	route.Setup(env, timeout, db, gin)

	gin.Run(env.ServerAddress)
}





// func main() {
// 	// let's create a new database connection
// 	db := NewMongoDatabase()
// 	llm := infrastructure.NewLlmClient(utils.MESSAGE_TELL_ROLE)
// 	// let's create a new repository
// 	repo := repository.NewAIRepository(db.Database("chatbot"))
// 	// let's create a new usecase
// 	usecase := usecase.NewChatUseCase(repo, llm)
// 	// let's create a new controller
// 	controller := delivery.NewAIController(usecase)

// 	// set up router
// 	r := gin.Default()
// 	//
// 	r.GET("/chat", controller.GetChats)
// 	r.GET("/chat/:id", controller.GetChat)
// 	r.POST("/chat", controller.CreateChat)
// 	r.PUT("/chat/:id", controller.UpdateChat)
// 	// r.DELETE("/chat/:id", controller.DeleteChat)

// 	r.Run() // listen and serve on

// }

// func NewMongoDatabase() mongo.Client {
// 	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
// 	defer cancel()

// 	mongodbURI :=  "mongodb://localhost:27017"

// 	client, err := mongo.NewClient(mongodbURI)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = client.Connect(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	err = client.Ping(ctx)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return client
// }