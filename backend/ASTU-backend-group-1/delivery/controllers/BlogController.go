package controllers

import (
	usecase "astu-backend-g1/usecases"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	usecase usecase.BlogUsecase
}

func NewBlogController(uc usecase.BlogUsecase) *BlogController {
	return &BlogController{usecase: uc}
}
func (cont *BlogController) HandleCreateBlog(ctx *gin.Context) {
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	authorId := ctx.PostForm("author_id")
	date := ctx.PostForm("date")
	tags := ctx.PostForm("tags")
	blog, err := cont.usecase.CreateBLog(title, content, authorId, date, tags)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, blog)
	}
}
func (cont *BlogController) HandleGetAllBlogs(ctx *gin.Context) {

	blogs, err := cont.usecase.GetAllBlogs()
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, blogs)
	}

}
func (cont *BlogController) HandleFilterBlogs(ctx *gin.Context) {
	title := ctx.PostForm("title")
	blogid := ctx.PostForm("_id")
	authorId := ctx.PostForm("author_id")
	date := ctx.PostForm("date")
	tags := ctx.PostForm("tags")

	strLike := ctx.PostForm("likes")
	likeSort := 0
	if strLike != "" {
		val, err := strconv.Atoi(strLike)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		}
		likeSort = val

	}
	strDislike := ctx.PostForm("dislikes")
	dislikeSort := 0
	if strDislike != "" {
		val, err := strconv.Atoi(strDislike)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		}
		dislikeSort = val
	}

	strComment := ctx.PostForm("comments")
	commentSort := 0
	if strComment != "" {
		val, err := strconv.Atoi(strComment)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		}
		commentSort = val
	}

	strView := ctx.PostForm("views")
	viewSort := 0
	if strView != "" {
		val, err := strconv.Atoi(strView)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		}
		viewSort = val
	}

	blogs, err := cont.usecase.FilterBlogs(title, blogid, date, tags, authorId, likeSort, dislikeSort, commentSort, viewSort)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, blogs)
	}

}
func (cont *BlogController) HandleUpdate(ctx *gin.Context) {
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	authorId := ctx.PostForm("author_id")
	// date := ctx.PostForm("date")
	tags := ctx.PostForm("tags")
	blog, err := cont.usecase.UpdateBLog(ctx.Param("blogId"), title, content, authorId, tags, "", "", "")
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, blog)
	}

}
func (cont *BlogController) HandleDelete(ctx *gin.Context) {
	err := cont.usecase.DeleteBLog(ctx.Param("blogId"))
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Blog deleted"})
	}

}
func (cont *BlogController) HandleBlogLikeOrDislike(ctx *gin.Context) {
	interactionType := ctx.Param("type")
	if interactionType == "1" {
		fmt.Println("like some blog")
		err := cont.usecase.LikeBlog(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("authorId"))
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Blog liked successfully"})
		}
	} else if interactionType == "-1" {
		err := cont.usecase.DislikeBlog(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("authorId"))
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Blog disliked successfully"})
		}
	} else {
		err := cont.usecase.ViewBlogs(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("authorId"))
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Blog viewed successfully"})
		}
	}
}

func (cont *BlogController) HandleCommentLikeOrDislike(ctx *gin.Context) {
	interactionType := ctx.Param("type")
	if interactionType == "1" {
		err := cont.usecase.LikeComment(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("commentId"), ctx.Request.FormValue("authorId"))
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Comment liked successfully"})
		}
	} else if interactionType == "-1" {
		err := cont.usecase.DislikeComment(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("commentId"), ctx.Request.FormValue("authorId"))
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Comment disliked successfully"})
		}
	} else {
		err := cont.usecase.ViewComment(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("commentId"), ctx.Request.FormValue("authorId"))
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Comment viewed successfully"})
		}
	}
}

func (cont *BlogController) HandleReplyLikeOrDislike(ctx *gin.Context) {
	like := ctx.Param("like")
	if like == "1" {
		err := cont.usecase.LikeReply(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("commentId"), ctx.Request.FormValue("replyId"), ctx.Request.FormValue("authorId"))
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Reply liked successfully"})
		}
	} else if like == "-1" {
		err := cont.usecase.DislikeReply(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("commentId"), ctx.Request.FormValue("replyId"), ctx.Request.FormValue("authorId"))
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Reply disliked successfully"})
		}
	} else {
		err := cont.usecase.ViewReply(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("commentId"), ctx.Request.FormValue("replyId"), ctx.Request.FormValue("authorId"))
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Reply viewed successfully"})
		}
	}
}

// func (cont *BlogController) HandleDislike(ctx *gin.Context) {
// 	cont.usecase.DislikeBlog(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("authorId"))

// }
// func (cont *BlogController) HandleView(ctx *gin.Context) {
// 	cont.usecase.ViewsBlogs(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("authorId"))

// }
func (cont *BlogController) HandleCommentOnBlog(ctx *gin.Context) {
	err := cont.usecase.AddComment(ctx.Request.FormValue("content"), ctx.Request.FormValue("blogId"), ctx.Request.FormValue("authorId"))
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
	}

}
