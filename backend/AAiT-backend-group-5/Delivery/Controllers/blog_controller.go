package controllers

import (
	"net/http"

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

func (c *blogController) getAuthorID(ctx *gin.Context) (string, bool) {
	authorID, ok := ctx.Get("id")
	if !ok {
		return "", false
	}

	authorIDStr, ok := authorID.(string)
	if !ok {
		return "", false
	}

	return authorIDStr, true
}

func (c *blogController) CreateBlogController(ctx *gin.Context) {
	var newBlog dtos.CreateBlogRequest

	if err := ctx.ShouldBind(&newBlog); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	authorID, ok := c.getAuthorID(ctx)

	if !ok {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

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
	blogs, err := c.usecase.GetBlogs(ctx)

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

	authorID, ok := c.getAuthorID(ctx)
	if !ok {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

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
	authorID, ok := c.getAuthorID(ctx)

	if !ok {
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

	userID, ok := c.getAuthorID(ctx)
	if !ok {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

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

	userID, ok := c.getAuthorID(ctx)

	if !ok {
		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	comment.UserID = userID
	if err := c.usecase.AddComment(ctx, comment); err != nil {
		ctx.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.IndentedJSON(http.StatusCreated, gin.H{"message": "Comment added successfully"})
}
