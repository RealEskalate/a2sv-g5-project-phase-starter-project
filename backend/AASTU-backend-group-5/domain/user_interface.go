package domain

import "github.com/gin-gonic/gin"

type User_Controller_interface interface {
	GetOneUser() gin.HandlerFunc
	GetUsers() gin.HandlerFunc
	UpdateUser() gin.HandlerFunc
	DeleteUser() gin.HandlerFunc
	FilterUser() gin.HandlerFunc
}

type User_Usecase_interface interface {
	GetOneUser(id string) (ResponseUser , error) 
	GetUsers() ([]ResponseUser , error)
	UpdateUser(id string , user UpdateUser) (ResponseUser , error)
	DeleteUser(id string) (error)
	FilterUser(map[string]string) ([]ResponseUser , error)
}

type User_Repository_interface interface{
	GetUserDocumentByID(id string) (User , error)
	GetUserDocuments() ([]User , error)
	UpdateUserDocument(id string , user UpdateUser) (User , error)
	DeleteUserDocument(id string) (error)
	FilterUserDocument(filter map[string]string) ([]User , error)
}

