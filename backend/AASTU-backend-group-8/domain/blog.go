package domain

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogPost struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Title        string             `bson:"title"`
	Content      string             `bson:"content"`
	AuthorID     primitive.ObjectID `bson:"author_id"`
	Tags         []string           `bson:"tags"`
	CreatedAt    time.Time          `bson:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at"`
	Veiws        int                `json:"views"`
	LikeCount    int                `bson:"like_count"`
	DislikeCount int                `bson:"dislike_count"`
}

type SearchBlogPost struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

type BlogRepositoryInterface interface {
	Save(blog *BlogPost) (string, error)
	GetAllBlog(pagination Pagination, sortBy string, sortOrder int, filter BlogFilter) ([]BlogPost, error)
	GetBlogByID(id primitive.ObjectID) (*BlogPost, error)
	Update(id primitive.ObjectID, blog *BlogPost) (*BlogPost, error)
	Delete(id primitive.ObjectID) error
	ToggleLikeDislike(blogID, userID primitive.ObjectID, like bool) error

	HasUserLiked(blogID, userID primitive.ObjectID) (bool, error)
	HasUserDisliked(blogID, userID primitive.ObjectID) (bool, error)
	UpdateLikeDislikeCount(blogID, userID primitive.ObjectID, like bool) error
	AddComment(blogID primitive.ObjectID, comment Comment) error
	GetBlogPostByID(blogID primitive.ObjectID) (BlogPost, error)
}
type BlogUsecaseInterface interface {
	CreateBlogPost(blog *BlogPost) (string, error)
	GetAllBlogPosts(Pagination, string, int, BlogFilter) ([]BlogPost, error)
	GetBlogByID(id primitive.ObjectID) (*BlogPost, error)
	UpdateBlogPost(id primitive.ObjectID, blog *BlogPost) (*BlogPost, error)
	DeleteBlogPost(id primitive.ObjectID) error

	LikeBlogPost(blogID primitive.ObjectID, userID primitive.ObjectID) error
	DislikeBlogPost(blogID primitive.ObjectID, userID primitive.ObjectID) error
	AddCommentToBlogPost(blogID primitive.ObjectID, comment Comment) error
}
