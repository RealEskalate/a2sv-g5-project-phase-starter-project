package route

import (
	"AAiT-backend-group-6/delivery/controller"
	"AAiT-backend-group-6/domain"
	"AAiT-backend-group-6/mongo"
	"AAiT-backend-group-6/repository"
	"AAiT-backend-group-6/usecase"

	"github.com/gin-gonic/gin"
)

// func NewSignupRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
// 	emailService := infrastructure.NewEmailService(env.SmtpServer, env.Mail, env.MailPassword)
// 	ur := repository.NewUserRepository(db, domain.UserCollection)
// 	su := usecase.NewSignupUsecase(ur, timeout, *emailService)
// 	sc := controller.SignupController{
// 		SignupUsecase: su,
// 		Env:           env,
// 	}

// 	group.POST("/signup", sc.Signup)
// }

func NewBlogRouter(db mongo.Database, gin *gin.Engine) {
	tr := repository.NewBlogRepository(db, domain.CollectionBlogs)
	tu := usecase.NewBlogUseCase(tr)
	tc := controller.BlogController{
		BlogUsecase: tu,
	}
	// protectedRoute := gin.Group("")
	publicRoute := gin.Group("")
	// protectedRoute.Use(infrastructure.AdminOnlyMiddleware(), infrastructure.JWTAuthMiddleware())
	publicRoute.GET("/blogs", tc.GetBlogs)
	publicRoute.GET("/blogs/:id", tc.GetBlog)
	publicRoute.POST("/blogs", tc.CreateBlog)
	publicRoute.PUT("/blogs/:id", tc.UpdateBlog)
	publicRoute.DELETE("/blogs/:id", tc.DeleteBlog)
	publicRoute.POST("/blogs/:id/like", tc.LikeBlog)
	publicRoute.POST("/blogs/:id/unlike", tc.UnlikeBlog)
	publicRoute.POST("/blogs/:id/comment", tc.CommentBlog)

}
