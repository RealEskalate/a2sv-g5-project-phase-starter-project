package controller

import (
	"Blog_Starter/domain"
	"context"
	"net/http"
	"github.com/gin-gonic/gin"
)

type BlogRatingController struct {
	blogratingUSeCase 	domain.BlogRatingUseCase
	ctx          		context.Context
}

func NewBlogRatingController(blogRatingUseCase domain.BlogRatingUseCase, ctx context.Context) *BlogRatingController {
	return &BlogRatingController{
		blogratingUSeCase: blogRatingUseCase,
		ctx:           ctx,
	}
}

func (bc *BlogRatingController) InserttAndUpdateRating(c *gin.Context) {
	var newRating domain.BlogRatingRequest
	if err := c.BindJSON(&newRating); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : "invalid request format"})
	}
	
	if newRating.RatingID != "" {
		exisitingRating, err := bc.blogratingUSeCase.GetRatingByID(bc.ctx, newRating.RatingID)
		if exisitingRating != nil {
			updatedRating, err := bc.blogratingUSeCase.UpdateRating(bc.ctx, newRating.Rating, newRating.RatingID)
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

	insertedRating, err := bc.blogratingUSeCase.InsertRating(bc.ctx, &newRating)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error" : "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"inserted_rating" : insertedRating})
}

func (bc *BlogRatingController) DeleteRating(c *gin.Context) {
	var toDelete domain.BlogRatingRequest
	if err := c.BindJSON(&toDelete); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error" : "invalid request format"})
		return
	}

	deletedRating, err := bc.blogratingUSeCase.DeleteRating(bc.ctx, toDelete.RatingID)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"error" : "internal server error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"'deleted_rating" : deletedRating})
}