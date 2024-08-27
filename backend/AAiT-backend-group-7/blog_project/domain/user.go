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
	AddBlog(ctx context.Context, userID int, blog Blog) (User, error)
	StoreRefreshToken(ctx context.Context, userID int, refreshToken string) error
	ValidateRefreshToken(ctx context.Context, userID int, refreshToken string) (bool, error)
	GetRefreshToken(ctx context.Context, userID int) (string, error)
}

type IUserUsecase interface {
	GetAllUsers(ctx context.Context) ([]User, error)
	GetUserByID(ctx context.Context, id int) (User, error)
	CreateUser(ctx context.Context, user User) (User, error)
	UpdateUser(ctx context.Context, id int, user User) (User, error)
	DeleteUser(ctx context.Context, id int) error
	GetUserByUsername(ctx context.Context, username string) (User, error)
	AddBlog(ctx context.Context, userID int, blog Blog) (User, error)
	DeleteBlog(ctx context.Context, userID int, blogID int) (User, error)
	Login(ctx context.Context, username, password string) (string, string, error)
	Logout(ctx context.Context, token string) error
	ForgetPassword(ctx context.Context, email string) error
	ResetPassword(ctx context.Context, username, password string) error
	PromoteUser(ctx context.Context, userID int) (User, error)
	DemoteUser(ctx context.Context, userID int) (User, error)
	RefreshToken(ctx context.Context, refreshToken string) (string, error)
}

type IUserController interface {
	GetAllUsers(c *gin.Context)
	GetUserByID(c *gin.Context)
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	DeleteUser(c *gin.Context)
	Login(c *gin.Context)
	Logout(c *gin.Context)
	ForgetPassword(c *gin.Context)
	ResetPassword(c *gin.Context)
	PromoteUser(c *gin.Context)
	DemoteUser(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type ITokenRepository interface {
	BlacklistToken(ctx context.Context, token string) error
	IsBlacklisted(token string) (bool, error)
}
