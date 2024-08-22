package controllers

import (
	"log"
	"net/http"
	"strconv"
	"strings"

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

	title := ctx.DefaultQuery("title", "")
	authorName := ctx.DefaultQuery("author_name", "")
	date := ctx.DefaultQuery("date", "")
	viewCount := ctx.DefaultQuery("view_count", "-1")
	likeCount := ctx.DefaultQuery("like_count", "-1")
	dislikeCount := ctx.DefaultQuery("dislike_count", "-1")
	tags := ctx.DefaultQuery("tags", "")

	viewCountInt, _ := strconv.Atoi(viewCount)
	likeCountInt, _ := strconv.Atoi(likeCount)
	dislikeCountInt, _ := strconv.Atoi(dislikeCount)

	tagsSlice := []string{}
	if tags != "" {
		tagsSlice = strings.Split(tags, ",")
	}

	filter := dtos.FilterBlogRequest{
		Title:        title,
		AuthorName:   authorName,
		Date:         date,
		ViewCount:    viewCountInt,
		LikeCount:    likeCountInt,
		DislikeCount: dislikeCountInt,
		Tags:         tagsSlice,
	}

	log.Println(title, "filter")
	log.Println(tags, "filter")

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

	log.Println(updateBlog, "updateBlog")

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
	blogID := ctx.Param("id")

	blogPopularity.UserID = userID
	blogPopularity.BlogID = blogID

	if err := c.usecase.TrackPopularity(ctx, blogPopularity); err != nil {
		ctx.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Popularity tracked successfully"})
}
