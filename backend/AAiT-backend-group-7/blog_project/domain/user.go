package domain

import (
	"context"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID         int    `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Blogs      []int  `json:"blogs"`
	Role       string `json:"role"`
	Bio        string `json:"bio"`
	Phone      string `json:"phone"`
	ProfilePic string `json:"profile_pic"`
}

type IUserRepository interface {
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUserByID(ctx context.Context, id int) (User, error)
	CreateUser(ctx context.Context, user User) (User, error)
	UpdateUser(ctx context.Context, id int, user User) (User, error)
	DeleteUser(ctx context.Context, id int) error
	SearchByUsername(ctx context.Context, username string) (User, error)
	SearchByEmail(ctx context.Context, email string) (User, error)
	AddBlog(ctx context.Context, userID int, blogID int) (User, error)
	RefreshToken(ctx context.Context, username string, token string) (User, error)
	CreateToken(ctx context.Context, username string, token string) (User, error)
}

type IUserUsecase interface {
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUserByID(ctx context.Context, id int) (User, error)
	CreateUser(ctx context.Context, user User) (User, error)
	UpdateUser(ctx context.Context, id int, user User) (User, error)
	DeleteUser(ctx context.Context, id int) error
	AddBlog(ctx context.Context, userID int, blogID int) (User, error)
	Login(ctx context.Context, username, password string) (User, error)
	ForgetPassword(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, username, password string) error
}

type IUserController interface {
	GetAllUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	AddBlog(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	ForgetPassword(c *gin.Context)
	ResetPassword(c *gin.Context)
	PromoUser(c *gin.Context)
	DemoteUser(c *gin.Context)
	RefreshToken(c *gin.Context)
}
