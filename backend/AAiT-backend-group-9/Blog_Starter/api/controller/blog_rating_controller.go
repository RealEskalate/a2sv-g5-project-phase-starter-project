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

func (bc *BlogRatingController) InsertAndUpdateRating(c *gin.Context) {
	var newRating domain.BlogRatingRequest
	blogID := c.Param("id")
	if err := c.BindJSON(&newRating); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : "invalid request format"})
	}

	user, err := utils.CheckUser(c)
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}
	
	newRating.UserID = user.UserID
	newRating.BlogID = blogID
	if newRating.RatingID != "" {
		exisitingRating, err := bc.blogratingUSeCase.GetRatingByID(c, newRating.RatingID)
		if exisitingRating != nil {
			updatedRating, err := bc.blogratingUSeCase.UpdateRating(c, newRating.Rating, newRating.RatingID)
			if err != nil {
				c.IndentedJSON(http.StatusInternalServerError, gin.H{"error" : "internal server error"})
				return
			}
			c.IndentedJSON(http.StatusOK, gin.H{"updated_rating" : updatedRating})
			return
		}
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"error" : "internal server error"})
			return
		}
	}

	insertedRating, err := bc.blogratingUSeCase.InsertRating(c, &newRating)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error" : "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"inserted_rating" : insertedRating})
}

func (bc *BlogRatingController) DeleteRating(c *gin.Context) {

	ratingID := c.Param("id")
	deletedRating, err := bc.blogratingUSeCase.DeleteRating(c, ratingID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error" : "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"'deleted_rating" : deletedRating})
}