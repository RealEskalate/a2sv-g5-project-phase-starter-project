package domain

type User struct {
	ID        string `bson:"_id" json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Password  string `json:"password"`
	IsAdmin   bool   `json:"is_admin"`
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
