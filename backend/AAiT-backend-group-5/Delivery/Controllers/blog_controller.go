package controllers

import (
	"log"
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

<<<<<<< HEAD
func (c *blogController) getAuthorID(ctx *gin.Context) string {
	return ctx.GetString("id")
=======
func (blogController BlogController) GetBlog(ctx *gin.Context) {
	blogID := ctx.Param("id")

	blog, err := blogController.BlogUsecase.GetBlog(ctx, blogID)
	if err != nil {
		ctx.JSON(err.Code, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, blog)
>>>>>>> f6022233 (finish blog controller)
}

func (c *blogController) CreateBlogController(ctx *gin.Context) {
	var newBlog dtos.CreateBlogRequest

	if err := ctx.ShouldBind(&newBlog); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	authorID := c.getAuthorID(ctx)

	log.Println(authorID, "This is the author ID")

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

<<<<<<< HEAD
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
=======
func (blogController BlogController) UpdateBlog(ctx *gin.Context) {
	var updated_request dtos.UpdateBlogRequest

	if e := ctx.BindJSON(updated_request); e != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// get user ID from the context
	userId, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid token"})
		return
	}

	// check if the user matches the authorID of the blog
	if userId != updated_request.AuthorID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		return
	}

	// set fields for the updated blog
	blogID, e := primitive.ObjectIDFromHex(updated_request.BlogID)
	if e != nil {
		ctx.JSON(http.StatusInternalServerError, errors.New("internal server error"))
		return
	}

	updated_blog := models.Blog{
		ID:      blogID,
		Title:   updated_request.Title,
		Content: updated_request.Content,
		Tags:    updated_request.Tags,
	}

	err := blogController.BlogUsecase.UpdateBlog(ctx, updated_request.BlogID, &updated_blog)
	if err != nil {
		ctx.JSON(err.Code, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "blog update successful"})
}

func (blogController BlogController) DeleteBlog(ctx *gin.Context) {
	var delete_request dtos.DeleteBlogRequest
>>>>>>> f6022233 (finish blog controller)

	if err := ctx.ShouldBind(&filter); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

<<<<<<< HEAD
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
=======
	// get user ID and role from the context
	userId, exists := ctx.Get("id")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid token"})
		return
	}

	userRole, exists := ctx.Get("role")
	if !exists {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid token"})
		return
	}

	// check if the user matches the authorID of the blog or if user is an admin
	if userId != delete_request.AuthorID && userRole != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		return
	}

	err := blogController.BlogUsecase.DeleteBlog(ctx, delete_request.BlogID)
	if err != nil {
		ctx.JSON(err.Code, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "blog delete successful"})
}

func (blogController BlogController) HandelTrackPopularity(ctx *gin.Context) {
	var request dtos.TrackPopularityRequest

	// attempt to bind payload to the request
	err := ctx.ShouldBind(request)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	// get userId from context
	userId := ctx.GetString("id")
	if userId == "" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		return
	}

	// append userId to the UserID field of the request
	request.UserID = userId

	// invoke TrackPopularity Usecase
	e := blogController.BlogUsecase.TrackPopularity(ctx, request.BlogID, request)
	if e != nil {
		ctx.JSON(e.Code, e.Error())
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "successful"})
>>>>>>> f6022233 (finish blog controller)
}
