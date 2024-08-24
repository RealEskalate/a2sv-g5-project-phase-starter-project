package controller

import (
	"AAiT-backend-group-8/Domain"
	infrastructure "AAiT-backend-group-8/Infrastructure"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (controller *Controller) CreateBlog(ctx *gin.Context) {
	var blog Domain.Blog

	if err := ctx.ShouldBind(&blog); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// fmt.Println(blog)

	struct_err := struct_validator.Struct(blog)

	if struct_err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": struct_err.Error()})
		return
	}

	claims, err := infrastructure.Parse(ctx)

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
	key := infrastructure.GetCacheKey(ctx)

	jsonList, err := controller.rdb.LRange(ctx, key, 0, -1).Result()

	var blogs []Domain.Blog

	if err == nil && len(jsonList) != 0 {
		for _, j := range jsonList {
			var blog Domain.Blog
			err := json.Unmarshal([]byte(j), &blog)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			blogs = append(blogs, blog)
		}
		ctx.JSON(http.StatusOK, gin.H{"blogs": blogs})
		return
	}

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

	blogs, err = controller.blogUseCase.GetAllBlogs(page, pageSize, sortBy)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"blogs": blogs})

	for _, blog := range blogs {
		err := controller.cacheUseCase.Update(blog.Id, key)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		jsonBlog, err := json.Marshal(blog)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Push to Redis list and handle any errors
		err = controller.rdb.RPush(ctx, key, jsonBlog).Err()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	controller.rdb.Expire(ctx, key, 10*time.Minute)

}

func (controller *Controller) GetBlogByID(ctx *gin.Context) {

	id := ctx.Param("id")

	var blog = &Domain.Blog{}

	jsonData, err := controller.rdb.Get(context.Background(), id).Result()

	if err == nil {
		var blog Domain.Blog
		err = json.Unmarshal([]byte(jsonData), &blog)

		if err == nil {
			fmt.Println("from here")
			_ = controller.blogUseCase.UpdateBlogViewCount(id)

			blog.ViewCount += 1
			ctx.JSON(http.StatusOK, gin.H{"blog": blog})

			_ = controller.blogUseCase.UpdateBlogViewCount(id)
			return

		}

	}

	blog, err = controller.blogUseCase.GetBlogByID(id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = controller.blogUseCase.UpdateBlogViewCount(blog.Id.Hex())
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"blog": blog})
	jsonBlog, err := json.Marshal(blog)
	if err != nil {
		return
	}

	controller.rdb.Set(ctx, id, jsonBlog, 0)
}

func (controller *Controller) UpdateBlog(ctx *gin.Context) {
	var blog Domain.Blog

	if err := ctx.ShouldBind(&blog); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claim, err := infrastructure.Parse(ctx)

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

	keys, err := controller.cacheUseCase.Delete(blog.Id)
	if err != nil {
		return
	}

	err = controller.rdb.Get(ctx, blog.Id.Hex()).Err()

	if err == nil {
		blogJson, err := json.Marshal(blog)
		if err != nil {
			return
		}
		controller.rdb.Set(context.Background(), blog.Id.Hex(), blogJson, 10*time.Minute)
	}

	for _, key := range keys {
		controller.rdb.Del(ctx, key)
	}

}

func (controller *Controller) DeleteBlog(ctx *gin.Context) {
	id := ctx.Param("id")

	blog, err := controller.blogUseCase.GetBlogByID(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	claims, err := infrastructure.Parse(ctx)

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

	_ = controller.commentUseCase.DeleteCommentsOfBlog(blog.Id.Hex())

	_ = controller.LikeUseCase.DeleteLikesOfBlog(blog.Id.Hex())

	ctx.JSON(http.StatusOK, gin.H{"message": "Blog deleted"})

	//delete if the blog is cached
	controller.rdb.Del(context.Background(), blog.Id.Hex())

	//find all keys the blog belongs to
	keys, err := controller.cacheUseCase.Delete(blog.Id)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	//iterate over the keys and delete them from the cache
	for _, key := range keys {
		controller.rdb.Del(context.Background(), key)
	}

}

func (controller *Controller) SearchBlog(ctx *gin.Context) {
	searchParams := infrastructure.GetSearchParams(ctx)
	key := infrastructure.GetCacheKey(ctx)

	jsonList, err := controller.rdb.LRange(context.Background(), key, 0, -1).Result()
	var blogs []Domain.Blog

	if err == nil && len(jsonList) != 0 {
		fmt.Println("here in the search")

		var blog Domain.Blog
		for _, j := range jsonList {
			err = json.Unmarshal([]byte(j), &blog)
			if err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			blogs = append(blogs, blog)
		}
		ctx.JSON(http.StatusOK, gin.H{"blogs": blogs})
		return
	}

	blogs, err = controller.blogUseCase.SearchBlog(searchParams)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"blogs": blogs})
	for _, blog := range blogs {
		err := controller.cacheUseCase.Update(blog.Id, key)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		jsonBlog, err := json.Marshal(blog)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		// Push to Redis list and handle any errors
		err = controller.rdb.RPush(ctx, key, jsonBlog).Err()
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
	}
	controller.rdb.Expire(ctx, key, 10*time.Minute)
}
