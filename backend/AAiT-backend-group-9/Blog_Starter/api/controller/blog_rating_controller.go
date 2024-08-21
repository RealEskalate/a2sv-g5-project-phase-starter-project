package controller

import (
	"Blog_Starter/domain"
	"time"
	"net/http"
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

func (bc *BlogRatingController) InserttAndUpdateRating(c *gin.Context) {
	var newRating domain.BlogRatingRequest
	blogID := c.Param("blog_id")
	if err := c.BindJSON(&newRating); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : "invalid request format"})
	}


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
	var toDelete domain.BlogRatingRequest
	ratingID := c.Param("rating_id")
	toDelete.RatingID = ratingID
	deletedRating, err := bc.blogratingUSeCase.DeleteRating(c, toDelete.RatingID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error" : "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"'deleted_rating" : deletedRating})
}