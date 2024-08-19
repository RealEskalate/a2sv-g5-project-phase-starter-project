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

	blogrepo := repositories.NewBlogRepository(client)
	bloguse := usecase.NewBlogUsecase(blogrepo, time.Second*300)

	likerepo := repositories.NewLikeRepository(client)
	likeuse := usecase.NewLikeUsecase(likerepo, time.Second*300)

	dislrepo := repositories.NewDislikeRepository(client)
	disluse := usecase.NewDislikeUsecase(dislrepo, time.Second*300)

	commrepo := repositories.NewCommentRepository(client)
	commuse := usecase.NewCommentUsecase(commrepo, time.Second*300)

	blogcont := controllers.NewBlogController(bloguse, likeuse, commuse, disluse)

	userrepo := repositories.NewUserRepository(client)
	useruse := usecase.NewUserUsecase(userrepo, time.Second*300)
	usercont := controllers.NewUserController(useruse)

	//the router gateway
	r := gin.Default()
	router.SetRouter(r, blogcont, usercont, client)
	r.Run("localhost:8080")

}
