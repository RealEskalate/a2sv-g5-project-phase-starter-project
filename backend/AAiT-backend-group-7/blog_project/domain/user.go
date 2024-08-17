package domain

import "github.com/gin-gonic/gin"

type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	Password     string `json:"password"`
	Email        string `json:"email"`
	Blogs        []int  `json:"blogs"`
	RefreshToken string `json:"refresh_token"`
	Role         string `json:"role"`
	Bio          string `json:"bio"`
	Phone        string `json:"phone"`
	ProfilePic   string `json:"profile_pic"`
}

type IUser_Repository interface {
	GetAllUsers() ([]User, error)
	GetUserByID(id int) (User, error)
	CreateUser(user User) (User, error)
	UpdateUser(id int, user User) (User, error)
	DeleteUser(id int) error
	SearchByUsername(username string) (User, error)
	SearchByEmail(email string) (User, error)
	AddBlog(user_id int, blog_id int) (User, error)
	RefreshToken(username string, token string) (User, error)
	CreateToken(username string, token string) (User, error)
}

type IUser_Usecases interface {
	GetAllUsers() ([]User, error)
	GetUserByID(id int) (User, error)
	CreateUser(user User) (User, error)
	UpdateUser(id int, user User) (User, error)
	DeleteUser(id int) error
	AddBlog(user_id int, blog_id int) (User, error)
	Login(username string, password string) (User, error)
	ForgetPassword(email string) (User, error)
	ResetPassword(username string, password string) (User, error)
}

type IUser_Controller interface {
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
