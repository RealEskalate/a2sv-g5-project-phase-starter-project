package blog

import "time"

type Tag struct {
	ID   string `json:"id,omitempty" bson:"_id,omitempty"`
	Name string `json:"name,omitempty"`
}

type Blog struct {
	ID            string    `json:"id,omitempty" bson:"_id,omitempty"`
	AuthorID      string    `json:"author_id,omitempty"`
	Title         string    `json:"title,omitempty"`
	Content       string    `json:"content,omitempty"`
	Tags          []Tag     `json:"tags,omitempty"`
	ViewsCount    int       `json:"views_count,omitempty"`
	CommentsCount int       `json:"comments_count,omitempty"`
	LikesCount    int       `json:"likes_count,omitempty"`
	DislikesCount int       `json:"dislikes_count,omitempty"`
	CreatedAt     time.Time `json:"created_at,omitempty"`
	UpdatedAt     time.Time `json:"updated_at,omitempty"`
}

type Comment struct {
	ID        string    `json:"id,omitempty" bson:"_id,omitempty"`
	BlogID    string    `json:"blog_id,omitempty"`
	AuthorID  string    `json:"author_id,omitempty"`
	Content   string    `json:"content,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

type Like struct {
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	BlogID string `json:"blog_id,omitempty"`
	UserID string `json:"user_id,omitempty"`
}

type Dislike struct {
	ID     string `json:"id,omitempty"`
	BlogID string `json:"blog_id,omitempty"`
	UserID string `json:"user_id,omitempty"`
}

type FilterOption struct {
	Field    string      `json:"field,omitempty"`
	Operator string      `json:"operator,omitempty"`
	Value    interface{} `json:"value,omitempty"`
}

type PaginationRequest struct {
	Limit int `json:"limit,omitempty"`
	Page  int `json:"page,omitempty"`
}

type PaginationResponse[T interface{}] struct {
	Limit      int   `json:"limit,omitempty"`
	Page       int   `json:"page,omitempty"`
	Count      int64 `json:"count,omitempty"`
	TotalPages int   `json:"total_pages,omitempty"`
	Items      []T   `json:"items,omitempty"`
}

type BlogUseCase interface {
	CreateBlog(blog Blog) (Blog, error)
	UpdateBlog(id string, blog Blog) (Blog, error)
	DeleteBlog(id string) error
	GetBlogByID(id string) (Blog, error)
	GetBlogs(filterOptions []FilterOption, pagination PaginationRequest) (PaginationResponse[Blog], error)
	SearchBlogs(query string) (PaginationResponse[Blog], error)
	GetCommentsByBlogID(blogID string, pagination PaginationRequest) (PaginationResponse[Comment], error)
	CreateComment(comment Comment) error
	DeleteComment(id string) error
	LikeBlog(like Like) error
	DislikeBlog(dislike Dislike) error
}

type BlogRepository interface {
	CreateBlog(blog Blog) (string, error)
	UpdateBlog(id string, blog Blog) error
	DeleteBlog(id string) error
	GetBlogByID(id string) (Blog, error)
	GetBlogs(filterOptions []FilterOption, pagination PaginationRequest) (PaginationResponse[Blog], error)
	SearchBlogs(query string) (PaginationResponse[Blog], error)
	GetCommentsByBlogID(blogID string, pagination PaginationRequest) (PaginationResponse[Comment], error)
	CreateComment(comment Comment) error
	DeleteComment(id string) error
	LikeBlog(like Like) error
	DislikeBlog(dislike Dislike) error
}
