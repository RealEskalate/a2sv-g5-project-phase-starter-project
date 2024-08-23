package Domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Blog struct {
	Id           primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Title        string             `bson:"title" json:"title" validate:"required,min=3,max=100"`
	Body         string             `bson:"body" json:"body" validate:"required"`
	Tags         []string           `bson:"tags" json:"tags"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	LastUpdated  time.Time          `bson:"last_updated" json:"last_updated"`
	AuthorName   string             `bson:"author_name" json:"author_name" validate:"max=100"`
	AuthorID     primitive.ObjectID `bson:"author_id" json:"author_id"`
	ViewCount    int                `bson:"view_count" json:"view_count"`
	LikeCount    int                `bson:"like_count" json:"like_count"`
	CommentCount int                `bson:"comment_count" json:"comment_count"`
}

type SearchCriteria struct {
	Title     string    `form:"title" bson:"title" json:"title"`
	Author    string    `form:"author" bson:"author" json:"author"`
	Tags      []string  `form:"tags" bson:"tags" json:"tags"`
	StartDate time.Time `form:"startDate" bson:"start_date" json:"start_date"`
	EndDate   time.Time `form:"endDate" bson:"end_date" json:"end_date"`
	MinViews  int       `form:"minViews" bson:"min_views" json:"min_views"`
	SortBy    string    `form:"sortBy" bson:"sort_by" json:"sort_by"`
	Page      int       `form:"page" bson:"page" json:"page"`
	PageSize  int       `form:"pageSize" bson:"page_size" json:"page_size"`
}
