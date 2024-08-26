package main

import (
	"astu-backend-g1/config"
	_ "astu-backend-g1/delivery/docs"
	"astu-backend-g1/delivery/routers"
	"astu-backend-g1/infrastructure"
	"log"

	"github.com/gin-gonic/gin"
)

// @title TODO APIs
// @version 1.0
// @description Testing Swagger APIs.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @securityDefinitions.apiKey JWT
// @in header
// @name token
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8000
// @BasePath /
// @schemes http
func main() {
	config, err := config.LoadConfig()
	// clientOptions := options.Client().ApplyURI(config.Database.Uri)
	// if err != nil {
	// 	panic(err)
	// }
	// client, err := mongo.Connect(context.TODO(), clientOptions)
	// if err != nil {
	// 	panic(err)
	// }
	// blogCollections := client.Database("BlogAPI").Collection("Blogs")
	// userCollections := client.Database("BlogAPI").Collection("Users")
	// _ = client.Database("BlogAPI").Collection("Tokens")
	// blogRepo := repository.NewBlogRepository(mongoifc.WrapCollection(blogCollections))
	// blogUsecase := usecase.NewBlogUsecase(blogRepo)
	// blogController := controllers.NewBlogController(*blogUsecase)
	// authController := infrastructure.NewAuthController(blogRepo)
	// userRepo := repository.NewUserRepository(mongoifc.WrapCollection(userCollections))
	// userUsecase, err := usecase.NewUserUsecase(userRepo)
	// if err != nil {
	// 	panic(err)
	// }
	// UserController := controllers.NewUserController(userUsecase)
	// Router := routers.NewMainRouter(*UserController, *blogController, authController)
	// Router.GinBlogRouter()
	router := gin.Default()
	prompts, err := infrastructure.LoadPrompt("./prompts.json")
	if err != nil {
		log.Fatal(err)
	}
	routers.AddAIRoutes(router, *config, prompts)
	router.Run(":9000")
}
