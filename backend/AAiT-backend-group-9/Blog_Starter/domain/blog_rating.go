package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	CollectionRating = "ratings"
)

type BlogRating struct {
	RatingID  primitive.ObjectID `json:"rating_id" bson:"_id"`
	UserID    string             `json:"user_id" bson:"user_id"`
	BlogID    string             `json:"blog_id" bson:"blog_id"`
	Rating    int                `json:"rating" bson:"rating"`
	CreatedAt time.Time          `json:"createtimestamp" bson:"createtimestamp"`
	UpdatedAt time.Time          `json:"updatetimestamp" bson:"updatetimestamp"`
}

type BlogRatingRequest struct {
	RatingID string `json:"rating_id"`
	UserID   string `json:"user_id"`
	BlogID   string `json:"blog_id"`
	Rating   int    `json:"rating"`
}

type BlogRatingRepository interface {
	InsertRating(c context.Context, rating *BlogRating) (*BlogRating, error)
	GetRatingByBlogID(c context.Context, blogID string) ([]*BlogRating, error)
	GetRatingByID(c context.Context, ratingID string) (*BlogRating, error)
	UpdateRating(c context.Context, rating int, ratingID string) (*BlogRating, int, error)
	DeleteRating(c context.Context, ratingID string) (*BlogRating, error)
	DeleteRatingByBlogID(c context.Context, blogID string) error
}

type BlogRatingUseCase interface {
	InsertRating(c context.Context, rating *BlogRatingRequest) (*BlogRating, error)
	GetRatingByBlogID(c context.Context, blogID string) ([]*BlogRating, error)
	GetRatingByID(c context.Context, ratingID string) (*BlogRating, error)
	UpdateRating(c context.Context, rating int, ratingID string) (*BlogRating, error)
	DeleteRating(c context.Context, ratingID string) (*BlogRating, error)
}
