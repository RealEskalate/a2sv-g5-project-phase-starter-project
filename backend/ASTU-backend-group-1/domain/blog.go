package domain

import "time"

type Blog struct {
	ID       string    `json:"id,omitempty"`
	Title    string    `json:"title,omitempty"`
	Content  string    `json:"content,omitempty"`
	AuthorID string    `json:"author_id,omitempty"`
	Date     time.Time `json:"date,omitempty"`
	Tags     []string  `json:"tags,omitempty"`
}
type BlogFilters struct {
	Title    string
	Tags     []string
	AuthorId string
	Date     time.Time
	BlogId   string
}

type BlogSortOption struct {
	Like     bool
	Dislike  bool
	Comments bool
	View     bool
}

type BlogFilterOption struct {
	Filter     BlogFilters
	Pagination PaginationInfo
	Order      BlogSortOption
}
type BlogRepository interface {
	Create(b Blog) (Blog, error)
	Get(opts BlogFilterOption) ([]Blog, error)
	Update(blogId string, updateData Blog) (Blog, error)
	Delete(blogId string) error
}
