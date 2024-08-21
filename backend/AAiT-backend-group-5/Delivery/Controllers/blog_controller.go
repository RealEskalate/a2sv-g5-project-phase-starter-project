package controllers

import (
	"net/http"
	"strconv"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/gin-gonic/gin"
)

type blogController struct {
	usecase interfaces.BlogUsecase
}

func NewBlogController(usecase interfaces.BlogUsecase) interfaces.BlogController {
	return &blogController{
		usecase: usecase,
	}
}

func (c *blogController) getAuthorID(ctx *gin.Context) string {
	return ctx.GetString("id")
}

func (c *blogController) CreateBlogController(ctx *gin.Context) {
	var newBlog dtos.CreateBlogRequest

	if err := ctx.ShouldBind(&newBlog); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	authorID := c.getAuthorID(ctx)

	blog := models.Blog{
		Title:    newBlog.Title,
		Content:  newBlog.Content,
		Tags:     newBlog.Tags,
		AuthorID: authorID,
	}

	createdBlog, err := c.usecase.CreateBlog(ctx, &blog)
	if err != nil {
		ctx.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"data": createdBlog})

}

func (c *blogController) GetBlogController(ctx *gin.Context) {
	blogId := ctx.Param("id")

	blog, err := c.usecase.GetBlog(ctx, blogId)

	if err != nil {
		ctx.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": blog})
}

func (c *blogController) GetBlogsController(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")

	pageInt, ok := strconv.Atoi(page)

	if ok != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	blogs, err := c.usecase.GetBlogs(ctx, pageInt)

	if err != nil {
		ctx.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": blogs})
}

func (c *blogController) SearchBlogsController(ctx *gin.Context) {
	var filter dtos.FilterBlogRequest

	if err := ctx.ShouldBind(&filter); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	blogs, err := c.usecase.SearchBlogs(ctx, filter)

	if err != nil {
		ctx.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": blogs})

}
func (c *blogController) UpdateBlogController(ctx *gin.Context) {
	var updateBlog dtos.UpdateBlogRequest
	blogID := ctx.Param("id")

	if err := ctx.ShouldBind(&updateBlog); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	authorID := c.getAuthorID(ctx)

	blog := models.Blog{
		ID:       blogID,
		AuthorID: authorID,
		Title:    updateBlog.Title,
		Content:  updateBlog.Content,
		Tags:     updateBlog.Tags,
	}

	err := c.usecase.UpdateBlog(ctx, blog.ID, &blog)

	if err != nil {
		ctx.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Blog updated successfully"})
}

func (c *blogController) DeleteBlogController(ctx *gin.Context) {
	blogID := ctx.Param("id")
	authorID := c.getAuthorID(ctx)

	if authorID == "" {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	deleteBlogReq := dtos.DeleteBlogRequest{
		BlogID:   blogID,
		AuthorID: authorID,
	}

	if err := c.usecase.DeleteBlog(ctx, deleteBlogReq); err != nil {
		ctx.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

func (c *blogController) TrackPopularityController(ctx *gin.Context) {
	var blogPopularity dtos.TrackPopularityRequest

	if err := ctx.ShouldBind(&blogPopularity); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	userID := c.getAuthorID(ctx)

	blogPopularity.UserID = userID
	if err := c.usecase.TrackPopularity(ctx, blogPopularity); err != nil {
		ctx.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Popularity tracked successfully"})
}

func (c *blogController) AddCommentController(ctx *gin.Context) {
	var comment models.Comment

	if err := ctx.ShouldBind(&comment); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	userID := c.getAuthorID(ctx)
	comment.UserID = userID
	if err := c.usecase.AddComment(ctx, comment); err != nil {
		ctx.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Comment added successfully"})
}
