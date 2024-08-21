package controller

import (
	"Blog_Starter/domain"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type LikeController struct {
	LikeUsecase domain.LikeUseCase

}

func(lc *LikeController) LikeBlog(c *gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	var likeRequest domain.Like
	err:= c.BindJSON(&likeRequest)
	if err!=nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
		return
	}
	likeResponse,err := lc.LikeUsecase.LikeBlog(ctx, &likeRequest)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot like blog"})
		return
	}
	c.JSON(http.StatusOK, likeResponse)
}

func(lc *LikeController) UnlikeBlog(c *gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	likeID := c.Param("like_id")  // TODO find better ways to get the likeID
	likeResponse,err := lc.LikeUsecase.UnlikeBlog(ctx, likeID)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot unlike blog"})
		return
	}
	c.JSON(http.StatusOK, likeResponse)
}

func (lc *LikeController) GetByID(c *gin.Context){
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	userID := c.Param("user_id")
	blogID := c.Param("blog_id")
	likeResponse,err := lc.LikeUsecase.GetByID(ctx, userID, blogID)
	if err!=nil{
		c.JSON(http.StatusInternalServerError, gin.H{"error": "cannot get like"})
		return
	}
	c.JSON(http.StatusOK, likeResponse)
}