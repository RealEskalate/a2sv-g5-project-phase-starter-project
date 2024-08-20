package main

import (
	"AAiT-backend-group-6/bootstrap"
	"AAiT-backend-group-6/delivery/routers"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	app := bootstrap.App()

	env := app.Env

	db := app.Mongo.Database(env.DBName)

	defer func() {
		if r := recover(); r != nil {
			log.Fatalf("Application panicked: %v", r)
		}
		app.CloseDBConnection()
	}()

	log.Println("Starting the application...")

	// timeout := time.Duration(env.ContextTimeout) * time.Second

	gin := gin.Default()

	log.Println("Setting up routers...")
	routers.SetupRouter(db, gin)

	log.Println("Running the server...")
	err := gin.Run(env.ServerAddress)
	if err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
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
