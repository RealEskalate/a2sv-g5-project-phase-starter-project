package domain

type UserRepository interface {
	FindById(id string) (*User, Error)
	FindByEmail(email string) (*User, Error)
	FindByUsername(username string) (*User, Error)
	FindAll() ([]User, Error)
	Create(user *User) (*User, Error)
	Update(id string, user *User) Error
	Delete(id string) Error
}

type BlogRepository interface {
	FindById(id string) (*Blog, Error)
	FindAll() ([]Blog, Error)
	Create(blog *Blog) (*Blog, Error)
	Update(blog *Blog) (*Blog, Error)
	Delete(id string) Error
	Search(query string) ([]Blog, Error)
	Filter(filters map[string]interface{}) ([]Blog, Error)
	AddComment(blogID string, comment *Comment) Error
	DeleteComment(blogID, commentID string) Error
	EditComment(blogID, commentID string, comment *Comment) Error
	Like(id string) Error
	DisLike(id string) Error
}