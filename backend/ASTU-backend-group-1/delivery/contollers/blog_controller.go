package contollers

import (
	"astu-backend-g1/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogController struct{
	usecase usecases.BlogUsecase
}

func NewBlogController(uc usecases.BlogUsecase) *BlogController{
    return &BlogController{usecase: uc}
}
func (cont *BlogController) handleCreateBlog(ctx *gin.Context){
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	authorId := ctx.PostForm("author_id")
	date := ctx.PostForm("date")
	tags := ctx.PostForm("tags")
	cont.usecase.CreateBLog(title, content, authorId, date, tags)

}
func (cont *BlogController) handleGetAllBlogs(ctx *gin.Context){
	cont.usecase.GetAllBlogs()
	
}
func (cont *BlogController) handleFilterBlogs(ctx *gin.Context){
	title := ctx.PostForm("title")
	blogid := ctx.PostForm("_id")
	authorId := ctx.PostForm("author_id")
	date := ctx.PostForm("date")
	tags := ctx.PostForm("tags")

	likeSort,err := strconv.Atoi(ctx.PostForm("likes"))
	if err != nil{
		panic(err)
	}
	DislikeSort,err := strconv.Atoi(ctx.PostForm("dislikes"))
	if err != nil{
		panic(err)
	}
	CommentSort,err := strconv.Atoi(ctx.PostForm("comments"))
	if err != nil{
		panic(err)
	}
	ViewSort,err := strconv.Atoi(ctx.PostForm("views"))
	if err != nil{
		panic(err)
	}
	cont.usecase.FilterBlogs(title,blogid,date,tags,authorId,likeSort,DislikeSort,CommentSort,ViewSort)
	
}
func (cont *BlogController) handleUpdate(ctx *gin.Context){
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	authorId := ctx.PostForm("author_id")
	// date := ctx.PostForm("date")
	tags := ctx.PostForm("tags")
	cont.usecase.UpdateBLog(ctx.Request.FormValue("blogId"),title,content,authorId,tags,"","","")
	
}
func (cont *BlogController) handleDelete(ctx *gin.Context){
	cont.usecase.DeleteBLog(ctx.Request.FormValue("blogId"))
	
}
func (cont *BlogController) handleLike(ctx *gin.Context){
	cont.usecase.LikeBlog(ctx.Request.FormValue("blogId"),ctx.Request.FormValue("authorId"))
	
}
func (cont *BlogController) handleDislike(ctx *gin.Context){
	cont.usecase.DislikeBlog(ctx.Request.FormValue("blogId"),ctx.Request.FormValue("authorId"))
	
}
func (cont *BlogController) handleView(ctx *gin.Context){
	cont.usecase.ViewsBlogs(ctx.Request.FormValue("blogId"),ctx.Request.FormValue("authorId"))
	
}
func (cont *BlogController) handleCommentOnBlog(ctx *gin.Context){
	cont.usecase.AddComment(ctx.Request.FormValue("content"),ctx.Request.FormValue("blogId"),ctx.Request.FormValue("authorId"))
	
}