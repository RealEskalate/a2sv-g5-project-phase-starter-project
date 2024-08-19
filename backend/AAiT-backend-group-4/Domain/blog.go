package domain

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Tag string
type FilterParam string

const (
	Date       FilterParam = "date"
	Popularity FilterParam = "popularity"
)

const (
	Tech            Tag = "Tech"
	Education       Tag = "Education"
	HealthWellness  Tag = "Health & Wellness"
	Lifestyle       Tag = "Lifestyle"
	FinanceBusiness Tag = "Finance & Business"
)

type Author struct {
	Author_ID string
	Name      string
}

type Comment struct {
	User_ID   string    `json:"user_id"`
	User_Name string    `json:"user_name"`
	Content   string    `json:"content"`
	Date      time.Time `json:"date"`
}

type Feedback struct {
	View_count int       `json:"view_count"`
	Dislikes   int       `json:"dislikes"`
	Likes      int       `json:"likes"`
	Comments   []Comment `json:"comments"`
}

type Blog struct {
	ID          primitive.ObjectID `bson:"_id"`
	Title       string             `json:"title" bson:"title" validate:"required,min=5,max=100"`
	Content     string             `json:"content" bson:"content" validate:"required,min=5"`
	Author_Info Author             `json:"author_info" bson:"author_info"`
	Tags        []Tag              `json:"tags" bson:"tags"`
	Popularity  float64            `json:"popularity" bson:"popularity"`
	Feedbacks   Feedback           `json:"feedbacks" bson:"feedbacks"`
	Created_At  time.Time          `json:"created_at" bson:"created_at"`
	Updated_At  time.Time          `json:"updated_at" bson:"updated_at"`
}

type Filter struct {
	AuthorName *string
	Tags       *[]Tag
	BlogTitle  *string
	Popularity *float64
	Sort_By    *FilterParam
}

type BlogUpdate struct {
	Title       *string    `json:"title" bson:"title" validate:"required,min=5,max=100"`
	Content     *string    `json:"content" bson:"content" validate:"required,min=5"`
	Author_Info *Author    `json:"author_info" bson:"author_info"`
	Tags        *[]Tag     `json:"tags" bson:"tags"`
	Popularity  *float64   `json:"popularity" bson:"popularity"`
	Feedbacks   *Feedback  `json:"feedbacks" bson:"feedbacks"`
	Updated_At  *time.Time `json:"updated_at" bson:"updated_at"`
}


type BlogRepository interface {
	// CreateBlog inserts a new blog into the collection
	CreateBlog(c context.Context, blog *Blog) error
	// FetchByBlogID retrieves a blog by its ID
	FetchByBlogID(c context.Context, blogID string) (Blog, error)
	// FetchAll retrieves all blogs from the collection
	FetchAll(c context.Context) ([]Blog, error)
	// FetchByBlogAuthor retrieves blogs by the author's ID
	FetchByBlogAuthor(c context.Context, authorID string) ([]Blog, error)
	// FetchByBlogTitle retrieves blogs by their title
	FetchByBlogTitle(c context.Context, title string) ([]Blog, error)
	// UpdateBlog updates a blog in the collection by its ID
	UpdateBlog(ctx context.Context, id primitive.ObjectID, blog BlogUpdate) error
	// DeleteBlog deletes a blog from the collection by its ID
	DeleteBlog(ctx context.Context, id primitive.ObjectID) error
	// BlogExists checks if a blog exists by its ID
	BlogExists(ctx context.Context, id primitive.ObjectID) (bool, error)
	// UserIsAuthor checks if a user is the author of a blog by their user ID and the blog ID
	UserIsAuthor(ctx context.Context, blogID primitive.ObjectID, userID string) (bool, error)
	// UpdateFeedback retrives blogs by it's id and updates it's feedback
	UpdateFeedback(ctx context.Context, id string, updateFunc func(*Feedback) error) error

}