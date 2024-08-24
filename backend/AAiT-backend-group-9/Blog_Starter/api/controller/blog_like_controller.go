package controller

import (
    "Blog_Starter/domain"
    "Blog_Starter/utils"
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
)

type LikeController struct {
    LikeUsecase domain.LikeUseCase
    timeout     time.Duration
}

func NewLikeController(blogLikeUseCase domain.LikeUseCase, timeout time.Duration) *LikeController {
    return &LikeController{
        LikeUsecase: blogLikeUseCase,
        timeout:     timeout,
    }
}

func (lc *LikeController) LikeBlog(c *gin.Context) {
    var likeRequest domain.Like
    user, err := utils.CheckUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, domain.Response{
            Success: false,
            Message: "Unauthorized user",
        })
        return
    }

    blogID := c.Param("blog_id")
    likeRequest.UserID = user.UserID
    likeRequest.BlogID = blogID
    likeResponse, err := lc.LikeUsecase.LikeBlog(c, &likeRequest)
    if err != nil {
        c.JSON(http.StatusInternalServerError, domain.Response{
            Success: false,
            Message: "Cannot like blog",
        })
        return
    }
    c.JSON(http.StatusOK, domain.Response{
        Success: true,
        Message: "Blog liked successfully",
        Data:    likeResponse,
    })
}

func (lc *LikeController) UnlikeBlog(c *gin.Context) {
    likeID := c.Param("id")
    likeResponse, err := lc.LikeUsecase.UnlikeBlog(c, likeID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, domain.Response{
            Success: false,
            Message: "Cannot unlike blog",
        })
        return
    }
    c.JSON(http.StatusOK, domain.Response{
        Success: true,
        Message: "Blog unliked successfully",
        Data:    likeResponse,
    })
}

func (lc *LikeController) GetByID(c *gin.Context) {
    blogID := c.Param("blog_id")
    user, err := utils.CheckUser(c)
    if err != nil {
        c.JSON(http.StatusUnauthorized, domain.Response{
            Success: false,
            Message: "Unauthorized user",
        })
        return
    }
    likeResponse, err := lc.LikeUsecase.GetByID(c, user.UserID, blogID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, domain.Response{
            Success: false,
            Message: "Cannot get like",
        })
        return
    }
    c.JSON(http.StatusOK, domain.Response{
        Success: true,
        Message: "Like retrieved successfully",
        Data:    likeResponse,
    })
}