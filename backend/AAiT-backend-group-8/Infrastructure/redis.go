package infrastructure

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func InitRedis() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("redis init success")
	return rdb

}

func GetCacheKey(ctx *gin.Context) string {
	Title := ctx.Query("title")
	Author := ctx.Query("author")
	Tags := ctx.QueryArray("tags")
	StartDateStr := ctx.Query("startDate")
	EndDateStr := ctx.Query("endDate")
	MinViewsStr := ctx.Query("minViews")
	SortBy := ctx.DefaultQuery("sortBy", "view_count")
	PageStr := ctx.DefaultQuery("page", "10")
	PageSizeStr := ctx.DefaultQuery("pageSize", "1")

	res := Title + ":" + Author + ":" + StartDateStr + ":" + EndDateStr + ":" + MinViewsStr + ":" + SortBy + ":" + PageStr + ":" + PageSizeStr

	for _, v := range Tags {
		res += v
	}

	return res
}
