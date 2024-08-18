package domain

import (
	"context"
	"time"
)

const (
	CollectionRating = "ratings"
)

type BlogRating struct {
	RatingID string `json:"rating_id" bson:"_id"`
	UserID   string `json:"user_id" bson:"user_id"`
	BlogID   string `json:"blog_id" bson:"blog_id"`
	Rating   int   `json:"rating" bson:"rating"`
	CreatedAt   time.Time `json:"createtimestamp" bson:"createtimestamp"`
	UpdatedAt    time.Time `json:"updatetimestamp" bson:"updatetimestamp"`
}

type BlogRatingRepository interface {
	InsertRating(c context.Context, rating *BlogRating) (*BlogRating, error)
	GetRatingByBlogID(c context.Context, blogID string) ([]*BlogRating, error)
	GetRatingByID(c context.Context, ratingID string) (*BlogRating, error)
	UpdateRating(c context.Context, rating int, ratingID string) (*BlogRating, error)
	DeleteRating(c context.Context, ratingID string) (*BlogRating, error)

}

type RatingUseCase interface {
	InsertRating(c context.Context, rating *BlogRating) (*BlogRating, error)
	GetRatingByBlogID(c context.Context, blogID string) ([]*BlogRating, error)
	GetRatingByID(c context.Context, ratingID string) (*BlogRating, error)
	UpdateRating(c context.Context, rating int, ratingID string) (*BlogRating, error)
	DeleteRating(c context.Context, ratingID string) (*BlogRating, error)


}