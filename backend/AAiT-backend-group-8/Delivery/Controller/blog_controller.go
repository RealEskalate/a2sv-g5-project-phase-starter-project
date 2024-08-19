package controllers

import (
	"AAiT-backend-group-8/Domain"
	"AAiT-backend-group-8/Helper"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (controller *Controller) CreateBlog(ctx *gin.Context) {
	var blog Domain.Blog

	if err := ctx.ShouldBind(&blog); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := controller.blogUseCase.CreateBlog(&blog)

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
	sortBy := ctx.DefaultQuery("sortBy", "ViewCount")
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

	ctx.JSON(http.StatusOK, gin.H{"blogs": *blogs, "page": page, "pageSize": len(*blogs)})

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

	err := Helper.Authenticate(ctx, false, &blog)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
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

	err = Helper.Authenticate(ctx, true, blog)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
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

	ctx.JSON(http.StatusOK, gin.H{"message": "Blog deleted"})
}

func (controller *Controller) SearchBlog(ctx *gin.Context) {
	searchParams := Helper.GetSearchParams(ctx)
	blogs, err := controller.blogUseCase.SearchBlog(searchParams)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"blogs": *blogs})
}
