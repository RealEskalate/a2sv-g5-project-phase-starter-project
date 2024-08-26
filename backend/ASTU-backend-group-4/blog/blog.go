package blog

import (
	"context"
	"time"

	"github.com/RealEskalate/-g5-project-phase-starter-project/astu/backend/g4/pkg/infrastructure"
)

var weightsToUse = map[string]float32{
	"views":    0.5,
	"comments": 0.3,
	"likes":    0.2,
	"dislikes": 0.1,
}

type Blog struct {
	ID            string    `json:"id,omitempty" bson:"_id,omitempty"`
	AuthorID      string    `json:"author_id" bson:"author_id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	Tags          []string  `json:"tags"`
	ViewsCount    int       `json:"views_count" bson:"views_count"`
	CommentsCount int       `json:"comments_count" bson:"comments_count"`
	LikesCount    int       `json:"likes_count" bson:"likes_count"`
	DislikesCount int       `json:"dislikes_count" bson:"dislikes_count"`
	Popularity    float32   `json:"popularity"`
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" bson:"updated_at"`
}

func (b *Blog) CalculatePopularity() {
	b.Popularity = float32(
		float32(b.ViewsCount)*weightsToUse["views"] +
			// float32(b.CommentsCount)*weightsToUse["comments"] +
			float32(b.LikesCount)*weightsToUse["likes"] -
			float32(b.DislikesCount)*weightsToUse["dislikes"],
	)
}

func (b *Blog) IncrementViewsCount() {
	b.ViewsCount++
	b.CalculatePopularity()
}

func (b *Blog) IncrementCommentsCount() {
	b.CommentsCount++
	b.CalculatePopularity()
}

func (b *Blog) IncrementLikesCount() {
	b.LikesCount++
	b.CalculatePopularity()
}

func (b *Blog) IncrementDislikesCount() {
	b.DislikesCount++
	b.CalculatePopularity()
}

func (b *Blog) DecrementLikesCount() {
	b.LikesCount--
	b.CalculatePopularity()
}

func (b *Blog) DecrementDislikesCount() {
	b.DislikesCount--
	b.CalculatePopularity()
}

func (b *Blog) DecrementCommentsCount() {
	b.CommentsCount--
	b.CalculatePopularity()
}

type BlogSummary struct {
	ID            string    `json:"id" bson:"_id,omitempty"`
	AuthorID      string    `json:"author_id" bson:"author_id"`
	Title         string    `json:"title"`
	Tags          []string  `json:"tags"`
	ViewsCount    int       `json:"views_count" bson:"views_count"`
	CommentsCount int       `json:"comments_count" bson:"comments_count"`
	LikesCount    int       `json:"likes_count" bson:"likes_count"`
	DislikesCount int       `json:"dislikes_count" bson:"dislikes_count"`
	Popularity    float32   `json:"popularity"`
	CreatedAt     time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" bson:"updated_at"`
}

type Comment struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	BlogID    string    `json:"blog_id" bson:"blog_id"`
	AuthorID  string    `json:"author_id" bson:"author_id"`
	Content   string    `json:"content" bson:"content"`
	CreatedAt time.Time `json:"created_at" bson:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at"`
}

type Like struct {
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	BlogID string `json:"blog_id,omitempty" bson:"blog_id"`
	UserID string `json:"user_id,omitempty" bson:"user_id"`
}

type Dislike struct {
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	BlogID string `json:"blog_id,omitempty" bson:"blog_id"`
	UserID string `json:"user_id,omitempty" bson:"user_id"`
}

type FilterQuery struct {
	Tags           []string
	CreatedAtFrom  string
	CreatedAtTo    string
	PopularityFrom float32
	PopularityTo   float32
}

type BlogUseCase interface {
	CreateBlog(ctx context.Context, authorID string, blog CreateBlogRequest) (Blog, error)
	UpdateBlog(ctx context.Context, id, userID string, blog UpdateBlogRequest) (Blog, error)
	DeleteBlog(ctx context.Context, id, userID string) error
	GetBlogByID(ctx context.Context, id string) (Blog, error)
	GetBlogs(ctx context.Context, filterQuery FilterQuery, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[BlogSummary], error)
	SearchBlogs(ctx context.Context, query string, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[BlogSummary], error)
	GetCommentsByBlogID(ctx context.Context, blogID string, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[Comment], error)
	CreateComment(ctx context.Context, userID, blogId string, comment CreateCommentRequest) error
	DeleteComment(ctx context.Context, id, userID string) error
	LikeBlog(ctx context.Context, userID string, blogID string) error
	UnLikeBlog(ctx context.Context, userID string, blogID string) error
	DislikeBlog(ctx context.Context, userID string, blogID string) error
	UnDislikeBlog(ctx context.Context, userID string, blogID string) error
}

type BlogRepository interface {
	CreateBlog(ctx context.Context, blog Blog) (string, error)
	UpdateBlog(ctx context.Context, id string, blog Blog) error
	DeleteBlog(ctx context.Context, id string) error
	GetBlogByID(ctx context.Context, id string) (Blog, error)
	GetBlogs(ctx context.Context, filterQuery FilterQuery, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[BlogSummary], error)
	SearchBlogs(ctx context.Context, query string, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[BlogSummary], error)
	GetCommentsByBlogID(ctx context.Context, blogID string, pagination infrastructure.PaginationRequest) (infrastructure.PaginationResponse[Comment], error)
	DeleteCommentsByBlogID(ctx context.Context, blogID string) error
	CreateComment(ctx context.Context, comment Comment) (string, error)
	GetCommentByID(ctx context.Context, id string) (Comment, error)
	DeleteComment(ctx context.Context, id string) error
	LikeBlog(ctx context.Context, like Like) error
	DeleteLikesByBlogID(ctx context.Context, blogID string) error
	DeleteDislikesByBlogID(ctx context.Context, blogID string) error
	DislikeBlog(ctx context.Context, dislike Dislike) error
	UnlikeBlog(ctx context.Context, like Like) error
	UndislikeBlog(ctx context.Context, dislike Dislike) error
}
