package blog

import "time"

type Tag struct {
	ID   string
	Name string
}

type Blog struct {
	ID            string
	AuthorID      string
	Title         string
	Content       string
	Tags          []Tag
	ViewsCount    int
	CommentsCount int
	LikesCount    int
	DislikesCount int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Comment struct {
	ID        string
	BlogID    string
	AuthorID  string
	Content   string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Like struct {
	ID     string
	BlogID string
	UserID string
}

type Dislike struct {
	ID     string
	BlogID string
	UserID string
}

type FilterOption struct {
	Field    string
	Operator string
	Value    interface{}
}
