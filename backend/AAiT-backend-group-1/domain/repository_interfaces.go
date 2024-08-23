package domain

import "context"

type UserRepository interface {
	CheckExistence(ctx context.Context, id string) (int, Error)
	CountByEmail(cxt context.Context, email string) (int, Error)
	CountByUsername(cxt context.Context, username string) (int, Error)
	FindById(cxt context.Context, id string) (*User, Error)
	FindByEmail(cxt context.Context, email string) (*User, Error)
	FindByUsername(cxt context.Context, username string) (*User, Error)
	FindAll(cxt context.Context) ([]User, Error)
	Create(cxt context.Context, user *User) (*User, Error)
	UpdateProfile(cxt context.Context, id string, user map[string]interface{}) Error
	UpdatePassword(cxt context.Context, id string, password string) Error
	UpdateRole(cxt context.Context, id string, role string) Error
	Delete(cxt context.Context, id string) Error
	UploadProfilePicture(cxt context.Context, picture Photo, id string) Error
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
	FindAll(page_number string) ([]Blog, Error)
	Create(blog *Blog) (*Blog, Error)
	Update(blogID string, blog *Blog) (*Blog, Error)
	Delete(id string) Error
	SearchByTitle(title string, page_number string) ([]Blog, Error)
	SearchByAuthor(author string, page_number string) ([]Blog, Error)
	Filter(filters map[string]interface{}) ([]Blog, Error)
	AddComment(blogID string, comment *Comment) Error
	DeleteComment(blogID, commentID string) Error
	EditComment(blogID, commentID string, comment *Comment) Error
	Like(id string, userID string) Error
	DisLike(id string, userID string) Error
}
