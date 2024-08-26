package main

import (
	"blog_g2/deliveries/controllers"
	"blog_g2/deliveries/router"
	"blog_g2/infrastructure"
	"blog_g2/repositories"
	"blog_g2/usecase"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	client := infrastructure.MongoDBInit() //mongodb initialization
	infrastructure.InitGoogleOAuthConfig()
	aiserv, _ := infrastructure.NewGeminiAIService()

	blogrepo := repositories.NewBlogRepository(client)
	bloguse := usecase.NewBlogUsecase(blogrepo, aiserv, time.Second*300)

	likerepo := repositories.NewLikeRepository(client)
	likeuse := usecase.NewLikeUsecase(likerepo, time.Second*300)

	dislrepo := repositories.NewDislikeRepository(client)
	disluse := usecase.NewDislikeUsecase(dislrepo, time.Second*300)

	commrepo := repositories.NewCommentRepository(client)
	commuse := usecase.NewCommentUsecase(commrepo, aiserv, time.Second*300)
	comcont := controllers.NewCommentController(commuse)

	blogcont := controllers.NewBlogController(bloguse, likeuse, commuse, disluse, aiserv)

	userrepo := repositories.NewUserRepository(client)
	useruse := usecase.NewUserUsecase(userrepo, time.Second*300)
	usercont := controllers.NewUserController(useruse)
	oauthController := controllers.NewOAuthController(useruse)

	//the router gateway
	r := gin.Default()
	router.SetRouter(r, comcont, blogcont, usercont, oauthController, client)
	r.Run("127.0.0.1:8080")

}
