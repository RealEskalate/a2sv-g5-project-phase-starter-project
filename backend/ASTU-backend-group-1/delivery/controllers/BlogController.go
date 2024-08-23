package controllers

import (
	"astu-backend-g1/domain"
	"astu-backend-g1/infrastructure"
	usecase "astu-backend-g1/usecases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type BlogController struct {
	usecase usecase.BlogUsecase
}

func NewBlogController(uc usecase.BlogUsecase) *BlogController {
	return &BlogController{usecase: uc}
}

func (cont *BlogController) HandleCreateBlog(ctx *gin.Context) {
	var blog domain.Blog
	err := ctx.ShouldBindJSON(&blog)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	claims, err := infrastructure.GetClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the claims"})
		return
	}
	blog.AuthorId = claims.ID
	blog, err = cont.usecase.CreateBLog(blog)
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

func (cont *BlogController) HandleGetBlogById(ctx *gin.Context) {
	blogs, err := cont.usecase.GetBlogByBLogId(ctx.Param("blogId"))
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, blogs)
	}
}

func (cont *BlogController) HandleGetPopularBlog(ctx *gin.Context) {
	blogs, err := cont.usecase.FindPopularBlog()
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, blogs)
	}
}

func (cont *BlogController) HandleFilterBlogs(ctx *gin.Context) {
	var blf domain.BlogFilterOption
	err := ctx.ShouldBindJSON(&blf)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}

	blogs, err := cont.usecase.FilterBlogs(blf)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, blogs)
	}
}

// todo: start tesing  here
func (cont *BlogController) HandleBlogUpdate(ctx *gin.Context) {
	var updateBlog domain.Blog
	err := ctx.ShouldBindJSON(&updateBlog)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	blog, err := cont.usecase.UpdateBLog(ctx.Param("blogId"), updateBlog)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, blog)
	}
}

func (cont *BlogController) HandleBlogDelete(ctx *gin.Context) {
	claims, err := infrastructure.GetClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the claims"})
		return
	}
	err = cont.usecase.DeleteBLog(ctx.Param("blogId"), claims.ID)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Blog deleted"})
	}
}

func (cont *BlogController) HandleBlogLikeOrDislike(ctx *gin.Context) {
	interactionType := ctx.Param("type")
	blogId := ctx.Param("blogId")
	claims, err := infrastructure.GetClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the claims"})
		return
	}
	if interactionType == "like" {
		err := cont.usecase.LikeBlog(blogId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Blog liked successfully"})
		}
	} else if interactionType == "dislike" {
		err := cont.usecase.DislikeBlog(blogId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Blog disliked successfully"})
		}
	} else if interactionType == "view" {
		err := cont.usecase.ViewBlogs(blogId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Blog viewed successfully"})
		}
	} else {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid interaction type"})
	}
}

func (cont *BlogController) HandleCommentOnBlog(ctx *gin.Context) {
	var newComment domain.Comment
	err := ctx.ShouldBindJSON(&newComment)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	claims, err := infrastructure.GetClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the claims"})
		return
	}
	newComment.AuthorId = claims.ID
	err = cont.usecase.AddComment(ctx.Param("blogId"), newComment)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
	}
}

func (cont *BlogController) HandleGetAllComments(ctx *gin.Context) {
	comments, err := cont.usecase.GetAllComments(ctx.Param("blogId"))
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, comments)
	}
}

func (cont *BlogController) HandleGetCommentById(ctx *gin.Context) {
	comments, err := cont.usecase.GetCommentById(ctx.Param("blogId"), ctx.Param("commentId"))
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, comments)
	}
}

func (cont *BlogController) HandleCommentLikeOrDislike(ctx *gin.Context) {
	interactionType := ctx.Param("type")
	commentId := ctx.Param("commentId")
	blogId := ctx.Param("blogId")
	claims, err := infrastructure.GetClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the claims"})
		return
	}

	if interactionType == "like" {
		err := cont.usecase.LikeComment(blogId, commentId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Comment liked successfully"})
		}
	} else if interactionType == "dislike" {
		err := cont.usecase.DislikeComment(blogId, commentId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Comment disliked successfully"})
		}
	} else if interactionType == "view" {
		err := cont.usecase.ViewComment(blogId, commentId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Comment viewed successfully"})
		}
	} else {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid interaction type"})
	}
}

func (cont *BlogController) HandleReplyOnComment(ctx *gin.Context) {
	var newReply domain.Reply
	err := ctx.ShouldBindJSON(&newReply)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	claims, err := infrastructure.GetClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the claims"})
		return
	}
	newReply.AuthorId = claims.ID
	err = cont.usecase.ReplyToComment(ctx.Param("blogId"), ctx.Param("commentId"), newReply)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Reply added successfully"})
	}
}

func (cont *BlogController) HandleGetAllRepliesForComment(ctx *gin.Context) {
	replies, err := cont.usecase.GetAllRepliesForComment(ctx.Param("blogId"), ctx.Param("commentId"))
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, replies)
	}
}

func (cont *BlogController) HandleGetReplyById(ctx *gin.Context) {
	replies, err := cont.usecase.GetReplyById(ctx.Param("blogId"), ctx.Param("commentId"), ctx.Param("replyId"))
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, replies)
	}
}

func (cont *BlogController) HandleReplyLikeOrDislike(ctx *gin.Context) {
	like := ctx.Param("type")
	commentId := ctx.Param("commentId")
	blogId := ctx.Param("blogId")
	replyId := ctx.Param("replyId")
	claims, err := infrastructure.GetClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the claims"})
		return
	}
	if like == "like" {
		err := cont.usecase.LikeReply(blogId, commentId, replyId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Reply liked successfully"})
		}
	} else if like == "dislike" {
		err := cont.usecase.DislikeReply(blogId, commentId, replyId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Reply disliked successfully"})
		}
	} else if like == "view" {
		err := cont.usecase.ViewReply(blogId, commentId, replyId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Reply viewed successfully"})
		}
	} else {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "invalid interaction type"})
	}
}