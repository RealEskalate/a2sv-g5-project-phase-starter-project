package controller

import (
	"AAiT-backend-group-8/Domain"
	"AAiT-backend-group-8/Helper"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (controller *Controller) CreateBlog(ctx *gin.Context) {
	var blog Domain.Blog

	if err := ctx.ShouldBind(&blog); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	struct_err := struct_validator.Struct(blog)

	if struct_err != nil{
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": struct_err.Error()})
		return
	}

	claims, err := Helper.Parse(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		log.Fatal("parse claims error: " + err.Error())
		return
	}

	blog.AuthorID, _ = primitive.ObjectIDFromHex((*claims)["id"].(string))
	blog.AuthorName = (*claims)["name"].(string)

	err = controller.blogUseCase.CreateBlog(&blog)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Blog created"})
	_ = controller.blogUseCase.UpdateBlogViewCount(blog.Id.Hex())
}

func (controller *Controller) GetBlogs(ctx *gin.Context) {

	pageStr := ctx.DefaultQuery("page", "1")
	pageSizeStr := ctx.DefaultQuery("pageSize", "10")
	sortBy := ctx.DefaultQuery("sortBy", "view_count")
	page, err := strconv.Atoi(pageStr)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pageSize, err1 := strconv.Atoi(pageSizeStr)

	if err1 != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
		return
	}

	blogs, err2 := controller.blogUseCase.GetAllBlogs(page, pageSize, sortBy)

	if err2 != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err2.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"blogs": blogs, "page": page, "pageSize": len(blogs)})

}

func (controller *Controller) GetBlogByID(ctx *gin.Context) {

	id := ctx.Param("id")

	blog, err := controller.blogUseCase.GetBlogByID(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"blog": blog})
	_ = controller.blogUseCase.UpdateBlogViewCount(blog.Id.Hex())
}

func (controller *Controller) UpdateBlog(ctx *gin.Context) {
	var blog Domain.Blog

	if err := ctx.ShouldBind(&blog); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claim, err := Helper.Parse(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if (*claim)["role"] != "admin" && (*claim)["id"] != blog.AuthorID.Hex() {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to perform this action"})
		return
	}

	blog.AuthorID, _ = primitive.ObjectIDFromHex((*claim)["id"].(string))
	blog.AuthorName = (*claim)["name"].(string)

	err = controller.blogUseCase.UpdateBlog(&blog)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Blog updated"})
}

func (controller *Controller) DeleteBlog(ctx *gin.Context) {
	id := ctx.Param("id")

	blog, err := controller.blogUseCase.GetBlogByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, err := Helper.Parse(ctx)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	if id, _ := primitive.ObjectIDFromHex((*claims)["id"].(string)); id != blog.AuthorID && (*claims)["role"].(string) != "admin" {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "You are not authorized to perform this action"})
		return
	}

	err = controller.blogUseCase.DeleteBlog(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = controller.commentUseCase.DeleteCommentsOfBlog(blog.Id.Hex())

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = controller.LikeUseCase.DeleteLikesOfBlog(blog.Id.Hex())

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "Blog deleted"})
}

func (controller *Controller) SearchBlog(ctx *gin.Context) {
	searchParams := Helper.GetSearchParams(ctx)
	blogs, err := controller.blogUseCase.SearchBlog(searchParams)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"blogs": blogs})
}
