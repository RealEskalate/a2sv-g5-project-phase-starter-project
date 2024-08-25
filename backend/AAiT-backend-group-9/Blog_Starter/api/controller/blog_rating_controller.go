package controller

import (
    "Blog_Starter/domain"
    "Blog_Starter/utils"
    "net/http"
    "time"
    "github.com/gin-gonic/gin"
)

type BlogRatingController struct {
    blogratingUSeCase 	domain.BlogRatingUseCase
    timeout        		time.Duration
}

func NewBlogRatingController(blogRatingUseCase domain.BlogRatingUseCase, timeout time.Duration) *BlogRatingController {
    return &BlogRatingController{
        blogratingUSeCase: blogRatingUseCase,
        timeout : timeout,
    }
}

func (bc *BlogRatingController) InsertRating(c *gin.Context) {
    blogID := c.Param("blog_id")

    user, err := utils.CheckUser(c)
    if err != nil {
        c.JSON(http.StatusBadRequest, domain.Response{
            Success: false,
            Message: err.Error(),
        })
        return
    }

    var newRating domain.BlogRatingRequest
    if err := c.BindJSON(&newRating); err != nil {
        c.JSON(http.StatusBadRequest, domain.Response{
            Success: false,
            Message: "Invalid request format",
        })
        return
    }
    newRating.UserID = user.UserID
    newRating.BlogID = blogID

    insertedRating, err := bc.blogratingUSeCase.InsertRating(c, &newRating)
    if err != nil {
        c.JSON(http.StatusInternalServerError, domain.Response{
            Success: false,
            Message: "Internal server error",
        })
        return
    }

    c.JSON(http.StatusOK, domain.Response{
        Success: true,
        Message: "Rating inserted successfully",
        Data:    insertedRating,
    })
}

func (bc *BlogRatingController) UpdateRating(c *gin.Context) {
    ratingID := c.Param("id")

    var updatedRating domain.BlogRatingRequest
    if err := c.BindJSON(&updatedRating); err != nil {
        c.JSON(http.StatusBadRequest, domain.Response{
            Success: false,
            Message: "Invalid request format",
        })
        return
    }

    existingRating, err := bc.blogratingUSeCase.GetRatingByID(c, ratingID)
    if existingRating != nil {
        updatedRating, err := bc.blogratingUSeCase.UpdateRating(c, updatedRating.Rating, ratingID)
        if err != nil {
            c.JSON(http.StatusInternalServerError, domain.Response{
                Success: false,
                Message: "Internal server error",
            })
            return
        }
        c.JSON(http.StatusOK, domain.Response{
            Success: true,
            Message: "Rating updated successfully",
            Data:    updatedRating,
        })
        return
    }
    if err != nil {
        c.JSON(http.StatusInternalServerError, domain.Response{
            Success: false,
            Message: "Internal server error",
        })
        return
    }
}

func (bc *BlogRatingController) DeleteRating(c *gin.Context) {
    ratingID := c.Param("id")
    deletedRating, err := bc.blogratingUSeCase.DeleteRating(c, ratingID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, domain.Response{
            Success: false,
            Message: "Internal server error",
        })
        return
    }

    c.JSON(http.StatusOK, domain.Response{
        Success: true,
        Message: "Rating deleted successfully",
        Data:    deletedRating,
    })
}