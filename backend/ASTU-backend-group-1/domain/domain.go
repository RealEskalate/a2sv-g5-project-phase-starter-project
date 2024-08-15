package domain

import "time"

type User struct {
	UserId    string
	Username  string
	Email     string
	FirstName string
	LastName  string
	Password  string
	IsAdmin   bool
}
type UserFilterOption struct {
	Filter struct {
		UserId    string
		Username  string
		Email     string
		FirstName string
		LastName  string
		IsAdmin   bool
	}
	Pagination int
}
type Blog struct {
	ID       string    `json:"id,omitempty"`
	Title    string    `json:"title,omitempty"`
	Content  string    `json:"content,omitempty"`
	AuthorID string    `json:"author_id,omitempty"`
	Date     time.Time `json:"date,omitempty"`
	Tags     []string  `json:"tags,omitempty"`
}
type BlogFilterOption struct {
	Filter struct {
		Title    string
		Tags     []string
		AuthorId string
		Date     time.Time
		BlogId   string
	}
	Pagination int
}
type BlogRepository interface {
	Get(opts BlogFilterOption) ([]Blog, error)
	Create(b Blog) (Blog, error)
	Update(blogId string, updateData Blog) (Blog, error)
	Delete(blogId string) error
}
type UserRepository interface {
	Get(opts UserFilterOption) ([]User, error)
	Create(u User) (User, error)
	Update(userId string, updateData User) (User, error)
	Delete(userId string) error
}
