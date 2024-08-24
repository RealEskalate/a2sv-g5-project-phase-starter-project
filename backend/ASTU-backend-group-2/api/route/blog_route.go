package route

import (
	"time"

	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/api/controller"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/bootstrap"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/domain"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/repository"
	"github.com/a2sv-g5-project-phase-starter-project/backend/ASTU-backend-group-2/usecase"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func NewProtectedBlogsRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	br := repository.NewBlogRepository(*db, domain.CollectionBlog)
	cr := repository.NewCommentRepository(db)
	rr := repository.NewReactionRepository(db)

	bc := controller.BlogController{
		BlogUsecase:     usecase.NewBlogUsecase(br, timeout),
		Env:             env,
		CommentUsecase:  usecase.NewCommentUsecase(cr, timeout),
		ReactionUsecase: usecase.NewReactionUsecase(rr, br, timeout),
	}

	group.POST("/blogs", bc.CreateBlog())
	group.POST("/blogs/batch", bc.BatchCreateBlog())
	group.PUT("/blogs/:id", bc.UpdateBlog())
	group.PATCH("/blogs/:id", bc.UpdateBlog())
	group.DELETE("/blogs/:id", bc.DeleteBlog())

	// comments
	group.POST("/blogs/:id/comments", bc.CreateComment())

	// // only authenticated users can access
	group.PUT("/comments/:comment_id", bc.UpdateComment())
	group.DELETE("/comments/:comment_id", bc.DeleteComment())
	//like and dislike
	group.POST("/blogs/:id/like", bc.Like())
	group.POST("/blogs/:id/dislike", bc.Dislike())
}

func NewPublicBlogsRouter(env *bootstrap.Env, timeout time.Duration, db *mongo.Database, group *gin.RouterGroup) {
	br := repository.NewBlogRepository(*db, domain.CollectionBlog)
	cr := repository.NewCommentRepository(db)
	// rr := repository.NewReactionRepository(db)
	bc := controller.BlogController{
		BlogUsecase:    usecase.NewBlogUsecase(br, timeout),
		Env:            env,
		CommentUsecase: usecase.NewCommentUsecase(cr, timeout),
		// ReactionUsecase: usecase.NewReactionUsecase(rr, br, timeout),
	}

	group.GET("/blogs", bc.GetBlogs())
	group.GET("/blogs/:id", bc.GetBlog())
	group.GET("blogs/popular", bc.GetbyPopularity())

	group.GET("/blogs/tags/", bc.GetByTags())
	group.GET("/blogs/search/", bc.Search())
	group.GET("/blogs/recent", bc.SortByDate())
	//comments
	group.GET("/blogs/:id/comments", bc.GetComments())
	group.GET("/comments/:comment_id", bc.GetComment())
}
