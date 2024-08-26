package controllers

import (
	"astu-backend-g1/domain"
	"astu-backend-g1/infrastructure"
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

// @Summary Create a new blog
// @Description Create a new blog post with the provided data.
// @Tags blogs
// @Accept json
// @Produce json
// @Param blog body domain.Blog true "Blog data"
// @Success 200 {object} domain.Blog
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /blogs [post]
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

// @Summary Get all blogs
// @Description Retrieve a list of all blogs with pagination.
// @Tags blogs
// @Accept json
// @Produce json
// @Param pageNumber query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {array} domain.Blog
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /blogs [get]
func (cont *BlogController) HandleGetAllBlogs(ctx *gin.Context) {
	page := ctx.Query("pageNumber")
	ipage, err := strconv.Atoi(page)
	if err != nil || ipage < 1 {
		ipage = 1
	}
	pageSize := ctx.Query("pageSize")
	ipageSize, err := strconv.Atoi(pageSize)
	if err != nil || ipageSize < 1 {
		ipageSize = 5
	}
	x := domain.PaginationInfo{}
	x.Page = ipage
	x.PageSize = ipageSize
	blogs, err := cont.usecase.GetAllBlogs(x)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, blogs)
	}
}

// @Summary Get a blog by ID
// @Description Retrieve a specific blog by its ID.
// @Tags blogs
// @Accept json
// @Produce json
// @Param blogId path string true "Blog ID"
// @Success 200 {object} domain.Blog
// @Failure 404 {object} map[string]string
// @Router /blogs/{blogId} [get]
func (cont *BlogController) HandleGetBlogById(ctx *gin.Context) {
	blogs, err := cont.usecase.GetBlogByBLogId(ctx.Param("blogId"))
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, blogs)
	}
}

// @Summary Get popular blogs
// @Description Retrieve the most popular blogs.
// @Tags blogs
// @Accept json
// @Produce json
// @Success 200 {array} domain.Blog
// @Failure 404 {object} map[string]string
// @Router /blogs/popular [get]
func (cont *BlogController) HandleGetPopularBlog(ctx *gin.Context) {
	blogs, err := cont.usecase.FindPopularBlog()
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, blogs)
	}
}

// @Summary Filter blogs
// @Description Filter blogs based on provided filters and pagination.
// @Tags blogs
// @Accept json
// @Produce json
// @Param filter body domain.BlogFilterOption true "Filter and pagination options"
// @Success 200 {array} domain.Blog
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /blogs/filter [post]
func (cont *BlogController) HandleFilterBlogs(ctx *gin.Context) {

	var blf domain.BlogFilterOption
	err := ctx.ShouldBindJSON(&blf)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, err)
		return
	}
	page := ctx.Query("pageNumber")
	ipage, err := strconv.Atoi(page)
	if err != nil || ipage < 1 {
		ipage = 1
	}
	pageSize := ctx.Query("pageSize")
	ipageSize, err := strconv.Atoi(pageSize)
	if err != nil || ipageSize < 1 {
		ipageSize = 5
	}
	x := domain.PaginationInfo{}
	x.Page = ipage
	x.PageSize = ipageSize
	blf.Pagination = x

	blogs, err := cont.usecase.FilterBlogs(blf)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, blogs)
	}
}

// @Summary Update a blog
// @Description Update an existing blog post by its ID.
// @Tags blogs
// @Accept json
// @Produce json
// @Param blogId path string true "Blog ID"
// @Param blog body domain.Blog true "Updated blog data"
// @Success 200 {object} domain.Blog
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /blogs/{blogId} [put]
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

// @Summary Delete a blog
// @Description Delete a blog post by its ID.
// @Tags blogs
// @Accept json
// @Produce json
// @Param blogId path string true "Blog ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /blogs/{blogId} [delete]
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

// @Summary Like or dislike a blog
// @Description Like, dislike, or view a blog post by its ID.
// @Tags blogs
// @Accept json
// @Produce json
// @Param type path string true "Interaction type (like, dislike, view)"
// @Param blogId path string true "Blog ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /blogs/{blogId}/{type} [post]
func (cont *BlogController) HandleBlogLikeOrDislike(ctx *gin.Context) {
	interactionType := ctx.Param("type")
	blogId := ctx.Param("blogId")
	claims, err := infrastructure.GetClaims(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "could not get the claims"})
		return
	}
	if interactionType == "like" {
		message, err := cont.usecase.LikeBlog(blogId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": message, "error": err})
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": message, "error": err})
		}
	} else if interactionType == "dislike" {
		message, err := cont.usecase.DislikeBlog(blogId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": message, "error": err})
		}
	} else if interactionType == "view" {
		message, err := cont.usecase.ViewBlogs(blogId, claims.ID)
		if err != nil {
			ctx.IndentedJSON(http.StatusNotFound, err)
		} else {
			ctx.IndentedJSON(http.StatusOK, gin.H{"message": message, "error": err})
		}
	} else {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "allowed:[like,view,dislike]", "error": "unknown interaction type"})
	}
}

// @Summary Comment on a blog
// @Description Add a comment to a blog post.
// @Tags comments
// @Accept json
// @Produce json
// @Param blogId path string true "Blog ID"
// @Param comment body domain.Comment true "Comment data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /blogs/{blogId}/comments [post]
func (cont *BlogController) HandleCommentOnBlog(ctx *gin.Context) {
	blogId := ctx.Param("blogId")
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
	err = cont.usecase.AddComment(blogId, newComment)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Comment added successfully"})
	}
}

