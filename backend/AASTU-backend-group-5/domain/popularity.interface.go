package domain

import "github.com/gin-gonic/gin"

type Like_Controller_inteface interface {
	GetLikes() gin.HandlerFunc
	CreateLike() gin.HandlerFunc
	DeleteLike() gin.HandlerFunc
}

type Like_Usecase_interface interface {
	GetLikes(post_id string) ([]Like , error)
	CreateLike(user_id string , post_id string) error
	DeleteLike(like_id string) error
}

type Like_Repositroy_interface interface {
	GetLikes(post_id string) ([]Like , error)
	CreateLike(user_id string , post_id string) error
	DeleteLike(like_id string) error
}

type DisLike_Controller_inteface interface {
	GetDisLikes() gin.HandlerFunc
	CreateDisLike() gin.HandlerFunc
	DeleteDisLike() gin.HandlerFunc
}

type DisLike_Usecase_interface interface {
	GetDisLikes(post_id string) ([]Like , error)
	CreateDisLike(user_id string , post_id string) error
	DeleteDisLike(like_id string) error
}

type DisLike_Repository_interface interface {
	GetDisLikes(post_id string) ([]Like , error)
	CreateDisLike(user_id string , post_id string) error
	DeleteDisLike(like_id string) error
}

type Comment_Controller_inteface interface {
	GetComments() gin.HandlerFunc
	CreateComment() gin.HandlerFunc
	DeleteComment() gin.HandlerFunc
	UpdateComment() gin.HandlerFunc
}

type Comment_Usecase_interface interface {
	GetComments(post_id string) ([]Comment , error)
	CreateComment(post_id string , user_id string) error
	DeleteComment(comment_id string) error
	UpdateComment(comment_id string) error
}

type Comment_Repository_interface interface {
	GetComments(post_id string) ([]Comment , error)
	CreateComment(post_id string , user_id string) error
	DeleteComment(comment_id string) error
	UpdateComment(comment_id string) error
}

