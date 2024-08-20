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

type PaginationData struct {
	NextPage     int `json:"next_page"`
	PreviousPage int `json:"previous_page"`
	CurrentPage  int `json:"current_page"`
	TotalPages   int `json:"total_pages"`
	TotalItems   int `json:"total_items"`
}

type PaginatedBlogs struct {
	Blogs      []Blog         `json:"blogs"`
	Pagination PaginationData `json:"pagination"`
}

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

	// SearchBlogs filters Blogs with given filter parameters and retrives blogs form the dataBase
	SearchBlogs(c context.Context, filter Filter, limit, offset int) ([]Blog, int, error)
	// CreateBlog inserts a new blog into the collection
	CreateBlog(c context.Context, blog *Blog) error
	// FetchByBlogID retrieves a blog by its ID
	FetchByBlogID(c context.Context, blogID string) (Blog, error)
	// FetchAll retrieves all blogs from the collection
	FetchAll(c context.Context, limit, offset int) ([]Blog, int, error)
	// FetchByBlogAuthor retrieves blogs by the author's ID
	FetchByBlogAuthor(c context.Context, authorID string, limit, offset int) ([]Blog, int, error)
	// FetchByBlogTitle retrieves blogs by their title
	FetchByBlogTitle(c context.Context, title string, limit, offset int) ([]Blog, int, error)

	FetchByPageAndPopularity(ctx context.Context, limit, offset int) ([]Blog, int, error)
	FetchByTags(ctx context.Context, tags []Tag, limit, offset int) ([]Blog, int, error)
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

	// AddLikes adds the number of Likes in the feedback field found in the blog info
	IncrementLikes(feedback *Feedback) error
	// DecrementsLikes decrements the number of likes from the feedback field of the blog
	DecrementLikes(feedback *Feedback) error
	// IncrementDislikes Increments the number of Dislikes from the feedback field of the blog
	IncrementDislike(feedback *Feedback) error
	// DecrementDislikes decrements the number of Dislikes from the feedback field of the blog
	DecrementDislikes(feedback *Feedback) error
	// AddComent adds a comment in to a blog feedback section
	AddComment(feedback *Feedback, comment Comment) error
	// UpdateComment updates a comment using the comment's user id and blogs id
	UpdateComment(feedback *Feedback, updatedComment Comment, userID string) error
	// RemoveComents removes a comment in a blog feedback section using the removing users id
	RemoveComment(feedback *Feedback, requesterUserID string, isAdmin bool) error
}

type BlogUsecase interface {
	// CreateBlog inserts a new blog into the collection
	CreateBlog(c context.Context, blog *Blog) error
	// FetchByBlogID retrieves a blog by its ID
	FetchByBlogID(c context.Context, blogID string) (Blog, error)
	// FetchAll retrieves all blogs from the collection
	FetchAll(c context.Context, limit, page int) (PaginatedBlogs, error)
	// FetchByBlogAuthor retrieves blogs by the author's ID
	FetchByBlogAuthor(c context.Context, authorID string, limit, page int) (PaginatedBlogs, error)
	// FetchByBlogTitle retrieves blogs by their title
	FetchByBlogTitle(c context.Context, title string, limit, page int) (PaginatedBlogs, error)

	FetchByPageAndPopularity(ctx context.Context, limit, page int) (PaginatedBlogs, error)

	FetchByTags(ctx context.Context, tags []Tag, limit, page int) (PaginatedBlogs, error)

	// UpdateBlog updates a blog in the collection by its ID
	UpdateBlog(ctx context.Context, id primitive.ObjectID, blog BlogUpdate, updatingID string) error
	// DeleteBlog deletes a blog from the collection by its ID
	DeleteBlog(ctx context.Context, id primitive.ObjectID, deletingID string) error

	SearchBlogs(c context.Context, filter Filter, limit, page int) (PaginatedBlogs, error)
	// AddComent adds a comment in to a blog feedback section
	AddComment(ctx context.Context, blogID string, comment Comment) error
	// UpdateComment updates a comment using the comment's user id
	UpdateComment(c context.Context, blogID primitive.ObjectID, userID string, updatedComment Comment) error
	// RemoveComents removes a comment in a blog feedback section using the removing users id
	RemoveComment(ctx context.Context, blogID primitive.ObjectID, requesterUserID string) error
}
