package domain

import "context"

type UserRepository interface {
	FindById(cxt context.Context, id string) (*User, Error)
	FindByEmail(cxt context.Context, email string) (*User, Error)
	FindByUsername(cxt context.Context, username string) (*User, Error)
	FindAll(cxt context.Context) ([]User, Error)
	Create(cxt context.Context, user *User) (*User, Error)
	Update(cxt context.Context, id string, user *User) Error
	Delete(cxt context.Context, id string) Error
}

type SessionRepository interface {
	FindTokenById(cxt context.Context, id string) (*Session, Error)
	CreateToken(cxt context.Context, session *Session) (*Session, Error)
	UpdateToken(cxt context.Context, id string, session *Session) Error
	DeleteToken(cxt context.Context, id string) Error
	FindTokenByUserUsername(cxt context.Context, username string) (*Session, bool, Error)
}
type BlogRepository interface {
	FindById(id string) (*Blog, Error)
	FindAll() ([]Blog, Error)
	Create(blog *Blog ) (*Blog, Error)
	Update(blogID string , blog *Blog) (*Blog, Error)
	Delete(id string) Error
	SearchByTitle(title string) ([]Blog, Error)
	SearchByAuthor(author string) ([]Blog, Error)
	Filter(filters map[string]interface{}) ([]Blog, Error)
	AddComment(blogID string, comment *Comment) Error
	DeleteComment(blogID, commentID string) Error
	EditComment(blogID, commentID string, comment *Comment) Error
	Like(id string , userID string) Error
	DisLike(id string , userID string) Error
}

