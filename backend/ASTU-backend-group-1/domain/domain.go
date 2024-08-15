package domain

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
	Tags     []string
}
type BlogRepository interface {
	Get() ([]Blog, error)
	GetByID(blogId string) ([]Blog, error)
	GetByAuthorID(authorID string) ([]Blog, error)
	GetByTags(tags []string) ([]Blog, error)
	Update(blogId string, updateData Blog) (Blog, error)
	Delete(blogId string) error
}
type UserRepository interface {
	Get() ([]User, error)
	GetByID(userID string) ([]User, error)
	GetByUsername(username string) ([]User, error)
	GetByEmail(email string) ([]User, error)
	Update(userId string, updateData User) (User, error)
	Delete(userId string) error
}
