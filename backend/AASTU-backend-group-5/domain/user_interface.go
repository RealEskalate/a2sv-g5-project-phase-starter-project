package domain

import "github.com/gin-gonic/gin"

type Controller_interface interface {
	GetOneUser() gin.HandlerFunc
	GetUsers() gin.HandlerFunc
	UpdateUser() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
	LogIn() gin.HandlerFunc
	Register() gin.HandlerFunc
	FilterUser() gin.HandlerFunc
}

type Usecase_interface interface {
	GetOneUser(id string) (ResponseUser , error) 
	GetUsers() ([]ResponseUser , error)
	UpdateUser(id string , user User) (ResponseUser , error)
	DeleteUser(id string) (error)
	LogIn(user User) (ResponseUser , error)
	Register(user User) (ResponseUser , error)
	FilterUser(map[string]string) ([]ResponseUser , error)
}

type Registery_interface interface{
	GetUserDocumentByID(id string) (ResponseUser , error) 
	GetUserDocuments() ([]ResponseUser , error)
	UpdateUserDocument(id string , user User) (ResponseUser , error)
	DeleteUserDocument(id string) (error)
	LogIn(user User) (ResponseUser , error)
	Register(user User) (ResponseUser , error)
	FilterUserDocument(map[string]string) ([]ResponseUser , error)
}
