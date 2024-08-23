package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	dtos "github.com/aait.backend.g5.main/backend/Domain/DTOs"
	interfaces "github.com/aait.backend.g5.main/backend/Domain/Interfaces"
	models "github.com/aait.backend.g5.main/backend/Domain/Models"
	"github.com/gin-gonic/gin"
)

type BlogController struct {
	usecase interfaces.BlogUsecase
}

func NewBlogController(usecase interfaces.BlogUsecase) *BlogController {
	return &BlogController{
		usecase: usecase,
	}
}

func (c *BlogController) CreateBlogController(ctx *gin.Context) {
	var newBlog dtos.CreateBlogRequest

	if err := ctx.ShouldBind(&newBlog); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := newBlog.Validate(); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "One or more fields are missing"})
		return
	}

	authorID := ctx.GetString("id")

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

func (c *BlogController) GetBlogController(ctx *gin.Context) {
	blogId := ctx.Param("id")

	blog, err := c.usecase.GetBlog(ctx, blogId)

	if err != nil {
		ctx.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": blog})
}

func (c *BlogController) GetBlogsController(ctx *gin.Context) {
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

func (c *BlogController) SearchBlogsController(ctx *gin.Context) {

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

	blogs, err := c.usecase.SearchBlogs(ctx, filter)

	if err != nil {
		ctx.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"data": blogs})
}

func (c *BlogController) UpdateBlogController(ctx *gin.Context) {
	var updateBlog dtos.UpdateBlogRequest
	blogID := ctx.Param("id")

	if err := ctx.ShouldBind(&updateBlog); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := updateBlog.Validate(); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "One or more fields are missing"})
		return
	}

	authorID := ctx.GetString("id")

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

func (c *BlogController) DeleteBlogController(ctx *gin.Context) {
	blogID := ctx.Param("id")
	authorID := ctx.GetString("id")

	if authorID == "" {
		fmt.Println("author id", authorID)
		fmt.Println("blog id", blogID)

		ctx.IndentedJSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}

	deleteBlogReq := dtos.DeleteBlogRequest{
		BlogID:   blogID,
		AuthorID: authorID,
	}

	if err := deleteBlogReq.Validate(); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "One or more fields are missing"})
		return
	}

	if err := c.usecase.DeleteBlog(ctx, deleteBlogReq); err != nil {
		ctx.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Blog deleted successfully"})
}

func (c *BlogController) TrackPopularityController(ctx *gin.Context) {
	var blogPopularity dtos.TrackPopularityRequest

	if err := ctx.ShouldBind(&blogPopularity); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := blogPopularity.Validate(); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "One or more fields are missing"})
		return
	}
	
	userID := ctx.GetString("id")
	blogID := ctx.Param("id")

	blogPopularity.UserID = userID
	blogPopularity.BlogID = blogID

	if err := c.usecase.TrackPopularity(ctx, blogPopularity); err != nil {
		ctx.IndentedJSON(err.Code, gin.H{"error": err.Message})
		return
	}

	ctx.IndentedJSON(http.StatusOK, gin.H{"message": "Popularity tracked successfully"})
}