// @Summary Get all comments for a blog
// @Description Retrieve all comments for a specific blog with pagination.
// @Tags comments
// @Accept json
// @Produce json
// @Param blogId path string true "Blog ID"
// @Param pageNumber query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {array} domain.Comment
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /blogs/{blogId}/comments [get]
func (cont *BlogController) HandleGetAllComments(ctx *gin.Context) {
	blogId := ctx.Param("blogId")
	page := ctx.Query("pageNumber")
	ipage, err := strconv.Atoi(page)
	if err != nil || ipage < 1 {
		ipage = 1
	}
	pageSize := ctx.Query("pageSize")
	ipageSize, err := strconv.Atoi(pageSize)
	if err != nil || ipageSize < 1 {
		ipageSize = 5
	}
	x := domain.PaginationInfo{}
	x.Page = ipage
	x.PageSize = ipageSize

	comments, err := cont.usecase.GetAllComments(blogId, x)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, comments)
	}
}

// @Summary Get a comment by ID
// @Description Retrieve a specific comment by its ID.
// @Tags comments
// @Accept json
// @Produce json
// @Param blogId path string true "Blog ID"
// @Param commentId path string true "Comment ID"
// @Success 200 {object} domain.Comment
// @Failure 404 {object} map[string]string
// @Router /blogs/{blogId}/comments/{commentId} [get]
func (cont *BlogController) HandleGetCommentById(ctx *gin.Context) {
	blogId := ctx.Param("blogId")
	commentId := ctx.Param("commentId")
	comments, err := cont.usecase.GetCommentById(blogId, commentId)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, comments)
	}
}

// @Summary Like or dislike a comment
// @Description Like, dislike, or view a comment on a blog post.
// @Tags comments
// @Accept json
// @Produce json
// @Param type path string true "Interaction type (like, dislike, view)"
// @Param blogId path string true "Blog ID"
// @Param commentId path string true "Comment ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /blogs/{blogId}/comments/{commentId}/{type} [post]
func (cont *BlogController) HandleCommentLikeOrDislike(ctx *gin.Context) {
	interactionType := ctx.Param("type")
	blogId := ctx.Param("blogId")
	commentId := ctx.Param("commentId")
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

// @Summary Reply to a comment
// @Description Add a reply to a comment on a blog post.
// @Tags replies
// @Accept json
// @Produce json
// @Param blogId path string true "Blog ID"
// @Param commentId path string true "Comment ID"
// @Param reply body domain.Reply true "Reply data"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /blogs/{blogId}/comments/{commentId}/replies [post]
func (cont *BlogController) HandleReplyOnComment(ctx *gin.Context) {
	var newReply domain.Reply
	blogId := ctx.Param("blogId")
	commentId := ctx.Param("commentId")
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
	err = cont.usecase.ReplyToComment(blogId, commentId, newReply)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Reply added successfully"})
	}
}

// @Summary Get all replies for a comment
// @Description Retrieve all replies for a specific comment with pagination.
// @Tags replies
// @Accept json
// @Produce json
// @Param blogId path string true "Blog ID"
// @Param commentId path string true "Comment ID"
// @Param pageNumber query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {array} domain.Reply
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /blogs/{blogId}/comments/{commentId}/replies [get]
func (cont *BlogController) HandleGetAllRepliesForComment(ctx *gin.Context) {
	blogId := ctx.Param("blogId")
	commentId := ctx.Param("commentId")
	page := ctx.Query("pageNumber")
	ipage, err := strconv.Atoi(page)
	if err != nil || ipage < 1 {
		ipage = 1
	}
	pageSize := ctx.Query("pageSize")
	ipageSize, err := strconv.Atoi(pageSize)
	if err != nil || ipageSize < 1 {
		ipageSize = 5
	}
	x := domain.PaginationInfo{}
	x.Page = ipage
	x.PageSize = ipageSize

	replies, err := cont.usecase.GetAllReplies(blogId, commentId, x)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, replies)
	}
}

// @Summary Get a reply by ID
// @Description Retrieve a specific reply by its ID.
// @Tags replies
// @Accept json
// @Produce json
// @Param blogId path string true "Blog ID"
// @Param commentId path string true "Comment ID"
// @Param replyId path string true "Reply ID"
// @Success 200 {object} domain.Reply
// @Failure 404 {object} map[string]string
// @Router /blogs/{blogId}/comments/{commentId}/replies/{replyId} [get]
func (cont *BlogController) HandleGetReplyById(ctx *gin.Context) {
	blogId := ctx.Param("blogId")
	commentId := ctx.Param("commentId")
	replyId := ctx.Param("replyId")
	replies, err := cont.usecase.GetReplyById(blogId, commentId, replyId)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, err)
	} else {
		ctx.IndentedJSON(http.StatusOK, replies)
	}
}

// @Summary Like or dislike a reply
// @Description Like, dislike, or view a reply to a comment on a blog post.
// @Tags replies
// @Accept json
// @Produce json
// @Param type path string true "Interaction type (like, dislike, view)"
// @Param blogId path string true "Blog ID"
// @Param commentId path string true "Comment ID"
// @Param replyId path string true "Reply ID"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Router /blogs/{blogId}/comments/{commentId}/replies/{replyId}/{type} [post]
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
