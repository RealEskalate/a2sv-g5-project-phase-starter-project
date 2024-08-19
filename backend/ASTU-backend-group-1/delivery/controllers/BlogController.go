package controllers

import (
	usecase "astu-backend-g1/usecases"
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
	cont.usecase.CreateBLog(title, content, authorId, date, tags)

}
func (cont *BlogController) HandleGetAllBlogs(ctx *gin.Context) {
	
	blogs,err:= cont.usecase.GetAllBlogs()
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound,err)
	}else{
		ctx.IndentedJSON(http.StatusOK,blogs)
	}


}
func (cont *BlogController) HandleFilterBlogs(ctx *gin.Context) {
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
	blogs,err:= cont.usecase.FilterBlogs(title, blogid, date, tags, authorId, likeSort, DislikeSort, CommentSort, ViewSort)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound,err)
	}else{
		ctx.IndentedJSON(http.StatusOK,blogs)
	}

}
func (cont *BlogController) HandleUpdate(ctx *gin.Context) {
	title := ctx.PostForm("title")
	content := ctx.PostForm("content")
	authorId := ctx.PostForm("author_id")
	// date := ctx.PostForm("date")
	tags := ctx.PostForm("tags")
	blog,err := cont.usecase.UpdateBLog(ctx.Request.FormValue("blogId"), title, content, authorId, tags, "", "", "")
	if err!= nil {
        ctx.IndentedJSON(http.StatusNotFound,err)
    }else{
        ctx.IndentedJSON(http.StatusOK,blog)
    }

}
func (cont *BlogController) HandleDelete(ctx *gin.Context) {
	err := cont.usecase.DeleteBLog(ctx.Request.FormValue("blogId"))
	if err!= nil {
        ctx.IndentedJSON(http.StatusNotFound,err)
    }else{
        ctx.IndentedJSON(http.StatusOK,gin.H{"message":"Blog deleted"})
    }

}
func (cont *BlogController) HandleBlogLikeOrDislike(ctx *gin.Context) {
	like := ctx.PostForm("like")
	if like == "1" {
		err := cont.usecase.LikeBlog(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("authorId"))
		if err!= nil {
            ctx.IndentedJSON(http.StatusNotFound,err)
        }else{
			ctx.IndentedJSON(http.StatusOK,gin.H{"message":"Blog liked successfully"})
		}
	} else if like == "-1" {
		err := cont.usecase.DislikeBlog(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("authorId"))
		if err!= nil {
            ctx.IndentedJSON(http.StatusNotFound,err)
        } else{
            ctx.IndentedJSON(http.StatusOK,gin.H{"message":"Blog disliked successfully"})
        }
	} else {
		err:=cont.usecase.ViewBlogs(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("authorId"))
		if err!= nil {
            ctx.IndentedJSON(http.StatusNotFound,err)
        } else{
            ctx.IndentedJSON(http.StatusOK,gin.H{"message":"Blog viewed successfully"})
        }
	}
}

func (cont *BlogController) HandleCommentLikeOrDislike(ctx *gin.Context) {
	like := ctx.PostForm("like")
	if like == "1" {
		err:=cont.usecase.LikeComment(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("commentId"), ctx.Request.FormValue("authorId"))
		if err!= nil {
            ctx.IndentedJSON(http.StatusNotFound, err)
        } else {
            ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Comment liked successfully"})
        }
	} else if like == "-1" {
		err := cont.usecase.DislikeComment(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("commentId"), ctx.Request.FormValue("authorId"))
		if err!= nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		}else{
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Comment disliked successfully"})
		}
	} else {
		err:=cont.usecase.ViewComment(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("commentId"), ctx.Request.FormValue("authorId"))
		if err!= nil {
            ctx.IndentedJSON(http.StatusNotFound, err)
        } else{
            ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Comment viewed successfully"})
        }
	}
}

func (cont *BlogController) HandleReplyLikeOrDislike(ctx *gin.Context) {
	like := ctx.PostForm("like")
	if like == "1" {
		err:=cont.usecase.LikeReply(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("commentId"), ctx.Request.FormValue("replyId"), ctx.Request.FormValue("authorId"))
		if err!= nil {
            ctx.IndentedJSON(http.StatusNotFound, err)
        } else{
            ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Reply liked successfully"})
        }
	} else if like == "-1" {
		err := cont.usecase.DislikeReply(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("commentId"), ctx.Request.FormValue("replyId"), ctx.Request.FormValue("authorId"))
		if err!= nil {
            ctx.IndentedJSON(http.StatusNotFound, err)
        } else{
            ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Reply disliked successfully"})
        }
	} else {
		err := cont.usecase.ViewReply(ctx.Request.FormValue("blogId"), ctx.Request.FormValue("commentId"), ctx.Request.FormValue("replyId"), ctx.Request.FormValue("authorId"))
		if err!= nil {
            ctx.IndentedJSON(http.StatusNotFound, err)
        } else{
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
	err:=cont.usecase.AddComment(ctx.Request.FormValue("content"), ctx.Request.FormValue("blogId"), ctx.Request.FormValue("authorId"))
	if err!= nil {
        ctx.IndentedJSON(http.StatusNotFound, err)
    } else{
        ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
    }

}
