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
type UserFilter struct {
	UserId    string
	Username  string
	Email     string
	FirstName string
	LastName  string
	IsAdmin   bool
}
type UserFilterOption struct {
	Filter     UserFilter
	Pagination int
}
type UserRepository interface {
	Get(opts UserFilterOption) ([]User, error)
	Create(u User) (User, error)
	Update(userId string, updateData User) (User, error)
	Delete(userId string) error
}
