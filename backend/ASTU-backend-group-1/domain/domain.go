package domain

import "time"

type User struct {
	ID        string
	Username  string
	Email     string
	FirstName string
	LastName  string
	Password  string
	IsAdmin   bool
}
type Blog struct {
	Title    string
	Content  string
	AuthorID string
	Date     time.Time
	Tags     []string
}
type BlogFilterOption struct {
	Tags       []string
	AuthorId   string
	Date       time.Time
	BlogId     string
	Pagination int
}
type UserFilterOption struct {
	UserID     string
	Email      string
	Username   string
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
