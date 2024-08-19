package contollers

import (
	"astu-backend-g1/usecases"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	usecase usecases.BlogUsecase
}

func NewBlogController(uc usecases.BlogUsecase) *BlogController {
	return &BlogController{usecase: uc}
}
func (cont *BlogController) handleCreateBlog(ctx *gin.Context) {
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	authorId := ctx.PostForm("author_id")
	date := ctx.PostForm("date")
	tags := ctx.PostForm("tags")
	cont.usecase.CreateBLog(title, content, authorId, date, tags)

}
func (cont *BlogController) handleGetAllBlogs(ctx *gin.Context) {
	cont.usecase.GetAllBlogs()

}
func (cont *BlogController) handleFilterBlogs(ctx *gin.Context) {
	title := ctx.PostForm("title")
	blogid := ctx.PostForm("_id")
	authorId := ctx.PostForm("author_id")
	date := ctx.PostForm("date")
	tags := ctx.PostForm("tags")

	likeSort, err := strconv.Atoi(ctx.PostForm("likes"))
	if err != nil {
		panic(err)
	}
	DislikeSort, err := strconv.Atoi(ctx.PostForm("dislikes"))
	if err != nil {
		panic(err)
	}
	CommentSort, err := strconv.Atoi(ctx.PostForm("comments"))
	if err != nil {
		panic(err)
	}
	ViewSort, err := strconv.Atoi(ctx.PostForm("views"))
	if err != nil {
		panic(err)
	}
	cont.usecase.FilterBlogs(title, blogid, date, tags, authorId, likeSort, DislikeSort, CommentSort, ViewSort)

}
func (cont *BlogController) handleUpdate(ctx *gin.Context) {
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	authorId := ctx.PostForm("author_id")
	// date := ctx.PostForm("date")
	tags := ctx.PostForm("tags")
	cont.usecase.UpdateBLog(ctx.Request.FormValue("blogId"), title, content, authorId, tags, "", "", "")

}
func (cont *BlogController) handleDelete(ctx *gin.Context) {
	cont.usecase.DeleteBLog(ctx.Request.FormValue("blogId"))

}
func (cont *BlogController) handleBlogLikeOrDislike(ctx *gin.Context) {
	like := ctx.PostForm("like")
	if like == "1" {
		cont.usecase.LikeBlog(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("authorId"))
	} else if like == "-1" {
		cont.usecase.DislikeBlog(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("authorId"))
	} else {
		cont.usecase.ViewBlogs(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("authorId"))
	}
}

func (cont *BlogController) handleCommentLikeOrDislike(ctx *gin.Context) {
	like := ctx.PostForm("like")
	if like == "1" {
		cont.usecase.LikeComment(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("commentId"), ctx.Request.FormValue("authorId"))
	} else if like == "-1" {
		cont.usecase.DislikeComment(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("commentId"), ctx.Request.FormValue("authorId"))
	} else {
		cont.usecase.ViewComment(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("commentId"), ctx.Request.FormValue("authorId"))
	}
}

func (cont *BlogController) handleReplyLikeOrDislike(ctx *gin.Context) {
	like := ctx.PostForm("like")
	if like == "1" {
		cont.usecase.LikeReply(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("commentId"), ctx.Request.FormValue("replyId"), ctx.Request.FormValue("authorId"))
	} else if like == "-1" {
		cont.usecase.DislikeReply(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("commentId"), ctx.Request.FormValue("replyId"), ctx.Request.FormValue("authorId"))
	} else {
		cont.usecase.ViewReply(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("commentId"), ctx.Request.FormValue("replyId"), ctx.Request.FormValue("authorId"))
	}
}

// func (cont *BlogController) handleDislike(ctx *gin.Context) {
// 	cont.usecase.DislikeBlog(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("authorId"))

// }
// func (cont *BlogController) handleView(ctx *gin.Context) {
// 	cont.usecase.ViewsBlogs(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("authorId"))

// }
func (cont *BlogController) handleCommentOnBlog(ctx *gin.Context) {
	cont.usecase.AddComment(ctx.Request.FormValue("content"), ctx.Request.FormValue("blogId"), ctx.Request.FormValue("authorId"))

}
